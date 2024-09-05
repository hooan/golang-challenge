package mail

import (
	"os"
	"strconv"

	"github.com/go-mail/mail"
)

func SendEmail( subject string, body string) {
	from := os.Getenv("FROM")
	password := os.Getenv("PASSWORD")
	to := os.Getenv("TO")
	user := os.Getenv("USERNAME")	
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := mail.NewDialer(os.Getenv("HOST"), port, user, password)

	if err := d.DialAndSend(m); err != nil {
	panic(err)
	}
}

