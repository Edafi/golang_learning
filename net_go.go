package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
)

type Server struct {
	clients map[net.Conn]bool
	mutex   sync.Mutex
}

func (server *Server) add_client(connection net.Conn) {
	server.mutex.Lock()
	server.clients[connection] = true
	defer server.mutex.Unlock()
}

func (server *Server) delete_client(connection net.Conn) {
	server.mutex.Lock()
	delete(server.clients, connection)
	connection.Close()
	server.mutex.Unlock()
}

func (server *Server) write_message(message []byte, sender net.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for client := range server.clients {
		if client != sender {
			_, err := client.Write(message)
			if err != nil {
				fmt.Println("Error sending message to client:", err)
				client.Close()
				server.delete_client(client)
			}
		}
	}
}

func http_request(domain string, port int) {

	var http_request = "GET / HTTP/1.1\r\n" +
		"Host: " +
		domain +
		"\r\n" +
		"User-Agent: Go-http-client/1.1\r\n" +
		"Accept: */*\r\n" +
		"Connection: close\r\n\r\n"
	var connection, error = tls.Dial("tcp", domain+":"+strconv.Itoa(port), &tls.Config{})
	if error != nil {
		fmt.Println("Connection not established")
		return
	}
	defer connection.Close()
	if _, error = connection.Write([]byte(http_request)); error != nil {
		fmt.Println(error)
		return
	}
	io.Copy(os.Stdout, connection)
}

func server_listen() {
	var server Server
	server.clients = make(map[net.Conn]bool)
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer listener.Close()
		fmt.Println("Server starts working!\n Listening ...")
		for {
			connection, err := listener.Accept()
			defer connection.Close()
			if err != nil {
				fmt.Println(err)
				connection.Close()
				continue
			}
			server.add_client(connection)
			go handler_connections(&server, connection)
		}
	}
}

func handler_connections(server *Server, connection net.Conn) {
	input := make([]byte, (1024 * 10))
	for {
		n, err := connection.Read(input)
		input = append(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Println(string(input[0:]))
		server.write_message(input, connection)
	}
}

func main() {
	server_listen()
}
