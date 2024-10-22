package stores

type Store interface {
	Create(item any) error
	GetAll() (any, error)
	Get(id int) (any, error)
	Update(id int, item any) error
	Delete(id int) error
}
