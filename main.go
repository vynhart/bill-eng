package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide command [http,work,migrate]")
	}

	switch os.Args[1] {
	case "migrate":
		migrateDB(GetConfig())
		return
	case "http":
		fmt.Println("to run http server, not implemented yet")
		// Initiate http handler object from http sub dir
		// compose all necessary dependency such as billing implementation from
		// package bill and loan service from package loan.
	case "work":
		fmt.Println("to run background job worker, not implemented yet")
		// Same as http, initiate worker object from worker sub dir
		// compose all necessary dependency and run the observer.
	}
}
