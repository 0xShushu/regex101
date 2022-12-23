package main

import (
	"net/http"
	"fmt"
	"github.com/0xshushu/regex101/server"
)

func main() {
	//make a new server
	s := server.NewServer()

	//mount handlers
	s.MountHandlers()

	//start serving
	fmt.Println(http.ListenAndServe(":3000", s.Router))
}