package main

import "log"

func main() {
	nextOdd := makeOddGenerator()
	log.Println(nextOdd()) // 1
	log.Println(nextOdd()) // 3
	log.Println(nextOdd()) // 5
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}