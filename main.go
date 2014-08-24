// main.go
// Graphite data visualization tool.
// Copyright © 2014 Brigham Toskin, Christopher R. Tooley
// MIT-licensed; see the file LICENSE for details.

// Graphite is a data visualization and graphing utility
package main

import (
	"fmt"
	"log"
	"os"
)

const LogLimit = 5e6

func main() {
	tLog, tErr := os.Open("testdata/log")
	if tErr != nil {
		log.Fatal(tErr)
	}
	defer func() { if tLog != nil { tLog.Close() } }()
	
	tStat, tErr := tLog.Stat()
	if tErr != nil {
		log.Fatal(tErr)
	}
	
	if tLogSize := tStat.Size(); tLogSize < LogLimit {
		fmt.Println("log size: ", tLogSize)
	} else {
		fmt.Println("too big!", tLogSize)
	}
}