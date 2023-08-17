package main

import (
	"fmt"
	"irg1008/next-go/pkg/config"
	"irg1008/next-go/pkg/mail"
)

func main() {
	config := config.NewConfig()
	mailService := mail.NewMail(config.Domain, config.ResendKey)

	sender := mailService.NewSender("No responder", "no-reply")
	url := "https://localhost:3000/confirm/123456789"

	err := sender.SendConfirmEmail("ivansudevlop@gmail.com", "Confirm email", url)

	fmt.Print(err)
}
