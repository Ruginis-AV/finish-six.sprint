package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type servise interface {
	ConvertString(input string) string
}

type Handler struct {
	logger  *log.Logger
	servise servise
}

func New(logger *log.Logger, servise servise) *Handler {
	return &Handler{
		logger:  logger,
		servise: servise,
	}
}

func (h *Handler) HtmlHandler(w http.ResponseWriter, r *http.Request) {
	possiblePaths := []string{
		"index.html",
		"../index.html",
		"./internal/index.html",
		"./service/index.html",
		"./server/index.html",
		"./pkg/morse/index.html",
	}

	var foundPath string
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			foundPath = path
			break
		}
	}

	if foundPath == "" {
		h.logger.Printf("index.html not found")
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, foundPath)
}
func (h *Handler) ParseHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(100); err != nil {
		h.logger.Printf("error form %v", err)
		http.Error(w, "error form", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		h.logger.Printf("error read file %v", err)
		http.Error(w, "error read file %v", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		h.logger.Printf("error input string %v", err)
		http.Error(w, "error input string %v", http.StatusInternalServerError)
		return
	}

	result := h.servise.ConvertString(string(input))

	w.Header().Set("Content-Type", "text/plane; charset=utf-8")

	filename := time.Now().UTC().Format("2006-01-02_15-04-05.999999") + ".txt"
	newFile, err := os.Create(filename)
	if err != nil {
		h.logger.Printf("error create file %v", err)
		http.Error(w, "error create file", http.StatusInternalServerError)
		return
	}

	_, err = newFile.WriteString(result)
	if err != nil {
		h.logger.Printf("error create file %v", err)
		http.Error(w, "error create file", http.StatusInternalServerError)
	}
	w.Write([]byte(result))
}
