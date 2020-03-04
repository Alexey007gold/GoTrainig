package main

import (
	"fmt"
	"io/ioutil"
)
import "encoding/json"

var a = 4

type data struct {
	Str string
	I   int
}

func main() {
	//a = 8
	//fmt.Println(a)
	//println("SS")
	//got := 7
	//println(got)
	//
	//array_name := [4]string{"s"}
	//array_name[3] = "32"
	//var arr[1]string
	//fmt.Println(len(arr))
	//println(array_name[3])

	dataA := data{"afsa", 3}
	_, err := json.Marshal(dataA)
	if err == nil {
		//fmt.Println(marshal)
	}

	file, _ := ioutil.ReadFile("test.json")
	dat := data{}
	err = json.Unmarshal(file, &dat)
	if err == nil {
		fmt.Println(dat)
	} else {
		fmt.Println(err)
	}
}

func b() {

}
