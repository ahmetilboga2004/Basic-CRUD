package stores

type Validator[T any] interface {
	Validate() error
}

type Store[T Validator[T]] interface {
	Create(item T) (T, error)
	GetAll() ([]T, error)
	Get(id int) (T, error)
	Update(id int, item T) error
	Delete(id int) error
}
