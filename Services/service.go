package Services

import "time"

type IMyService interface {
	Hello(name string) string
	Health() []Health
}

type MyService struct{}

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

func (s MyService) Health() []Health {
	var health []Health
	app := Health{"HelloWorld", "OK", time.Now().String()}
	health = append(health, app)
	return health
}

func (s MyService) Hello(name string) string {
	return name
}
