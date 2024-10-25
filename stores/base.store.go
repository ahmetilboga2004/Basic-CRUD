package stores

import "errors"

type Validator[T any] interface {
	Validate() error
	SetID(id int)
}

type Store[T Validator[T]] interface {
	Create(item T) error
	GetAll() ([]T, error)
	Get(id int) (T, error)
	Update(id int, item T) error
	Delete(id int) error
}

type BaseStore[T Validator[T]] struct {
	items  map[int]T
	nextID int
}

func NewBaseStore[T Validator[T]]() *BaseStore[T] {
	return &BaseStore[T]{
		items:  make(map[int]T),
		nextID: 1,
	}
}

func (bs *BaseStore[T]) Create(item T) error {
	if err := item.Validate(); err != nil {
		return err
	}
	item.SetID(bs.nextID)
	bs.items[bs.nextID] = item
	bs.nextID++
	return nil
}

func (bs *BaseStore[T]) GetAll() ([]T, error) {
	if len(bs.items) == 0 {
		return nil, errors.New("no items found")
	}
	var itemList []T
	for _, item := range bs.items {
		itemList = append(itemList, item)
	}

	return itemList, nil
}

func (bs *BaseStore[T]) Get(id int) (T, error) {
	if item, exists := bs.items[id]; exists {
		return item, nil
	}
	var zero T
	return zero, errors.New("item not found")
}

func (bs *BaseStore[T]) Update(id int, item T) error {
	if _, exists := bs.items[id]; !exists {
		return errors.New("item not found")
	}
	if err := item.Validate(); err != nil {
		return err
	}
	bs.items[id] = item
	return nil
}

func (bs *BaseStore[T]) Delete(id int) error {
	if _, exists := bs.items[id]; !exists {
		return errors.New("item not found")
	}
	delete(bs.items, id)
	return nil
}
