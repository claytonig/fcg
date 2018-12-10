package http

import (
	"fcg/cars"
	"fcg/logger"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/nu7hatch/gouuid"
)

//CarController - struct for car controller type
type CarController struct {
	repo cars.Repository
}

// GetCars - Get all cars
func (cc *CarController) GetCars(c echo.Context) error {

	var carsList []*cars.Car
	var err error

	carsList, err = cc.repo.GetAll()

	if err != nil {
		logger.Log().WithError(err).Errorf("Error occurred while getting cars")
		return c.JSON(http.StatusNotFound, make(map[string]interface{}))
	}

	return c.JSON(http.StatusOK, carsList)
}

//GetByID - get a car by ID
func (cc *CarController) GetByID(c echo.Context) error {
	carID := c.Param("carID")

	var car *cars.Car
	var err error

	car, err = cc.repo.GetByID(carID)

	if err != nil {
		logger.Log().WithError(err).Errorf("Error occurred while getting Car")
		return c.JSON(http.StatusNotFound, make(map[string]interface{}))
	}

	return c.JSON(http.StatusOK, car)
}

//CreateCar - create a car
func (cc *CarController) CreateCar(c echo.Context) error {
	var carMap map[string]interface{}
	err := c.Bind(&carMap)

	response := make(map[string]interface{})

	if err != nil {
		logger.Log().WithError(err).Errorf("Could not bind request body to car model, sending bad request")
		response["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	u4, err := uuid.NewV4()

	uuidString := u4.String()
	if err != nil {
		return fmt.Errorf("Error::Could not add item")
	}

	id, err := cc.repo.Create(carMap, uuidString)
	if err != nil {
		logger.Log().WithError(err).WithField("car", carMap).Errorf("Error while creating car")
		response["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusOK, id)
}

//DeleteCar - delete a car based on key
func (cc *CarController) DeleteCar(c echo.Context) error {
	carID := c.Param("carID")

	response := make(map[string]interface{})

	err := cc.repo.Delete(carID)

	if err != nil {
		logger.Log().WithError(err).WithField("id", carID).Error("Error while deleting car, sending bad request")
		response["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}
	response["result"] = "Successfully Deleted"
	return c.JSON(http.StatusOK, response)

}

//NewCarController - Create an instance on CarController
func NewCarController(e *echo.Echo, carRepo cars.Repository) {
	carHandler := &CarController{
		repo: carRepo,
	}

	e.GET("/fcg/cars", carHandler.GetCars)
	e.GET("/fcg/cars/:carID", carHandler.GetByID)
	e.POST("/fcg/cars", carHandler.CreateCar)
	e.DELETE("/fcg/cars/:carID", carHandler.DeleteCar)

}
