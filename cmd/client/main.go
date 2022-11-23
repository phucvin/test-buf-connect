package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/anypb"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	greetv1 "testbufconnect/gen/greet/v1"
	"testbufconnect/gen/greet/v1/greetv1connect"

	"github.com/bufbuild/connect-go"
)

func call(client greetv1connect.ServiceClient, req protoreflect.ProtoMessage, res protoreflect.ProtoMessage) error {
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
	client := greetv1connect.NewServiceClient(
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
