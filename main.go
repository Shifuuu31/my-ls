package main

import (
	"fmt"
	"log"

	"my-ls/source"
)

func main() {
	var ls source.MyLsFlag
	fmt.Printf("%+v\n", ls)
	if err := source.ParseArgs(&ls); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", ls)
	fmt.Printf("DIR:%+v\n", source.Dir)
	fmt.Printf("File:%+v\n", source.File)
}
