package services

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-mail/mail"
)

func FindProjectRoot() (string, error) {
    dir, err := os.Getwd()
    if err != nil {
        return "", err
    }

    for {
        if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
            return dir, nil
        }

        parent := filepath.Dir(dir)
        if parent == dir {
            return "", fmt.Errorf("project root not found")
        }
        dir = parent
    }
}

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
  //  root, err := FindProjectRoot()
  //  if err != nil {
    //    panic(err)
  //  }

    m.Embed("assets/stori.png")
	log.Println("assets/stori.png")
 //	m.Attach("assets/stori.png")
	d := mail.NewDialer(os.Getenv("HOST"), port, user, password)

	if err := d.DialAndSend(m); err != nil {
	log.Fatalln(err)
	panic(err)
	}
	log.Println("Email sent!")

}
