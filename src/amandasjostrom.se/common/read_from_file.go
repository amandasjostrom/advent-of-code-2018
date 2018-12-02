package common

import (
	"io/ioutil"
	"fmt"
	"log"
	"strings"
	"os"
	"strconv"
)

func FromFile(filename string, separator string) (values []string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(":(")
		log.Fatal(err)
	}
	values = strings.Split(string(file), separator)
	return
}

func GetInput(day int, separator string) []string {
	absolutePath, _ := os.Getwd()
	return FromFile(absolutePath+"/src/amandasjostrom.se/day"+strconv.Itoa(day)+"/input.txt", separator)
}
