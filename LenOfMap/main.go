package main

import (
	"fmt"
	"sync"
)

func Len(sm sync.Map) int {
	lengh := 0
	f := func(key, value interface{}) bool {
		lengh++
		return true
	}
	one:=lengh
	lengh=0
	sm.Range(f)
	if one != lengh {
		one = lengh
		lengh=0
		sm.Range(f)
		if one <lengh {
			return lengh
		}

	}
	return one
}

func main()  {

	a := sync.Map{}

	a.Store("name","zzt")
	a.Store("email","dishitin@qq.com")
	a.Store("age",14)
	a.Store("age",15)

	a.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	fmt.Println(Len(a))

}