package main

import (
	"utils"
)

func main() {
	chan1 := make(chan interface{})
	go utils.GetMatrix(chan1)
	go utils.CutImg(<-chan1)

}
