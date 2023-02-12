package api

import (
	"encoding/json"
	"errors"
	"image/jpeg"
	"io"
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
	}
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Info("Invalid request method")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	submissionBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Unable to decode submission", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	image, runError := controller.Run(submissionBytes)
	if runError != nil {
		var submissionErr *model.SubmissionError
		switch {
		case errors.As(runError, &submissionErr):
			log.Error("Returning client error", runError)
			http.Error(w, runError.Error(), http.StatusBadRequest)
		default:
			log.Error("Unable to create image", runError)
			http.Error(w, runError.Error(), http.StatusInternalServerError)
		}
		return
	}

	encodeError := jpeg.Encode(w, image, &jpeg.Options{Quality: 100})
	if encodeError != nil {
		log.Error("Unable to encode image", encodeError)
		http.Error(w, encodeError.Error(), http.StatusInternalServerError)
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
