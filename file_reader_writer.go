package main

import (
	"fmt"
	"io"
)

type phoneReader struct {
	phone string
}

func (phone phoneReader) Read(buffer []byte) (int, error) {
	var bytes int = 0
	for i := 0; i < len(phone.phone); i++ {
		if phone.phone[i] >= '0' && phone.phone[i] <= '9' {
			buffer[bytes] = phone.phone[i]
			bytes++
		}
	}
	return bytes, io.EOF
}

func main() {
	var phone1 phoneReader = phoneReader{"+1(234)567 9010"}
	var phone2 phoneReader = phoneReader{"+2-345-678-12-35"}

	var buffer []byte = make([]byte, len(phone1.phone))
	phone1.Read(buffer)
	fmt.Println(string(buffer))

	phone2.Read(buffer)
	fmt.Println(string(buffer))
}
