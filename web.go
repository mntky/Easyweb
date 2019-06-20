package main

import (
	"net/http"
	"bytes"
	"flag"
	"fmt"
	"log"
	"io"
	"strconv"
	"regexp"
	"os"

)

//flags
var (
	addr = flag.String("a", "127.0.0.1", "-a options is Listen address")
	port = flag.Int("p", 8080, "-p options is Listen Port")
	args = flag.String("t", "Test Page", "-t options is display words")
)

//logger
var (
	errbuf bytes.Buffer
	errlog = log.New(&errbuf, "[error]", 0)
	//log.New(io.write, string,	0 -- nothing
	//							1 -- yyyy/mm/dd
	//							2 -- hh:mm:ss
	//							3 -- day and time
	//							4 -- print the 6th decimal place
	//							5 -- day and [4]
)

func main() {
	flag.Parse()
	chk := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	tf := chk.MatchString(*addr)
	if tf == false {
		errlog.Print("mistaken ip address")
		fmt.Println(&errbuf)
		os.Exit(0)
	}

	fmt.Printf("Listen Start %v:%v\n", *addr,*port)
	start()
}

func start() {
	handle := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, *args)
	}
	http.HandleFunc("/", handle)
	http.ListenAndServe(*addr + ":" + strconv.Itoa(*port) ,nil)
}
