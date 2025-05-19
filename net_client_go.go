package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

type User struct {
	name       string
	index      int
	connection net.Conn
	err        error
}

func (user *User) init_conn_user(server string, port int) {

	fmt.Println("Input your user name")
	reader := bufio.NewReader(os.Stdin) // Создаем новый Reader
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		os.Exit(1)
	}
	user.name = name[:len(name)-1]
	server = server + ":" + strconv.Itoa(port)
	user.connection, user.err = net.Dial("tcp", server)
	if user.err != nil {
		defer user.connection.Close()
		fmt.Println("Something wrong")
		fmt.Println(user.err)
		os.Exit(1)
	}
}

func (user *User) send_message() {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}
		message = message[:len(message)-1]
		current_time := time.Now()
		message = user.name + "\t" + ": " + message + "\t || " + current_time.Format("02-01-2006 15-04-05"+"\n")
		user.connection.Write([]byte(message))
	}
}

func (user *User) update() {
	for {
		var buffer []byte = make([]byte, 1024*10)
		_, err := user.connection.Read(buffer)
		if err != nil {
			fmt.Println(err)
		}
		io.Copy(os.Stdout, user.connection)
	}
}

func main() {
	var user User
	user.init_conn_user("127.0.0.1", 8080)
	if user.err != nil {
		return
	}
	go user.send_message()
	user.update()
	select {}
}
