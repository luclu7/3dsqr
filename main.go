package main

import (
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	argsWithoutProg := os.Args[1:]
	fileInput := strings.Join(argsWithoutProg, " ")
	match, _ := regexp.MatchString(".cia$", fileInput)

	if match == false {
		fmt.Println("Please provide a CIA file.")
		os.Exit(1)
	}

	ip := GetOutboundIP()
	ipString := ip.String()
	content := "http://" + ipString + ":8000/" + fileInput
	obj := qrcodeTerminal.New()
	obj.Get([]byte(content)).Print()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("Serving your CIA at: " + content)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
