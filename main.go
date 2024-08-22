package main

import (
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
	}
}
