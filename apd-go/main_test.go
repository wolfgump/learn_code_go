package main

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name string
}

func TestAddress(t *testing.T) {
	user := &User{"zs"}
	fmt.Println(user)
	fmt.Println(reflect.TypeOf(user))
	fmt.Println(*user)
	fmt.Println(reflect.TypeOf(*user))
	fmt.Println(&user)
	fmt.Println(reflect.TypeOf(&user))
}
func TestDefault(t *testing.T)  {
	var id int
	fmt.Println(id)
}
