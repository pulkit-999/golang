package main

import (
	"fmt"
	"sync"
)

type singleton struct{} // singleton is private else Singleton becomes public with uppercase

var instance *singleton
var once sync.Once

func Getinstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
func main() {
	s1 := Getinstance()
	s2 := Getinstance()
	if s1 == s2 {
		fmt.Println("same instance")
	}
}
