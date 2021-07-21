package service

type UserService struct {
	Id   int
	Name string
	Impl UserServiceInterface
}

type UserServiceInterface interface {
	DoA()
	DoB()
}

type UserServiceImpl struct{}

func (p *UserServiceImpl) DoA() {
	println("UserServiceImpl DoA")
}
func (p *UserServiceImpl) DoB() {
	println("UserServiceImpl DoB")
}

func NewUserService(id int, name string) *UserService {
	return &UserService{Id: id, Name: name, Impl: &UserServiceImpl{}}
}
