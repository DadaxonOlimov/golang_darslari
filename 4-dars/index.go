package main

import (
	"fmt"
	"sort"
)

func main() {
	TestSlises3()
}
func testArrays1() {
	var myarr [3]string
	myarr[0] = "Go"
	myarr[1] = "dasturlash"
	myarr[2] = "vertualdars"
	fmt.Println("Qatorning elementlari:")
	fmt.Println("1: ",myarr[0])
	fmt.Println("2: ",myarr[1])
	fmt.Println("3: ",myarr[2])

}

func testArrays2() {
	myarr := [3] int{2,3,4}
	fmt.Println(myarr)
}

func testArrays3() {
	myarr1 := [...] int {1,3,5,6,}
	myarr2 := [4] int {1,3,5,6}
	fmt.Println(myarr1 == myarr2)
}

func testArrays4() {
	// ko'p o'lchamli qatorlar:
	myarr := [3][3]string {{"Java" ,"Go" ,"Python"},
	{"TypeSprict" ,"C++" ,"koltin"} ,
	{"Dart" ,"C#" ,"Swith"}}
	fmt.Println(myarr[1][2])
	}
 
func testArrays5() {
	myarr1 :=[3] int {1,2,5}
	// qatordan to'liq nusxa olish
	myarr2 := myarr1
	fmt.Println(myarr1)
	fmt.Println(myarr2)

	myarr1[2] = 100
	fmt.Println(myarr1)
	fmt.Println(myarr2)
}
 
func testArrays6() {
	myarr1 := [3] int{2,4,7}
	// qatordan reference 'li nusxa olish:
	myarr2 := &myarr1
	fmt.Println(myarr1)
	fmt.Println(*myarr2)

}

func testArrays9() {
	myarr := [2]int {2,5}
	myarr1 := &myarr
	myarr[0] = 1000 
	fmt.Println(myarr)
	fmt.Println(*myarr1)

	}
	
func TestSlises(){
	slises := []int {1,2,4,5,6,7,8,9}
	slises = append(slises, 10)
	fmt.Println(slises)
	fmt.Printf("bunda slises ichidagi elementlar soni sanaydi 9ta ekan : %d", len(slises))
}

func TestSlises2 () {
	myarr := [7]string {"anor","olma","shaftoli","gilos",
"o'rik","uzum", "anjir"}
// slice'ni qatordan yaratib olish 
slises := myarr[1:4]
fmt.Println(slises)

}

func TestSlises3(){
	myslice := []int {212,31,3,5,7,5656,67,77}
	sort.Ints(myslice)
	fmt.Println(myslice)

}




