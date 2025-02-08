package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Подключение к FTP-серверу
	conn, err := net.Dial("tcp", "localhost:2121")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Читаем приветственное сообщение от сервера
	go readServer(conn)

	// Читаем команды из стандартного ввода и отправляем на сервер
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Fprintln(conn, cmd)
		if strings.HasPrefix(cmd, "close") {
			break
		}
	}
}

func readServer(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		// Читаем данные от сервера и выводим в консоль
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error reading from server:", err)
			break
		}
		fmt.Print(msg)
	}
}
