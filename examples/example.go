package main

import (
	"fmt"
	"os"
	"time"

	"github.com/streem/streem-sdk-go"
)

const (
	apiKeyId       = "api_id"
	apiKeySecret   = "api_key"
	apiEnvironment = "prod-us"
)

func main() {
	err := streem.Init(apiKeyId, apiKeySecret, apiEnvironment)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	builder := streem.NewTokenBuilder()

	// required
	builder.SetUserId("someId")

	// recommended
	builder.SetName("T Rex")
	builder.SetEmail("some@email.com")
	builder.SetAvatarUrl("some.avatar.url")

	// optional
	addDurationToNow := builder.AddTime()
	builder.SetTokenExpirationMs(addDurationToNow(time.Minute * 30))
	builder.SetSessionExpirationMs(addDurationToNow(time.Hour * 12))

	token, err := builder.Build()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("Got token %s\n", token)

}
