package main

import (
	"fmt"
	"github.com/sqjian/bingo/example/echo/cmd"
)

var (
	_os   string
	_date string
	_ver  string
)

const dog = `
   / \__
  (    @\___
  /         O
 /   (_____/
/_____/   U
`

func logo() {
	fmt.Println("----------------------------------------")
	fmt.Printf("%s\n", dog)
	fmt.Printf("born in %v on the %v,ver:%v\n", _os, _date, _ver)
	fmt.Println("----------------------------------------")
}
func init() {
	logo()
}

func main() {
	cmd.Execute()
}
