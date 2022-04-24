package main

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/controllers"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/repositories"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/services"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/infrastructure"
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/router"
)

func main() {
	db := infrastructure.ConnectDatabase()
	repo := repositories.Init(db)
	serv := services.Init(repo)
	ctlr := controllers.Init(serv)
	api := router.Init(ctlr)
	api.Start()
}
