package main

import (
	"sync"

	"fcg/logger"

	adapterRepository "fcg/adapters/repository"
	carsController "fcg/cars/courier/http"
	cr "fcg/cars/repository"

	"github.com/labstack/echo"
)

var onceRest sync.Once

func main() {
	onceRest.Do(func() {
		e := echo.New()

		memAdapter := adapterRepository.NewMemory()
		carRepo := cr.NewCarRepository(memAdapter)
		carsController.NewCarController(e, carRepo)

		if err := e.Start("0.0.0.0:8085"); err != nil {
			logger.Log().WithError(err).Fatal("Unable to start the service")
		}
	})
}
