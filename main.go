package main

import (
	"TaskManager/Logging"
	"net"

	// "time"

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

func (t Task) AddTask(Title string, Description string, Overdate string, Status string) {
	request := ""
	request = Title + " " + Description + " " + Overdate + " " + Status
	Conn, err := net.Dial("tcp", "localhost:8080")

	defer Conn.Close()

	if err != nil {
		Logging.Logging(err)
	}
	if _, err := Conn.Write([]byte(request)); err != nil {
		Logging.Logging(err)
	} else {
		Logging.Logging("Adding Task Succesfuly")
	}
}

// func (t Task) ListTask(Title string, Description string, Overdate time.Time, Status string) {

// }

// func (t Task) UpdateTask(Title string, Description string, Overdate time.Time, Status string) {

// }

// func (t Task) DeleteTask(Title string, Description string, Overdate time.Time, Status string) {

// }

// func (t Task) ExportTasks(Title string, Description string, Overdate time.Time, Status string) {

// }

// func (t Task) ImportTasks(Title string, Description string, Overdate time.Time, Status string) {

// }
