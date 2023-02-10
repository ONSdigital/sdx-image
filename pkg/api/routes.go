package api

import (
	"encoding/json"
	"image/jpeg"
	"net/http"
	"sdxImage/pkg/controller"
	"sdxImage/pkg/log"
	"sdxImage/pkg/model"
)

func Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", healthCheck)
	mux.HandleFunc("/image", handleImage)
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		msg := "Failed to start server"
		log.Error(msg, err)
		panic(msg)
	} else {
		log.Info("listening...")
	}
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Info("Invalid request method")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	var submission model.Submission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		log.Error("Failed to decode submission", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	image, e := controller.Run(&submission)
	if e != nil {
		log.Error("Unable to create image", err)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	err = jpeg.Encode(w, image, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Error("Unable to encode image", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	return
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
