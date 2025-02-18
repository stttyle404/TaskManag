package AddTask

import (
	"log"
	"net"
	"os"
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
		log.Println(" Error type logging ", time.Now())
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

	}
}
func HandleConnection(Conn net.Conn) {

}
