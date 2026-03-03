package main

import "fmt"



type User struct {
	Name string
	Age  int
}


func (u User) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", u.Name, u.Age)
}


func main() {
	user := User{Name: "Alice", Age: 30}
	fmt.Println(user.Greet())
}