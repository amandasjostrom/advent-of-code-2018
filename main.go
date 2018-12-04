package main

import (
	"amandasjostrom.se/day1"
	"amandasjostrom.se/day2"
	"amandasjostrom.se/day3"
	"amandasjostrom.se/common"
)

func main(){
	//runDay1()
	//runDay2()
	runDay3()
}

func runDay1() {
	day1.Run()
}

func runDay2() {
	day2.Run(common.GetInput(2, "\n"))
}

func runDay3() {
	day3.Run(common.GetInput(3, "\n"))
}