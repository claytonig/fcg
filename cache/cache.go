package cache

import (
	"fmt"
)

//Cache struct to store cache data
type Cache struct {
	data map[string]string
}

//InitMemory - initialise instance of cache struct
func InitMemory() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

//ListAll - List all items in cache
func (c *Cache) ListAll() ([]string, error) {
	response := make([]string, 0)
	for _, value := range c.data {
		response = append(response, value)
	}
	return response, nil
}

// ListByKey - list item in cache based on key
func (c *Cache) ListByKey(key string) (string, error) {
	if value, ok := c.data[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("Error::Item not found with key:%s", key)
}

//Add - add item in cache
func (c *Cache) Add(uuid string, item string) error {
	c.data[uuid] = item
	return nil
}

//Remove - remove item from cache
func (c *Cache) Remove(key string) {
	delete(c.data, key)
}
