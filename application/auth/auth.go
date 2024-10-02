package auth

type CreateUserInput struct {
	Name           string
	Email          string
	Password       string
	RepeatPassword string
}

type AuthService interface {
	CreateUser(input CreateUserInput)
	AuthenticateUser()
}

type AuthServiceImpl struct {
}

func (service *AuthServiceImpl) CreateUser(input CreateUserInput) {

}
