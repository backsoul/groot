package main

import (
	"github.com/backsoul/groot/internal/database"
	"github.com/backsoul/groot/pkg/services"
)

func main() {
	database.InitializeDb()
	services.InitializeApi()
}
