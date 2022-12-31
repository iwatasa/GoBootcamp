package second

import (
	"fmt"
	"strconv"
)

type User struct {
	name string
	age int
}

func ConvertUser(name string, age int) User {
	return User{
		name: name,
		age: age,
	}
}

func PrintlnUser(user User) bool {
	fmt.Println("User name: "+user.name+", Age: "+strconv.Itoa(user.age))
	return true
}