// for loop haqida
package main

import (
	"fmt"
	"strconv"
)

func main() {
	TESTfor2()
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
	i := 1
	for i < 10 {
		fmt.Println("Alhamdulillah " + strconv. Itoa( i))
	i++
	}
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