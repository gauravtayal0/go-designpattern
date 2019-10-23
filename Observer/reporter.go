package Observer

import (
	"fmt"
)

type reporter struct {
	Id int
}

func (r *reporter) Update(message int64) {
	fmt.Printf("%d state receiver for id %d \n", message, r.Id)
}

func NewReporter(id int) Observer {
	observer := new(reporter)
	observer.Id = id

	return observer
}
