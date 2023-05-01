package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"example/internal"

	"github.com/garyburd/redigo/redis"
	"github.com/go-nacelle/httpbase"
	nacelle "github.com/go-nacelle/nacelle/v2"
)

type ServerInitializer struct {
	Logger nacelle.Logger `service:"logger"`
	Redis  redis.Conn     `service:"redis"`
}

func NewServerInitializer() httpbase.ServerInitializer {
	return &ServerInitializer{}
}

func (si *ServerInitializer) Init(ctx context.Context, server *http.Server) error {
	server.Handler = http.HandlerFunc(si.handle)
	return nil
}

func (si *ServerInitializer) handle(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Path[1:]; id != "" && r.Method == "GET" {
		si.handleGet(w, r, id)
		return
	}

	if id := r.URL.Path[1:]; id == "" && r.Method == "POST" {
		si.handlePost(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (si *ServerInitializer) handleGet(w http.ResponseWriter, r *http.Request, id string) {
	response, err := internal.GetResult(si.Redis, id)
	if err != nil {
		si.error(w, "Failed to read request body (%s)", err)
		return
	}

	if response == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialized, err := json.Marshal(response)
	if err != nil {
		si.error(w, "Failed to serialize response (%s)", err)
		return
	}

	si.Logger.Debug("Retrieved key %s", id)
	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

func (si *ServerInitializer) handlePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		si.error(w, "Failed to read request body (%s)", err)
		return
	}

	id, err := internal.PublishWork(si.Redis, string(content))
	if err != nil {
		si.error(w, "Failed to publish work (%s)", err)
		return
	}

	serialized, err := json.Marshal(map[string]string{"id": id})
	if err != nil {
		si.error(w, "Failed to serialize response (%s)", err)
		return
	}

	si.Logger.Debug("Published work with id %s", id)
	w.WriteHeader(http.StatusAccepted)
	w.Write(serialized)
}

func (si *ServerInitializer) error(w http.ResponseWriter, format string, err error) {
	si.Logger.Error(format, err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}
