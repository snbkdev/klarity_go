package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 {
	return 0
}

const sender = "test@test.asia"
const password = "p@ssw0rd"
const hostname = "smtp.test.asia"
const template = `Внимание, вы использовали %d байт хранилища, %d%% вашей квоты`

func CheckQuota(username string) {
	used := bytesInUse(username)

	const quota = 1000000000
	percent := 100 * used / quota
	if percent < 90 {
		return
	}

	msg := fmt.Sprintf(template, used, percent)
	auth  := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("Сбой smtp.SendMail(%s): %s", username, err)
	}
}