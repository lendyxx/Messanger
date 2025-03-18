package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("SERVER: ")

	listener, err := net.Listen("tcp", ":9988")
	log.Println("listening on :9988")

	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Connected ", conn.RemoteAddr())

		_, err = io.Copy(os.Stdout, conn)
		log.Println("disconnect with error: ", err)
	}
}
