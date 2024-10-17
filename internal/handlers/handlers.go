package handlers

import "github.com/Rizabekus/music-library/internal/services"

type Handlers struct {
	Service *services.Services
}

func HandlersInstance(services *services.Services) *Handlers {
	return &Handlers{Service: services}
}
