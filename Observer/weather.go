package Observer

type weather struct {
	temperature int64
	observers   map[Observer]int64 //[]Observer
}

func (this *weather) GetTemperature() int64 {
	return this.temperature
}

func (this *weather) SetMeasurement(temperatue int64) {
	this.temperature = temperatue
	this.dataChanges()
}

func (this *weather) dataChanges() {
	this.NotifyObservers()
}

func (this *weather) RegisterObserver(o Observer) {
	this.observers[o] = 1
}

func (this *weather) RemoveObserver(o Observer) {
	_, check := this.observers[o]
	if check {
		delete(this.observers, o)
	}
}

func (this *weather) NotifyObservers() {
	for observer := range this.observers {
		if observer != nil {
			observer.Update(this.GetTemperature())
		}
	}
}

func NewSubject() *weather {
	subject := new(weather)

	subject.observers = make(map[Observer]int64) //make([]Observer, 0)
	return subject
}
