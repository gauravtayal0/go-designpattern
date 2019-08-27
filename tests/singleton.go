package main

import (
	"github.com/gauravtayal0/Go-DesignPattern/Singleton"
	"fmt"
)

func main(){
	instance := Singleton.Getinstance()
	fmt.Printf("gaurav instance is %v \n", instance)

	instance.Test()
}
