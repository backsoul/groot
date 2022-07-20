package main

import (
	"github.com/backsoul/groot/pkg/services"
)

func main() {
	services.InitializeDb()
	services.InitializeApi()
}
