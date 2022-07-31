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
	fmt.Printf("U Tranform object : %v\n", rubik.U)
	r.MoveUpDirect()
	fmt.Printf("%v\n", r)
	r.Display(os.Stdout)

	t := rubik.InitialTransform(*r)
	fmt.Printf("Tranform object : %v\n", t)
	t = t.Reverse()
	fmt.Printf("Reverse tranform object : %v\n", t)
	newRubik := t.Apply(*r)
	fmt.Printf("%v\n", newRubik)
	newRubik.Display(os.Stdout)
	fmt.Println("===== Re-apply transform =======")
	newRubik = t.Apply(newRubik)
	fmt.Printf("%v\n", newRubik)
	newRubik.Display(os.Stdout)

}
