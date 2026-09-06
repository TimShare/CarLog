package system

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")
	if err := h.database.Ping(ctx); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "unhealthy"})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
