package helper

import (
	"fmt"
	"os"
)

func HandleError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	} 
}

func JsonResult(code string,message string) {
	fmt.Printf(`{"result":"%s","message":"%s"}`,code,message)
	fmt.Println("")
}