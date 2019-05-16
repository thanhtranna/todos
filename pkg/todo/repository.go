package todo

type Repository interface {
	Delete(id string, userId string) error
	GetAll(userId string) ([]*Todo, error)
	GetByID(id string, userId string) (*Todo, error)
	Store(u *Todo, userId string) error
	Update(u *Todo, userId string) error
}