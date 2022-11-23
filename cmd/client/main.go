package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/anypb"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"testbufconnect/gen/anyservice/anyserviceconnect"
	greetv1 "testbufconnect/gen/greet/v1"

	"github.com/bufbuild/connect-go"
)

func call(client anyserviceconnect.AnyServiceClient, req protoreflect.ProtoMessage, res protoreflect.ProtoMessage) error {
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
	client := anyserviceconnect.NewAnyServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res := new(greetv1.GreetResponse)
	err := call(client, &greetv1.GreetRequest{
		Name: "Bob",
	}, res)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Greeting)
}
