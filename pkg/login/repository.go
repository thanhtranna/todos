package login

type Repository interface {
	Signin(*Login) (string, error)
}