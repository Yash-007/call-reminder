package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	godotenv.Load(".env")
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &api.CreateCallParams{}
	params.SetTwiml("<Response><Say>Please take your medicine</Say></Response>")
	params.SetTo("+918770798425")
	params.SetFrom("+12513136008")
	params.SetStatusCallback("http://localhost:3000/call-res")
	params.SetStatusCallbackEvent([]string{"initiated,ringing, answered, completed"})
	params.SetStatusCallbackMethod("POST")

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}
