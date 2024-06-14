package repository

type CrudRepository[ID, T any] interface {
	Count() (int, error)
	DeleteById(id ID) error
	Delete(entity *T) error
	DeleteAll() error
	DeleteAllById(ids []ID) error
	ExistsById(id ID) (bool, error)
	FindAll() ([]T, error)
	FindById(id ID) (*T, error)
	Save(entity *T) (*T, error)
	SaveAll(entities ...T) ([]T, error)
}
