package main

import (
	"fmt"
	"github.com/gauravtayal0/Go-DesignPattern/Singleton"
)

func main(){
	instance := Singleton.Getinstance()
	fmt.Printf("gaurav instance is %v \n", instance)

	instance.Test()
}
