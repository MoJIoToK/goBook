//Clockwall - TCP-клиент только для чтения несколькиз серверов

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

// Server - сервер времени
type server struct {
	name    string
	host    string
	message string
}

func main() {
	delay := time.Duration(1 * time.Second)

	//Проверка аргументов командной строки
	if len(os.Args) < 2 {
		fmt.Println("Usage: name=localhost:port")
		os.Exit(1)
	}

	//Парсинг списка серверов
	servers, err := parseServer(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	//Подключение к серверу и запись времени из сервера в структуру.
	for _, srv := range servers {
		conn, err := net.Dial("tcp", srv.host)
		if err != nil {
			log.Fatalf("Error connecting to %s: %v", srv.host, err)
			return
		}
		fmt.Printf("Connecting to %s\n", srv.host)
		defer conn.Close()

		//Получение времени от сервера в горутине
		go func(conn io.ReadCloser) {
			if err := srv.getMessage(conn); err != nil {
				log.Fatal(err)
			}
		}(conn)

	}

	printMessage(servers, delay)

}

// printMessage - функция для печати времени из структуры с
// определенным интервалом.
func printMessage(servers []*server, delay time.Duration) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Server", "Message")
	fmt.Fprintf(tw, format, "----", "------", "-------")
	for {
		for _, c := range servers {
			fmt.Fprintf(tw, format, c.name, c.host, c.message)
		}
		fmt.Fprintf(tw, format, "----", "----", "----")
		tw.Flush()
		time.Sleep(delay)
	}
}

// parseServer - функция для парсинга серверов из командной строки
func parseServer(args []string) ([]*server, error) {
	servers := make([]*server, 0, len(args))
	for _, arg := range os.Args[1:] {
		part := strings.Split(arg, "=")
		if len(part) < 2 {
			return nil, fmt.Errorf("invalid argument", arg)
		}
		servers = append(servers,
			&server{
				part[0],
				part[1],
				"",
			})
	}
	return servers, nil
}

// getMessage - метод для получения времени из подключения и сохранения
// в поле структуры.
func (srv *server) getMessage(src io.ReadCloser) error {
	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		srv.message = scanner.Text()
	}

	return scanner.Err()
}
