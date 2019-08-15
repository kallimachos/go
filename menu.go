package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	menu := cli.NewApp()
	menu.Name = "boom"
	menu.Version = "0.0.1"
	menu.Usage = "make an explosive entrance"
	menu.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	err := menu.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
