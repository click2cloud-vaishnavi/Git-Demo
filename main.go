package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/save", saveImage)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func saveImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, _ := r.FormFile("image")
	defer file.Close()

	filename := handler.Filename
	outFile, _ := os.Create(filepath.Join(".", filename))
	defer outFile.Close()

	_, _ = io.Copy(outFile, file)

	w.Write([]byte("Image saved successfully"))
}
