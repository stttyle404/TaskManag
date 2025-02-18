package AddTask

import (
	"database/sql"
	"log"
	"net"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

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
		log.Println(v, " - ", time.Now())
	}
}

func AddTask() {
	Listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		Log(err)
	}
	defer Listener.Close()
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			Log(err)
		}
		HandleConnection(Conn)

	}
}
func HandleConnection(Conn net.Conn) {
	connStr := "user=postgres password=1331 dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	defer Conn.Close()
	input := make([]byte, (1024 * 4))
	ReadedRequest, err := Conn.Read(input)
	if ReadedRequest == 0 || err != nil {
		Log(err)
	}
	source := strings.Fields(string(input[0:ReadedRequest]))

	addtask, err := db.Exec("insert into Task values($1,$2,$3,$4)", source[0], source[1], source[2], source[3])
	if err != nil {
		Log(err)
	}
	Log(addtask)
}

func main() {
	AddTask()
}
