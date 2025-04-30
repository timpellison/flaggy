package server

import (
	"encoding/json"
	"github.com/timpellison/flaggy/internal/model"
	"log/slog"
	"net/http"
)

type DataService interface {
	GetFlag(key string) (*model.FeatureFlag, error)
	Add(featureFlag model.FeatureFlag) error
	Update(featureFlag model.FeatureFlag) error
	Delete(featureFlag model.FeatureFlag) error
}

type FlaggyService struct {
	dataService DataService
	logger      *slog.Logger
}

func NewFlaggyService(dataService DataService) *FlaggyService {
	l := slog.Default()
	return &FlaggyService{dataService: dataService, logger: l}
}

func (fs *FlaggyService) Start() error {

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Healthy!\n"))
	})
	http.HandleFunc("GET /featureflag/{key}", fs.getFlag())
	http.HandleFunc("POST /featureflag", fs.addFlag())
	http.HandleFunc("PUT /featureflag/{key}/{enabled}", fs.updateFlag())
	http.HandleFunc("DELETE /featureflag/{key}", fs.deleteFlag())
	fs.logger.Info("Starting server")
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		return e
	}
	return nil
}

func (fs *FlaggyService) getFlag() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		f, e := fs.dataService.GetFlag(key)
		if e != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		w.Header().Set("Content-Type", "application/json")
		en := json.NewEncoder(w)
		_ = en.Encode(f)
	}
}

func (fs *FlaggyService) addFlag() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var f model.FeatureFlag
		err := decoder.Decode(&f)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		_ = fs.dataService.Add(f)
		w.WriteHeader(http.StatusCreated)
	}
}

func (fs *FlaggyService) updateFlag() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		enabled := r.PathValue("enabled")
		if enabled == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		f := &model.FeatureFlag{
			Key:   key,
			Value: enabled == "true",
		}
		_ = fs.dataService.Update(*f)
		w.WriteHeader(http.StatusOK)
	}
}

func (fs *FlaggyService) deleteFlag() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		decoder := json.NewDecoder(r.Body)
		var f model.FeatureFlag
		err := decoder.Decode(&f)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		_ = fs.dataService.Delete(f)
		w.WriteHeader(http.StatusOK)
	}
}
