package cars

//Car - struct for car details
type Car struct {
	ID    string `json:"id,omitempty"`
	Make  string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
	Year  string `json:"year,omitempty"`
}

// Repository layer for cars
type Repository interface {
	GetAll() ([]*Car, error)
	GetByID(ID string) (*Car, error)
	Create(carMap map[string]interface{}, uuidString string) (string, error)
	Delete(ID string) error
}
