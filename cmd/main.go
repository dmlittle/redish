package main

import (
	"fmt"
	"os"

	"github.com/dmlittle/redish/pkg/cmd"
)

func main() {
	if err := cmd.NewRedishCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
