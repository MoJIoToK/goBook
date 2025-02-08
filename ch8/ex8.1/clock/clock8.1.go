// Clock8.1 - TCP сервер, выводящий периодически время параллельно.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

// Определение флагов из командной строки для указания порта и часового пояса
var (
	port = flag.String("port", "8000", "server port")
	tz   = flag.String("tz", "Europe/Moscow", "Time Zone")
)

func main() {

	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Clock server is running on %s port with timezone %s", *port, *tz)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, *tz)
		log.Printf("%s is connected", conn.LocalAddr())
	}
}

// handleConn обслуживает клиента, отправляя ему текущее время в заданном
// часовом поясе каждую секунду
func handleConn(c net.Conn, loc string) {
	defer c.Close()

	location, err := time.LoadLocation(loc)
	if err != nil {
		return
	}

	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
