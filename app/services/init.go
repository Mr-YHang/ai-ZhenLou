package services

type Services struct {
	Session *Session
}

func NewServices() *Services {
	return &Services{
		Session: NewSession(),
	}
}
