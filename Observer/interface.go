package Observer

type (
	Subject interface {
		RegisterObserver(Observer)
		RemoveObserver(Observer)
		NotifyObservers()
	}

	Observer interface {
		Update(int64)
	}
)
