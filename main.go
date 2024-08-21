package main

import (
	"fmt"
	"main/src/engine"
)

func main() {
	fmt.Println("starting...")
	var e engine.Engine
	
	e.Init()
	e.Load()
	e.Run()
	e.Unload()
	e.Close()
}
