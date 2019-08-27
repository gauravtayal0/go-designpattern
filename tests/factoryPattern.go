package main

import (
	"fmt"
	"github.com/gauravtayal0/Go-DesignPattern/Factory"
)

func main(){
	database := Factory.DatabaseFactory("production")
	
	database.PutData("test", "this is mongodb")
	fmt.Println(database.GetData("test"))
}
