package adapters

// MemoryAdapter - Adapter to talk to cache
type MemoryAdapter interface {
	GetAll() ([]string, error)
	GetByKey(id string) (string, error)
	Create(uuid string, item string) error
	Remove(id string) error
}
