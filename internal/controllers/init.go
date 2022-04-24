package controllers

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/services"
)

type Service struct{ service *services.Repository }

func Init(service *services.Repository) *Service {
	return &Service{service}
}
