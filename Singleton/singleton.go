package Singleton

import ("sync"
		"fmt")

type singleton struct {}

var instance *singleton
var once sync.Once

func Getinstance() *singleton {
	once.Do(func(){
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) Test(){
	fmt.Println("hello from the inside of singleton")
}