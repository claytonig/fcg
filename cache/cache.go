package cache

type cache struct {
	data map[string]interface{}
}

func initMemory() *cache {
	return nil
}

func (c *cache) ListAll() (interface{}, error) {
	return nil, nil
}

func (c *cache) ListByKey(key string) ([]interface{}, error) {
	return nil, nil
}

func (c *cache) Add(item interface{}) (string, error) {
	return "", nil
}

func (c *cache) Remove(key string) error {
	return nil
}
