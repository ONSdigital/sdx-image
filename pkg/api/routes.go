package api

import (
	"encoding/json"
	"image/jpeg"
	"log"
	"net/http"
	"sdxImage/pkg/controller"
	"sdxImage/pkg/model"
)

func Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", healthCheck)
	mux.HandleFunc("/image", handleImage)
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	var submission model.Submission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	image, e := controller.Run(&submission)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	err = jpeg.Encode(w, image, &jpeg.Options{Quality: 100})
	if err != nil {
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
