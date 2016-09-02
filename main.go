package main

import (
	"fmt"
	"os"
	"time"
)

// base time layout
const timeLayout = "2006-01-02 15:04:05"

func main() {
	args := os.Args
	var now time.Time

	// first Argument is time
	if len(args) > 1 {
		t, err := time.Parse(timeLayout, args[1])
		if err != nil {
			panic("fail to parse time you give : " + args[1])
		}
		now = t
	} else {
		now = time.Now()
	}
	fmt.Println(now.Unix())
}
