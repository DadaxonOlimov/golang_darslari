// map malimotlari tuzilishi haqida
// map[key tab] value tab
package main

import "fmt"

func main() {
	 statuses := make(map[string]int)

	// mapga qiymatlar qo'shadi
	statuses["pending"] = 0
	statuses["ineted"] = 1
	statuses["running"] = 2
	statuses["timedout"] = 3
	statuses["succussful"] = 4
	statuses["failed"] = 5
	
	// mapdan qiymat o'chirish
	 succussfulStatus := statuses["failed"]
	fmt.Println(succussfulStatus)
	// mapdan bitta element o'chirish
	delete(statuses, "pending")
}