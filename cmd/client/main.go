package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/anypb"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	anyservice "testbufconnect/gen/anyservice"
	"testbufconnect/gen/anyservice/anyserviceconnect"

	"github.com/bufbuild/connect-go"
)

func call(client anyserviceconnect.ServiceClient, req protoreflect.ProtoMessage, res protoreflect.ProtoMessage) error {
	reqMsg, err := anypb.New(req)
	if err != nil {
		return err
	}
	resWrapper, err := client.Call(
		context.Background(),
		connect.NewRequest(reqMsg),
	)
	if err != nil {
		return err
	}
	err = resWrapper.Msg.UnmarshalTo(res)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	client := anyserviceconnect.NewServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res := new(anyservice.GreetResponse)
	err := call(client, &anyservice.GreetRequest{
		Name: "Bob",
	}, res)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Greeting)
}
