package main

import (
	"net/http"

	"HelloWorld/Services"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {

	s := Services.MyService{}

	hello := Services.MakeServerEndPointHello(s)

	helloServer := httpTransport.NewServer(hello, Services.HelloDecodeRequest, Services.HelloEncodeResponse)

	go http.ListenAndServe("0.0.0.0:8000", helloServer)

	select {}
}
