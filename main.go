package main

import (
	"github.com/go-stomp/stomp/v3"
	"log"
)

func ExampleDial_1() error {
	conn, err := stomp.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}

	err = conn.Send(
		"/queue/test-1",           // destination
		"text/plain",              // content-type
		[]byte("Test message #1")) // body
	if err != nil {
		return err
	}

	return conn.Disconnect()
}
func main() {
	log.Println(ExampleDial_1())
	//netConn, err := net.DialTimeout("tcp", "127.0.0.1:8080", 10*time.Second)
	//if err != nil {
	//	log.Fatalf("Dial: %s", err)
	//}
	//
	//stompConn, err := stomp.Connect(netConn)
	//if err != nil {
	//	log.Printf("Connect: %s", err)
	//}
	//stompConn.Send("/", "text/plain", []byte("Hi!"))
	//defer stompConn.Disconnect()

}
