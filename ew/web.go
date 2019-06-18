package main

import (
	"net/http"
	"bytes"
	"flag"
	"fmt"
	"log"
	"io"
	"strconv"

)

//flags
var (
	addr = flag.String("a", "127.0.0.1", "-a options is Listen address")
	port = flag.Int("p", 8080, "-p options is Listen Port")
)

//logger
var (
	buf bytes.Buffer
	logger = log.New(&buf, "[error]", log.Lshortfile)
)

func main() {
	flag.Parse()
	fmt.Printf("v: %v	p: %v\n", *addr, &addr)
	fmt.Printf("v: %v		p: %v\n", *port, &port)
	fmt.Printf("Listen Start %v:%v\n", *addr,*port)

	handle := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Test Page\n")
	}

	http.HandleFunc("/", handle)
	http.ListenAndServe(*addr + ":" + strconv.Itoa(*port) , nil)
}

