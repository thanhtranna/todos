package todo

type Service interface {
	Delete(id string, userId string) error
	GetAll(userId string) ([]*Todo, error)
	GetByID(id string, userId string) (*Todo, error)
	Store(u *Todo) error
	Update(u *Todo) error
}

type todoService struct {
	repo Repository
}

func NewTodoService(repo Repository) Service {
	return &todoService{
		repo: repo,
	}
}

func (svc *todoService) Delete(id, userId string) error { return svc.repo.Delete(id, userId) }

func (svc *todoService) GetAll(userId string) ([]*Todo, error) { return svc.repo.GetAll(userId) }

func (svc *todoService) GetByID(id, userId string) (*Todo, error) { return svc.repo.GetByID(id, userId) }

func (svc *todoService) Store(u *Todo) error { return svc.repo.Store(u) }

func (svc *todoService) Update(u *Todo) error { return svc.repo.Update(u) }