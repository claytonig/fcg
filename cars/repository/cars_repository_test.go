package repository

import (
	"encoding/json"
	"fcg/adapters/mocks"
	"fcg/cars"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarRepository_GetAll(t *testing.T) {
	mockMemory := new(mocks.MemoryAdapter)
	mockCars := []*cars.Car{
		{
			ID:    "f0803291-6248-4620-4398-392cb80c4132",
			Make:  "Lamborghini",
			Model: "Gallardo",
			Year:  "2012",
		},
		{
			ID:    "t3453453-6248-4620-4398-392cb80c4132",
			Make:  "Lamborghini",
			Model: "Gallardo",
			Year:  "2012",
		},
	}

	mockSlice := make([]string, 0)
	for _, mockCar := range mockCars {
		b, err := json.Marshal(mockCar)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when Marshalling", err)
		}
		mockSlice = append(mockSlice, string(b))
	}

	mockMemory.On("GetAll").Return(mockSlice, nil)

	carRepo := NewCarRepository(mockMemory)
	cars, err := carRepo.GetAll()

	assert.NoError(t, err)
	assert.ObjectsAreEqualValues(mockCars, cars)
	mockMemory.AssertExpectations(t)
}

func TestCarRepository_GetByID(t *testing.T) {
	mockMemory := new(mocks.MemoryAdapter)
	mockCar := &cars.Car{
		ID:    "f0803291-6248-4620-4398-392cb80c4132",
		Make:  "Lamborghini",
		Model: "Gallardo",
		Year:  "2012",
	}

	b, err := json.Marshal(mockCar)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when Marshalling", err)
	}

	key := "f0803291-6248-4620-4398-392cb80c4132"
	mockMemory.On("GetByKey", key).Return(string(b), nil)

	carRepo := NewCarRepository(mockMemory)
	car, _ := carRepo.GetByID("f0803291-6248-4620-4398-392cb80c4132")

	assert.NoError(t, err)
	assert.ObjectsAreEqualValues(mockCar, car)
	mockMemory.AssertExpectations(t)
}

func TestCarRepository_Create(t *testing.T) {
	car := &cars.Car{
		ID:    "f0803291-6248-4620-4398-392cb80c4132",
		Make:  "Lamborghini",
		Model: "Gallardo",
		Year:  "2012",
	}

	mockMemory := new(mocks.MemoryAdapter)

	uuid := "f0803291-6248-4620-4398-392cb80c4132"

	carBytes, _ := json.Marshal(car)
	mockMemory.On("Create", uuid, string(carBytes)).Return(nil)

	carRepo := NewCarRepository(mockMemory)
	actualUUID, err := carRepo.Create(map[string]interface{}{
		"make":  "Lamborghini",
		"model": "Gallardo",
		"year":  "2012",
	}, uuid)

	assert.NoError(t, err)
	assert.ObjectsAreEqualValues(actualUUID, uuid)
	mockMemory.AssertExpectations(t)
}

func TestCarRepository_Delete(t *testing.T) {
	mockMemory := new(mocks.MemoryAdapter)

	deleteUUID := "f0803291-6248-4620-4398-392cb80c4132"

	mockMemory.On("Remove", deleteUUID).Return(nil)

	carRepo := NewCarRepository(mockMemory)
	err := carRepo.Delete(deleteUUID)

	assert.NoError(t, err)
	mockMemory.AssertExpectations(t)
}
