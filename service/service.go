package service

type Service interface {
	UserService
}

type service struct {
}

func NewService() Service {
	svc := &service{}
	return svc
}
