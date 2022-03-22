package main

import "fmt"

func main() {

	str1 := "Geeks"
	str2 := "Geek"
	str3 := "GeeksforGeeks"
	str4 := "Geeks"


	result1 := str1 == str2
	result2 := str2 == str3
	result3 := str3 == str4
	result4 := str1 == str4

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
	fmt.Println("Result 3: ", result3)
	fmt.Println("Result 4: ", result4)


	result5 := str1 != str2
	result6 := str2 != str3
	result7 := str3 != str4
	result8 := str1 != str4

	fmt.Println("\nResult 5: ", result5)
	fmt.Println("Result 6: ", result6)
	fmt.Println("Result 7: ", result7)
	fmt.Println("Result 8: ", result8)

}

