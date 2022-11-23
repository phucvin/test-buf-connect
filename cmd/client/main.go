package main

import (
    "context"
    "log"
    "net/http"

	"google.golang.org/protobuf/types/known/anypb"

    greetv1 "testbufconnect/gen/greet/v1"
    "testbufconnect/gen/greet/v1/greetv1connect"

    "github.com/bufbuild/connect-go"
)

func main() {
    client := greetv1connect.NewServiceClient(
        http.DefaultClient,
        "http://localhost:8080",
    )
	reqMsg, err := anypb.New(&greetv1.GreetRequest{
        Name: "John",
    })
	if err != nil {
        log.Println(err)
        return
	}
    res, err := client.Call(
        context.Background(),
        connect.NewRequest(reqMsg),
    )
    if err != nil {
        log.Println(err)
        return
    }
    log.Println(res.Msg.GetTypeUrl())
	resMsg := new(greetv1.GreetResponse)
	err = res.Msg.UnmarshalTo(resMsg)
	if err != nil {
        log.Println(err)
        return
	}
    log.Println(resMsg.GetGreeting())
}