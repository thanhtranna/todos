package login

type Service interface {
	Signin(userName string, password string) (string, error)
}

type loginService struct {
	repo Repository
}

func NewLoginService(repo Repository) Service {
	return &loginService{
		repo: repo,
	}
}

func (svc *loginService) Signin(email string, password string) (string, error) {
	return svc.repo.Signin(&Login{Email: email, Password: password})
}