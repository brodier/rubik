package main

import (
	"fmt"
	"os"

	"github.com/brodier/rubik/rubik"
)

func main() {
	file, err := os.Open("rubik.txt")
	if err != nil {
		fmt.Println("Unable to open file")
	}
	r := rubik.NewRubik(file)
	fmt.Printf("%v\n", r)
	r.Display(os.Stdout)
	fmt.Println("===== Now move up =======")
	r.MoveUpDirect()
	fmt.Printf("%v\n", r)
	r.Display(os.Stdout)
}
