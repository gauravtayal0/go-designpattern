package Observer

import (
	"fmt"
)

type reporter struct {
	Id int
}

func (r *reporter) Update(message int64) {
	fmt.Printf("%d subject state received for observer id %d \n", message, r.Id)
}

func NewReporter(id int) Observer {
	observer := new(reporter)
	observer.Id = id

	return observer
}
