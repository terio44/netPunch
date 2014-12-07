package main
 
import (
		"net"
		"log"
		"encoding/json"
		"errors"
		"os"
)

type msg struct {
    Name string
    IP string
    Time int64
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
	}

	m := msg{"Client", "192.0.0.1", 4444}
	m_json,err := json.Marshal(m)
	if err != nil {
		// handle error
	}

	_, err = conn.Write([]byte(m_json))

    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

	log.Printf("Conn:", conn)
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("Not connected to the network")
}