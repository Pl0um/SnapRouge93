package main

import (
	"fmt"
	"flag"
	"log"
	"main/src/engine"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	var e engine.Engine
	fullscreen := flag.Bool("fullscreen", false, "Active le mode plein écran")

	flag.Parse()

	if *fullscreen {
		fmt.Println("Le mode plein écran est activé")

	}

	e.Init()
	e.Load()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	e.Run()
	e.Unload()
	e.Close()
}
