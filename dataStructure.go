// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	// if 1 > 2 {
// 	// 	fmt.Println("True")
// 	// } else {
// 	// 	fmt.Println("False")
// 	// }

// 	// if 1 > 2 {
// 	// 	fmt.Println("True")
// 	// } else if 3 > 1 {
// 	// 	fmt.Println("False")
// 	// } else {
// 	// 	fmt.Println("Else")
// 	// }

// 	if result, err := test(); err != nil {
// 		fmt.Println("Result " + result)
// 		fmt.Println("Error " + err.Error())
// 	}
// 	fmt.Println("True")

// 	test := "felipe"

// 	switch test {
// 	case "test", "test3":
// 		fmt.Println("Caiu na primeira condição")
// 		fallthrough

// 	case "test2":
// 		fmt.Println("Caiu na segunda condição")
// 		fallthrough
// 	case "test4":
// 		fmt.Println("Caiu na segunda condição")
// 		break
// 	default:
// 		fmt.Println("Caiu no Default")
// 	}
// }

// func test() (string, error) {
// 	return "test", errors.New("Error test")
// }
