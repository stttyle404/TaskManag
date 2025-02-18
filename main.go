package main

import (
	"log"
	"net"
	"os"
	"time"

	_ "github.com/lib/pq"
	//"database/sql"
	//"net/http"
)

type Task struct {
	id          string
	title       string
	description string
	overdate    string
	status      string
}

func Log(Message any) {
	file, err := os.OpenFile("Err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
		log.Println(" Error type logging ", time.Now())
	}
}

func (t Task) AddTask(Title string, Description string, Overdate string, Status string) {
	request := ""
	request = Title + " " + Description + " " + Overdate + " " + Status
	Conn, err := net.Dial("tcp", "localhost:8080")

	defer Conn.Close()

	if err != nil {
		Log(err)
	}
	if _, err := Conn.Write([]byte(request)); err != nil {
		Log(err)
	} else {
		Log("Adding Task Succesfuly")
	}
}

func (t Task) ListTask(Title string, Description string, Overdate time.Time, Status string) {

}

func (t Task) UpdateTask(Title string, Description string, Overdate time.Time, Status string) {

}

func (t Task) DeleteTask(Title string, Description string, Overdate time.Time, Status string) {

}

func (t Task) ExportTasks(Title string, Description string, Overdate time.Time, Status string) {

}

func (t Task) ImportTasks(Title string, Description string, Overdate time.Time, Status string) {

}
