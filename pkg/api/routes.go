package api

import (
	"fmt"
	"image/jpeg"
	"net/http"
	"sdxImage/pkg/controller"
)

func Listen() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/image", handler)
	http.ListenAndServe(":5000", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	image := controller.Run("abs_1802")
	err := jpeg.Encode(w, image, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	return
}
