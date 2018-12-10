package adapters

import (
	c "fcg/cache"
	"fmt"
)

//MemoryAdapterRepository - struct to store storage connection data
type MemoryAdapterRepository struct {
	cacheConn *c.Cache
}

//NewMemory - Creates instance of memoryAdapterRepository
func NewMemory() *MemoryAdapterRepository {

	return &MemoryAdapterRepository{
		cacheConn: c.InitMemory(),
	}

}

//GetAll - Get all items
func (ma *MemoryAdapterRepository) GetAll() ([]string, error) {

	data, err := ma.cacheConn.ListAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

//GetByKey - Get item by key
func (ma *MemoryAdapterRepository) GetByKey(id string) (string, error) {

	data, err := ma.cacheConn.ListByKey(id)

	if err != nil {
		return "", err
	}

	return data, nil

}

//Create - Create an item
func (ma *MemoryAdapterRepository) Create(uuid string, item string) error {

	err := ma.cacheConn.Add(uuid, item)
	return err

}

//Remove - delete an item based on key
func (ma *MemoryAdapterRepository) Remove(id string) error {

	_, err := ma.cacheConn.ListByKey(id)

	if err != nil {
		return fmt.Errorf("Item does not exist: %s", id)
	}

	ma.cacheConn.Remove(id)

	return nil
}
