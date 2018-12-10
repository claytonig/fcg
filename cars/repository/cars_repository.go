package repository

import (
	"encoding/json"
	m "fcg/adapters"
	c "fcg/cars"
	"fcg/logger"
	"fmt"
)

//CarsRepository - struct for Car Repo
type CarsRepository struct {
	memAdapter m.MemoryAdapter
}

// NewCarRepository - Repository layer for cars
func NewCarRepository(memoryAdapter m.MemoryAdapter) c.Repository {
	return &CarsRepository{
		memAdapter: memoryAdapter,
	}
}

//GetAll - Get all cars
func (cr *CarsRepository) GetAll() ([]*c.Car, error) {

	vals, err := cr.memAdapter.GetAll()
	cars := make([]*c.Car, 0)
	if err != nil {
		logger.Log().WithError(err).Errorf("Error while fetching cars from storage")
		return nil, err
	}
	for _, x := range vals {
		bytes := []byte(x)
		var car *c.Car
		err = json.Unmarshal(bytes, &car)
		if err != nil {
			logger.Log().Errorf("Error while un-marshaling car")
			return nil, err
		}

		cars = append(cars, car)
	}

	return cars, nil
}

//GetByID - Get a car by ID
func (cr *CarsRepository) GetByID(ID string) (*c.Car, error) {

	val, err := cr.memAdapter.GetByKey(ID)

	if err != nil {
		logger.Log().WithError(err).Errorf("Error while fetching car from storage, key:%s", ID)
		return nil, err
	}

	if val == "" {
		logger.Log().WithError(err).Errorf("No Cars found")
		return nil, fmt.Errorf("No Cars found")
	}

	bytes := []byte(val)
	var car *c.Car

	err = json.Unmarshal(bytes, &car)
	if err != nil {
		logger.Log().WithError(err).WithField("uuid", ID).Errorf("Error while un-marshaling car")
		return nil, err
	}

	return car, nil

}

//Create - create a car
func (cr *CarsRepository) Create(carToCreate map[string]interface{}, uuidString string) (string, error) {

	var ok bool

	var make interface{}
	if make, ok = carToCreate["make"]; !ok && make.(string) != "" {
		logger.Log().WithField("make", carToCreate).Error("make is mandatory to create a car")
		return "", fmt.Errorf("make is mandatory to create a car")
	}

	var model interface{}
	if model, ok = carToCreate["model"]; !ok && model.(string) != "" {
		logger.Log().WithField("model", carToCreate).Error("model is mandatory to create a car")
		return "", fmt.Errorf("model is mandatory to create a car")
	}

	var year interface{}
	if year, ok = carToCreate["year"]; !ok && year.(string) != "" {
		logger.Log().WithField("year", carToCreate).Error("year is mandatory to create car")
		return "", fmt.Errorf("year is mandatory to create a car")
	}

	newCar := c.Car{
		ID:    uuidString,
		Make:  make.(string),
		Model: model.(string),
		Year:  year.(string),
	}

	carBytes, err := json.Marshal(newCar)

	if err != nil {
		logger.Log().WithError(err).Errorf("Error while converting car to json")
		return "", err
	}

	err = cr.memAdapter.Create(uuidString, string(carBytes))
	if err != nil {
		logger.Log().WithError(err).Errorf(
			"Error while storing car")
	}
	return uuidString, nil

}

// Delete - delete a car based on ID
func (cr *CarsRepository) Delete(ID string) error {

	err := cr.memAdapter.Remove(ID)

	if err != nil {
		logger.Log().WithField("uuid", ID).WithError(err).Errorf("Error while deleting car")
	}

	return nil
}
