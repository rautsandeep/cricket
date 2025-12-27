package main

import (
	"fmt"
)

type mango struct {
	test string
}

func (m mango) juice() {
	fmt.Println("Test of the mango juice is", m.test)
}

func (m *mango) sweet() {
	fmt.Println("test of the sweet prepared from the mango is", m.test)
}

type fruit interface{
juice()
sweet()
}

func display(f fruit){
 f.juice()
}


func main() {

	m1 := mango{
		test: "rasal",
	}
	m2:= &mango{
		test:"kharat",
	}
	m1.juice()
	m1.sweet()

	fmt.Println("-----------------")

	m2.juice()
	m2.sweet()
	fmt.Println("-----------------")
	display(&m1)
}
