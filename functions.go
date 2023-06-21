// package main

// import "fmt"

// func main() {
// 	// value, err := test()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println(value)

// 	// funcTest := func(test string, testInt int) {
// 	// 	fmt.Println(test, testInt)
// 	// }

// 	// value, err := testFunc(funcTest)

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// fmt.Println(value)

// 	// function := test()
// 	// function("Felipe G", 31)

// 	// test()

// 	// test("Felipe", "Test", "For", "stringArg")

// 	testFunc := func() {
// 		fmt.Println("Felipe G")
// 	}

// 	test(testFunc)
// }

// // multiple-return

// // func test() (string, error) {
// // 	return "Test", errors.New("Error")
// // }

// // named-return
// // func testFunc(value func(string, int)) (returnString string, returnError error) {
// // 	value("Felipe", 30)
// // 	returnString = "test"
// // 	returnError = nil
// // 	return
// // }

// // func test() func(string, int) {
// // 	funcTest := func(valueString string, valueInt int) {
// // 		fmt.Println(valueString, valueInt)
// // 	}
// // 	return funcTest
// // }

// // func test() {
// // 	func(valueString string, valueInt int) {
// // 		fmt.Println(valueString, valueInt)
// // 	}("Felipe G", 20)
// // }

// //	func test(valueString ...string) {
// //		for _, x := range valueString {
// //			fmt.Println(x)
// //		}
// //	}
// func test(valueString ...func()) {
// 	for _, x := range valueString {
// 		x()
// 	}
// }
