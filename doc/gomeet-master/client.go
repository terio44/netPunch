package main

import (
    "net"
    "net/http"
    "io/ioutil"
    "bufio"
    "fmt"
)

var goOut = "out"

func readSocket(conn net.Conn) {
    reader := bufio.NewReader(conn)
    for {
        s, _ := reader.ReadString('\n')
        fmt.Print(s)
    }
}

func writeSocket(conn net.Conn) {
    for {
        var message string
        _ , _ = fmt.Scanf("%s", &message)
        fmt.Fprintf(conn, message + "\n")
    }
}

func handleConnection(conn net.Conn) {
       go readSocket(conn)
       writeSocket(conn) 
}

func main() {
    var clientPairId string
    fmt.Print("Enter a unique ID for the client pair: ")
    _, _ = fmt.Scanf("%s", &clientPairId)
    response, _ := http.Get("http://localhost:8080/meet/" + clientPairId)
    address, _ := ioutil.ReadAll(response.Body)
    conn, _ := net.Dial("tcp", string(address))
    handleConnection(conn)
}
