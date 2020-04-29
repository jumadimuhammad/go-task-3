package main

import (
	"fmt"
	"strings"
)

func main() {
	temp := 0
	result := []string{}
	arr := [...]string{"Mie ayam", "Nasi goreng", "Pecel lele", "Soto daging"}
	key := "Nasi Goreng"

	for _, value := range arr {
		if strings.ToLower(value) == strings.ToLower(key) {
			temp++
			result = append(result, value)
		}
	}

	fmt.Println(temp, result)

	status, ok := Find(temp)
	fmt.Println(ok)
	if !ok {
		fmt.Println("Result :", result[0], ", Status", status)
	} else {
		fmt.Println(status)
	}
}

// Find is
func Find(value int) (notes string, status bool) {
	if value < 1 {
		notes = "Not found"
		status = true
	} else {
		notes = "Found"
		status = false
	}
	return
}
