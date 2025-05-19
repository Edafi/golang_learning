package main

/*
type Reader interface {
    Read(p []byte) (n int, err error)
}
*/

/*
type Writer interface {
    Write(p []byte) (n int, err error)
}
*/

import (
	"fmt"
	"io"
	"os"
)

type Phone struct {
	phone_number string
}

func (phone Phone) Read(buffer []byte) (int, error) {
	var index int = 0
	for i := 0; i < len(phone.phone_number); i++ {
		if phone.phone_number[i] >= '0' && phone.phone_number[i] <= '9' {
			buffer[index] = phone.phone_number[i]
			index++
		}
	}
	return index, io.EOF
}

func (phone Phone) Write(buffer []byte) (int, error) {
	if len(buffer) == 0 {
		return 0, io.EOF
	} else {
		file, err := os.OpenFile("GoLang_Print.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Error opening or creating file")
			os.Exit(1)
		}

		buffer = append(buffer, '\n')
		for i := 0; i < len(buffer); i++ {
			if buffer[i] >= '0' && buffer[i] <= '9' || buffer[i] == '\n' {
				file.Write(buffer[i : i+1])
			}
		}
		defer file.Close()
	}
	return len(buffer), nil
}

func main() {
	var phone1 Phone = Phone{"+1(234)567 9010"}
	var phone2 Phone = Phone{"+2-345-678-12-35"}

	var buffer []byte = make([]byte, len(phone1.phone_number))
	phone1.Read(buffer)
	//fmt.Println(string(buffer))
	phone1.Write(buffer)

	buffer = make([]byte, len(phone2.phone_number))
	phone2.Read(buffer)
	//fmt.Println(string(buffer))
	phone2.Write(buffer)

	file, err := os.OpenFile("GoLang_Print.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Something wrong! os.Open can't open the file")
		os.Exit(1)
	}
	fmt.Fprintf(file, "%-10s %-10d %-10.3f\n", "Tom", 21, 78.1545)
	file.Seek(0, 0)
	io.Copy(os.Stdout, file)
	defer file.Close()
}
