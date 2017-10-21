package main

import "fmt"

type Target interface {
	sayHello()
}

type Adaptee interface {
	printMessage(msg string)
}

type AdapteeObject struct {
	
}

func (adapteeObject1 AdapteeObject) printMessage(msg string) {
	fmt.Println(msg)	
}


type TargetObject struct {
	adaptee AdapteeObject 
}

func (targetObject1 TargetObject) sayHello() {
	adaptee := new(AdapteeObject)
	adaptee.printMessage("Hello")	
}

func Test(){
		obj := new(TargetObject) 
		obj.sayHello()
}

func main(){
	Test()
}