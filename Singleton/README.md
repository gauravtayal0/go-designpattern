Singleton is one of the most common and simplest design patterns in terms of its class diagram. Infact diagram just holds a single class.

Wikipedia defines Singleton Pattern as:
In software engineering, the singleton pattern is a software design pattern that restricts the instantiation of a class to one "single" instance. This is useful when exactly one object is needed to coordinate actions across the system. The term comes from the mathematical concept of a singleton.

But don't get too comfortable; despite its simplicity from a design perspective, we are going to encounter few potholes in its implementation.


Critics consider the singleton to be an anti-pattern in that it is frequently used in scenarios where it is not beneficial, introduces unnecessary restrictions in situations where a sole instance of a class is not actually required, and introduces global state into an application


Singleton ensures that a class has only one instance and provides a global point of access to it. To translate that in Go, whenever any goroutine tries to access an instance of a variable, you should get the same variable.

There are various ways to implement the singleton pattern in Go, but it is also quite easy to get it wrong. The right way to implement a singleton pattern in Go is to use the sync package’s Once.Do() function. This function makes sure that your specified code is executed only once and never more than once.

singleton package with Once.Do() 

```golang
package singleton

import (
    "sync"
)

type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```
