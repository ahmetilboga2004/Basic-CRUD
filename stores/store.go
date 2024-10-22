package stores

type Store interface {
	Create(item interface{}) error
	GetAll() (interface{}, error)
	Get(id int) (interface{}, error)
	Update(id int, item interface{}) error
	Delete(id int) error
}
