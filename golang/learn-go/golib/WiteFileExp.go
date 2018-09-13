// WiteFileExp
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

func main() {
	datas := "Hello Test!\nTest One\nTest Two\nTest Three\n"
	fmt.Println(reflect.TypeOf(datas))
	fmt.Println(datas)

	err := ioutil.WriteFile("testfile", []byte(datas), 0644)
	if err != nil {
		log.Fatal(err)
	}
	txts, _ := ioutil.ReadFile("testfile")
	fmt.Println(string(txts))

}
