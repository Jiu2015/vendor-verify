package main

import (
	"fmt"
	"os"

	"github.com/Jiu2015/vendor-verify/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("verify successfully")
}
