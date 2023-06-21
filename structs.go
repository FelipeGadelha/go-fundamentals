// package main

// import "fmt"

// func main() {

// 	var user User = User{
// 		name: "Felipe",
// 		age:  30,
// 	}
// 	fmt.Println(user)
// 	fmt.Println(user.name)
// 	fmt.Println(user.age)

// 	var end = Endereco{
// 		rua:    "rua 1",
// 		cidade: "São Paulo",
// 		estado: "SP",
// 	}
// 	var pessoa = Pessoa{
// 		nome:     "Felipe",
// 		idade:    30,
// 		endereco: end,
// 	}

// 	fmt.Println(pessoa)
// 	fmt.Println(pessoa.endereco.cidade)
// }

// type User struct {
// 	name string
// 	age  int
// }

// type Pessoa struct {
// 	nome     string
// 	idade    int
// 	endereco Endereco
// }

// type Endereco struct {
// 	rua    string
// 	cidade string
// 	estado string
// }