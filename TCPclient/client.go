package main

import (
    "net"
    "os"
    "fmt"
    "io/ioutil"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]

    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    //result, err := readFully(conn)
    result, err := ioutil.ReadAll(conn)
    checkError(err)

    fmt.Println(string(result))

    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	con, err := net.Dial("tcp", "webcode.me:80")
	checkErrors(err)

	req := "HEAD / HTTP/1.0\r\n\r\n"

	_, err = con.Write([]byte(req))
	checkErrors(err)

	res, err := ioutil.ReadAll(con)
	checkErrors(err)

	fmt.Println(string(res))
}

func checkErrors(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
 */
