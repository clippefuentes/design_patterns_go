package main

import "sync"

var instance singleInstance

type singleInstance struct{}

var once sync.Once

func GetInstance() singleInstance {
	once.Do(func() {
		instance = singleInstance{}
	})
	return instance
}

