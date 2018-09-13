// ReadAll useage
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

const filePath string = "/home/ming/scripts/MyProject/golang/Readme"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file error:", err)
	}
	data, err := ioutil.ReadAll(file)
	fmt.Println(reflect.TypeOf(filePath))
	fmt.Println(string(data))
	defer file.Close()
}
