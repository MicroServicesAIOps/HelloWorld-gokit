package Services

type IMyService interface {
	Hello(name string) string
}

type MyService struct{}

func (this MyService) Hello(name string) string {
	return name
}
