package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/entity"
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/infra/database"
)

type photoHandler struct {
	PhotoDB database.PhotoInterface
}

func NewPhotoHandler(db database.PhotoInterface) *photoHandler {
	return &photoHandler{
		PhotoDB: db,
	}
}

func (h *photoHandler) PhotoUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	validFiles := []string{
		"jpg",
		"png",
		"jpeg",
	}

	file, handler, err := r.FormFile("photo")

	if err != nil {
		http.Error(w, "Error in the process image", http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	p, err := entity.NewFile(handler.Filename, fileBytes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var valid bool

	for _, r := range validFiles {
		valid = strings.HasSuffix(strings.ToLower(handler.Filename), fmt.Sprintf(".%s", r))
		if valid {
			break
		}
	}

	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.PhotoDB.Upload(p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *photoHandler) GetPhotoById(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo, err := h.PhotoDB.FindPhotoByName(name)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photo)

	w.WriteHeader(http.StatusOK)

}
