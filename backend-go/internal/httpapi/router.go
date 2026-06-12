package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/Grimmjow06100/helloAi/backend-go/internal/config"
	"github.com/Grimmjow06100/helloAi/backend-go/internal/prompts"
)

type API struct {
	config  config.Config
	prompts *prompts.Store
}

func NewRouter(cfg config.Config, promptStore *prompts.Store) http.Handler {
	api := &API{
		config:  cfg,
		prompts: promptStore,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", api.health)
	mux.HandleFunc("GET /prompts", api.promptNames)

	return mux
}

func (api *API) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"env":    api.config.AppEnv,
	})
}

func (api *API) promptNames(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string][]string{
		"prompts": api.prompts.Names(),
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
