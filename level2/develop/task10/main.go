package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// server
	ln, err := net.Listen("tcp", "localhost:8008")
	if err != nil {
		fmt.Println("Cannot start server")
	}

	conn, _ := ln.Accept()

	go func() {
		for {
			mes, err := bufio.NewReader(conn).ReadString('\n')
			if err == io.EOF {
				return
			}

			fmt.Print("Server: ", mes)

			mes = "socket " + mes

			conn.Write([]byte(mes + "\n"))
		}
	}()

	// client
	timeOut := flag.Int("timeout", 10, "Time out flag")
	flag.Parse()
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments.\nUSE : -timeout <timeout> <host> <port>")
		return
	}

	conn2, err := net.DialTimeout("tcp", os.Args[3]+":"+os.Args[4], time.Duration(*timeOut)*time.Second)
	if err != nil {
		time.After(time.Duration(*timeOut) * time.Second)
		fmt.Println("Wrong server ip")
		return
	}

	if conn2 != nil {
		defer conn.Close()
		fmt.Println("Client Opened")
	}

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				conn.Close()
			}

			fmt.Fprint(conn, text+"\n")
		}
	}()

	for {
		mes, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("Client: " + mes)
	}
}
