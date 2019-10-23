package Observer

type weather struct {
	state     int64
	observers []Observer
}

func (w *weather) GetState() int64 {
	return w.state
}

func (w *weather) SetState(temperatue int64) {
	w.state = temperatue
	w.Notify()
}

func (w *weather) Subscribe(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *weather) Notify() {
	for _, o := range w.observers {
		if o != nil {
			o.Update(w.GetState())
		}
	}
}

func NewSubject() Subject {
	subject := new(weather)

	subject.observers = make([]Observer, 0)
	subject.state = 0

	return subject
}
