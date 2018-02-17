package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// open and listen on tcp port
	listener, err := net.Listen("tcp", ":1234")
	log.Println("Service started, listening on port 1234...")
	defer listener.Close()
	if err != nil {
		log.Fatalln(err)
	}

	commands := make(chan Command)
	// start the redis-imp server
	go commandServer(commands)

	// accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		// handle commands sent by user
		go processCmd(commands, conn)
	}

}

func processCmd(c chan Command, conn net.Conn) {
	// close connection at the end
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		response := make(chan string)
		c <- Command{
			Fields:   fs,
			Response: response,
		}

		fmt.Fprintf(conn, <-response+"\n")

	}
}

func commandServer(commands chan Command) {
	// create our redis(ish) kind of map to store values
	var data = make(map[string]string)
	for command := range commands {
		if len(command.Fields) < 2 {
			noBlockWrite(command.Response, "at least 2 arguments is needed!")
			continue
		}

		switch command.Fields[0] {
		case "GET": // GET {KEY}
			key := command.Fields[1]
			value := data[key]
			noBlockWrite(command.Response, "...->"+value)
		case "SET": // SET NEW {KEY: VALUE}
			if len(command.Fields) != 3 {
				noBlockWrite(command.Response, "expected value")
				continue
			}
			key := command.Fields[1]
			value := command.Fields[2]
			data[key] = value
			noBlockWrite(command.Response, "")
		case "DEL": // DELETE {KEY}
			key := command.Fields[1]
			delete(data, key)
			noBlockWrite(command.Response, "")
		default:
			noBlockWrite(command.Response, "INVALID COMMAND "+command.Fields[0])
		}
	}
}

// write to command channel without blocking
func noBlockWrite(c chan string, msg string) {
	go func() {
		c <- msg
	}()
}
