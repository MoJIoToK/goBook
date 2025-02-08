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
	// Запускаем FTP-сервер, слушая соединения на порту 2121
	ln, err := net.Listen("tcp", ":2121")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("FTP Server started on port 2121")

	// Принимаем входящие соединения в бесконечном цикле
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// Обрабатываем каждое соединение в отдельной горутине
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	cwd, _ := os.Getwd()
	fmt.Fprintf(conn, "Connected to FTP server. Current directory: %s\n", cwd)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := strings.Fields(scanner.Text()) // Разбираем введённую команду
		if len(cmd) == 0 {
			continue
		}
		switch cmd[0] {
		case "cd":
			// Меняем рабочий каталог
			if len(cmd) < 2 {
				fmt.Fprintln(conn, "Usage: cd <directory>")
				continue
			}
			if err := os.Chdir(cmd[1]); err != nil {
				fmt.Fprintln(conn, "Failed to change directory:", err)
			} else {
				cwd, _ = os.Getwd()
				fmt.Fprintln(conn, "Changed directory to", cwd)
			}
		case "ls":
			// Выводим список файлов в текущем каталоге
			files, err := os.ReadDir(".")
			if err != nil {
				fmt.Fprintln(conn, "Error reading directory:", err)
				continue
			}
			for _, file := range files {
				fmt.Fprintln(conn, file.Name())
			}
		case "get":
			// Отправляем клиенту содержимое файла
			if len(cmd) < 2 {
				fmt.Fprintln(conn, "Usage: get <filename>")
				continue
			}
			file, err := os.Open(cmd[1])
			if err != nil {
				fmt.Fprintln(conn, "Error opening file:", err)
				continue
			}
			io.Copy(conn, file)
			file.Close()
		case "close":
			// Закрываем соединение
			fmt.Fprintln(conn, "Closing connection...")
			return
		default:
			// Сообщаем об неизвестной команде
			fmt.Fprintln(conn, "Unknown command")
		}
	}
}
