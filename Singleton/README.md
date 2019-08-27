### Singleton is one of the most common and simplest design patterns in terms of its class diagram. Infact diagram just holds a single class.

# Wikipedia defines Singleton Pattern as:

In software engineering, the singleton pattern is a software design pattern that restricts the instantiation of a class to one *single instance*. This is useful when exactly one object is needed to coordinate actions across the system. 

The term comes from the mathematical concept of a singleton.


### Implementation

> There are various ways to implement the singleton pattern in Go, but it is also quite easy to get it wrong. The right way to implement a singleton pattern in Go is to use the sync package’s Once.Do() function. This function makes sure that your specified code is executed only once and never more than once.

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
