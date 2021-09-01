package main

import "fmt"

type Work interface {
	doJob()
}

type Job struct{}

type Human struct {
	age  int
}

type ProxyJob struct {
	job   Job
	human *Human
}

func (h *Human) doJob() {
	fmt.Println("Do Job")
}

func (p *ProxyJob) doJob() {
	if p.human.age > 10 {
		p.human.doJob()
	} else {
		fmt.Println("To young to do job")
	}
}

func main() {
	p1 := &ProxyJob{Job{}, &Human{1}}
	p2 := &ProxyJob{Job{}, &Human{11}}
	p1.doJob()
	p2.doJob()
}