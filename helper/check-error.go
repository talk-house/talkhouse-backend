package helper

import (
	"fmt"
	"log"
)

func CheckError(err error, message string) {
	if err != nil {
		fmt.Println(message)
		log.Fatal(err)
	}
}
