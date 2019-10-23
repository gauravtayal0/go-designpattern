package Observer

type (
	Subject interface {
		Subscribe(Observer)
		GetState() int64
		SetState(int64)
		Notify()
	}

	Observer interface {
		Update(int64)
	}
)
