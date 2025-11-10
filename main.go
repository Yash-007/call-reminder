package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func init() {
	godotenv.Load(".env")
}
func scheduleCall() {
	fmt.Println("scheduleCall")
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &api.CreateCallParams{}
	params.SetTwiml("<Response><Say>Please take your medicines and fruits</Say></Response>")
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

func main() {
	c := cron.New()
	fmt.Println("time: ", time.Now().Format("2006-01-02 15:04:05"))
	c.AddFunc("15 12 * * *", scheduleCall)
	c.Start()
	fmt.Println("Cron started")

	s := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Hello, World!")
			scheduleCall()
		}),
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Server started on port 3000")
}
