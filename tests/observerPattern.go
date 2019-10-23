package main

import (
	"fmt"
	"github.com/gauravtayal0/Go-DesignPattern/Observer"
)

func main() {
	weatherSubject := Observer.NewSubject()
	fmt.Println("hello", weatherSubject)

	for i := 0; i < 10; i++ {
		reporter1 := Observer.NewReporter(i)
		weatherSubject.Subscribe(reporter1)
	}

	weatherSubject.SetState(123)

	//weatherSubject.Notify()
}
