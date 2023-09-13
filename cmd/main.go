package main

import (
	"github.com/Pedroxer/CardChecker/util"
	"github.com/Pedroxer/CardChecker/web"
	"log"
)

func main() {
	conf, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}

	server := web.NewServer(conf)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}

}
