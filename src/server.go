package main
 
import "net"
import "io"
import "log"
import "encoding/json"
//import "strings"

type msg struct {
    Name string
    IP string
    Time int64
}

const (  
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
	ln, err := net.Listen(CONN_TYPE, ":"+ CONN_PORT)
 
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
	
	//remoteAddr := c.RemoteAddr()
	log.Printf("Network:%s",c.RemoteAddr())
	//s := strings.Split(remoteAddr, ":")
	//ip, port := s[0], s[1]
	//log.Printf("IP:%s , PORT:%s", ip, port)

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

func findExternalClientIP (c net.Conn) {

}

func getClientAddress (c net.Conn) {

}
