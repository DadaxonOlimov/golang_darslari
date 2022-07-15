// for loop haqida
package main

import (
	"fmt"
)

func main() {
	test()
} 
func testfor1() {
	for i :=19 ; i < 22; i++ {
		fmt.Println(i)

}
	for i :=19 ; i < 22; i++ {
		fmt.Println(i)
	}
}


func TESTfor() {
	n := 6
	i := 0
	for i< n {
		fmt.Println("hello DAddy Olimov")
		i++
	}
	fmt.Printf("the value of i is %d/n" , i)
	}


func TESTfor2() {
	// myarr :=  [3]string {"osmon" , "yer" ,"quyosh"}
	mymap := map[int]string{
		1 :"c++",
		2 : "go" ,
		3 : "python" ,
	}
	for key,value := range mymap {
		fmt.Println("key" , key , "value" , value)

	}
}


