package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/anypb"

	anyservice "testbufconnect/gen/anyservice"        // generated by protoc-gen-go
	"testbufconnect/gen/anyservice/anyserviceconnect" // generated by protoc-gen-connect-go
)

type ServiceServer struct{}

func (s *ServiceServer) Call(
	ctx context.Context,
	req *connect.Request[anypb.Any],
) (*connect.Response[anypb.Any], error) {
	log.Println("Request headers: ", req.Header())
	reqMsg := new(anyservice.GreetRequest)
	err := req.Msg.UnmarshalTo(reqMsg)
	if err != nil {
		return nil, err
	}
	resMsg, err := anypb.New(&anyservice.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", reqMsg.Name),
	})
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(resMsg)
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	server := &ServiceServer{}
	mux := http.NewServeMux()
	path, handler := anyserviceconnect.NewServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe("localhost:8080", handler)
}
