package main

import (
	"github.com/gauravtayal0/Go-DesignPattern/Observer"
)

func main() {
	weatherSubject := Observer.NewSubject()
	var weatherObserver Observer.Observer

	for i := 0; i < 10; i++ {
		weatherObserver = Observer.NewReporter(i)
		weatherSubject.RegisterObserver(weatherObserver)
	}

	weatherSubject.SetMeasurement(123)

	//lets remove the last observer
	weatherSubject.RemoveObserver(weatherObserver)
	weatherSubject.SetMeasurement(456)

}
