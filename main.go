package main

import (
    "fmt"
    "os"
    "github.com/twilio/twilio-go"
    "github.com/twilio/twilio-go/rest/api/v2010"
)

func sendMessage(client *twilio.RestClient) {
    from := "whatsapp:+12677992230" 
    to := "whatsapp:+5586994067338"  
    body := "Levanta, vai estudar Golang"

    _, err := client.Api.CreateMessage(&api.CreateMessageParams{
        To:   &to,
        From: &from,
        Body: &body,
    })

    if err != nil {
        fmt.Println("Error sending message:", err)
        return
    }
    fmt.Println("Message sent successfully!")
}

func main() {
  accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
    authToken := os.Getenv("TWILIO_AUTH_TOKEN")

    client := twilio.NewRestClient(accountSid, authToken)
    sendMessage(client)
}
