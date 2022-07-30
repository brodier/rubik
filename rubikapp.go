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

	t := rubik.InitialTransform(*r).Reverse()
	newRubik := t.Apply(*r)
	fmt.Printf("%v\n", newRubik)
	newRubik.Display(os.Stdout)
	fmt.Println("===== Re-apply transform =======")
	newRubik = t.Apply(newRubik)
	fmt.Printf("%v\n", newRubik)
	r.Display(os.Stdout)

}
