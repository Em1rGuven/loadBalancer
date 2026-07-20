package services

import (
	"net/http"

	"github.com/Em1rGuven/loadBalancer/storage"
)

func NewServiceContainer(backends []*storage.Backend, client *http.Client) *ServiceContainer {
	return &ServiceContainer{
		backends: backends,
		client:   client,
	}
}
