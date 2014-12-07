package main
 
import "net"
import "io"
import "log"
import "encoding/json"

type msg struct {
    Name string
    IP string
    Time int64
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
 
	if err != nil {
		log.Println("Cannot instantiate the listener, verify the port")
	}
 
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Cannot open listener.")
			continue
		}
		go handleConnection(conn)
	}
 
}
 
func handleConnection(c net.Conn) {
	buf := make([]byte, 255)
 	var m msg
	n, err := c.Read(buf)

	if err != nil {
		if err != io.EOF {
			log.Printf("Read error: %s", err)
		}
	}

	err = json.Unmarshal(buf[:n], &m)

	if err != nil {
	log.Printf("Read error: %s", err)
	c.Close()
	}

	log.Printf("Name:%s,IP:%s,Time:%d",m.Name,m.IP,m.Time)
 
	log.Printf("%d bytes read: %s", n, string(buf))
	c.Close()
}
