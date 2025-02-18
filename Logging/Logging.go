package Logging

import (
	"log"
	"os"
	"time"
)

func Logging(Message any) {
	file, err := os.OpenFile("App.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("FileReading Err - ", err, " - ", time.Now())
	}
	defer file.Close()

	log.SetOutput(file)

	switch v := Message.(type) {
	case string:
		log.Println(v, " - ", time.Now())
	case error:
		log.Println(v, " - ", time.Now())
	default:
		log.Println(v, " - ", time.Now())
	}
}
