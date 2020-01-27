package main

import (
	"fmt"
	"github.com/gauravtayal0/Go-DesignPattern/Factory"
)

func main() {
	database := Factory.DatabaseFactory(Factory.Sql)

	database.PutData("test", "database selected is")
	fmt.Println(database.GetData("test"))
}
