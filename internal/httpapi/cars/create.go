package cars

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"carlog/internal/service/car"
)

type createRequest struct {
	Make           string  `json:"make"`
	Model          string  `json:"model"`
	VIN            *string `json:"vin"`
	Year           int16   `json:"year"`
	CurrentMileage int32   `json:"current_mileage"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var request createRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		writeError(w, http.StatusBadRequest, "request body must contain one JSON object")
		return
	}

	created, err := h.service.Create(r.Context(), car.CreateParams{
		Make:           request.Make,
		Model:          request.Model,
		VIN:            request.VIN,
		Year:           request.Year,
		CurrentMileage: request.CurrentMileage,
	})
	var validationError car.ValidationError
	if errors.As(err, &validationError) {
		writeError(w, http.StatusBadRequest, validationError.Error())
		return
	}
	if errors.Is(err, car.ErrVINExists) {
		writeError(w, http.StatusConflict, car.ErrVINExists.Error())
		return
	}
	if err != nil {
		slog.Error("create car", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/cars/%d", created.ID))
	writeJSON(w, http.StatusCreated, created)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
