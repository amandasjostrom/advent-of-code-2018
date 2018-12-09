package main

import (
	"amandasjostrom.se/day1"
	"amandasjostrom.se/day2"
	"amandasjostrom.se/day3"
	"amandasjostrom.se/common"
	"amandasjostrom.se/day6"
	"amandasjostrom.se/day6"
)

func main(){
	//runDay1()
	//runDay2()
	//runDay3()
	runDay6_3()
	//runDay6()
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
func runDay6() {
	day6.Run(common.GetInput(6, "\n"))
}
func runDay6_3() {
	day6.Run(common.GetInput(6, "\n"))
}