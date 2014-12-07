package main
 
import "net"
import "io"
import "log"
 
func main() {
	ln, err := net.Listen("tcp", ":8080")
 
	if err != nil {
		// handle error
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
	bytesRead := 0
	buf := make([]byte, 81920)
 
	for {
		n, err := c.Read(buf)
		bytesRead += n
 
		if err != nil {
			if err != io.EOF {
				log.Printf("Read error: %s", err)
			}
			break
		}
 
		log.Println(n)
	}
 
	log.Printf("%d bytes read: %s", bytesRead, string(buf))
	c.Close()
}
