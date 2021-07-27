package main

import (
	"net/http"

	Services "HelloWorld/Services"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {

	s := Services.MyService{}

	hello := Services.MakeServerEndPointHello(s)
	health := Services.MakeServerEndpointHealth(s)

	helloServer := httpTransport.NewServer(hello, Services.HelloDecodeRequest, Services.HelloEncodeResponse)
	healthServer := httpTransport.NewServer(health, Services.HealthDecodeRequest, Services.HealthEncodeResponse)

	http.Handle("/hello/", helloServer)
	http.Handle("/health", healthServer)
	go http.ListenAndServe("0.0.0.0:8000", nil)

	select {}
}
