package modules

import (
	"TaskManager/Logging"
	"database/sql"
	"net"
	"strings"

	_ "github.com/lib/pq"
)

func ConnectToFunction(input string) {
	Listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		Logging.Logging(err)
	}
	defer Listener.Close()
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			Logging.Logging(err)
		}
		switch input {
		case "AddTask":
			AddTaskFunc(Conn)
		default:
			panic(err)
		}

	}
}
func AddTaskFunc(Conn net.Conn) {
	connStr := "user=postgres password=1331 dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	defer Conn.Close()
	input := make([]byte, (1024 * 4))
	ReadedRequest, err := Conn.Read(input)
	if ReadedRequest == 0 || err != nil {
		Logging.Logging(err)
	}
	source := strings.Fields(string(input[0:ReadedRequest]))

	addtask, err := db.Exec("insert into Task values($1,$2,$3,$4)", source[0], source[1], source[2], source[3])
	if err != nil {
		Logging.Logging(err)
	}
	Logging.Logging(addtask)
}
func DeleteTaskFunc(Conn net.Conn) {
	connStr := "user=postgres password=1331 dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	defer Conn.Close()
	input := make([]byte, (1024 * 4))
	ReadedRequest, err := Conn.Read(input)
	if ReadedRequest == 0 || err != nil {
		Logging.Logging(err)
	}
	source := strings.Fields(string(input[0:ReadedRequest]))

	addtask, err := db.Exec("delete into Task values($1,$2,$3,$4)", source[0], source[1], source[2], source[3])
	if err != nil {
		Logging.Logging(err)
	}
	Logging.Logging(addtask)
}

func main() {
	ConnectToFunction("AddTask")
}
