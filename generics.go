// package main

// import "fmt"

// type UserTest1 struct {
// 	name string
// 	age  int
// }

// type UserTest2 struct {
// 	documentNumber string
// 	address        string
// }

// func userTest[T UserTest1 | UserTest2](user T) {
// 	fmt.Println(user)
// }

// func identity[T any](arg T) T {
// 	return arg
// }

// func anyType(arg any) any {
// 	return arg
// }

// func main() {
// 	outIdentity1 := identity("test")
// 	outIdentity2 := identity(1234)

// 	out1 := anyType(1234)
// 	out2 := anyType("test")

// 	fmt.Println(outIdentity1)
// 	fmt.Println(outIdentity2)
// 	fmt.Println(out1)
// 	fmt.Println(out2)
// }
