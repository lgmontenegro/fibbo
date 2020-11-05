package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println(err)
		return
	}

	//fibbo := make(map[int]int, arg)
	fibbo := make([]int, arg)

	for key := range fibbo {
		switch key {
		case 0:
			continue
		case 1:
			fibbo[key] = 1
		default:
			fibbo[key] = fibbo[key-2] + fibbo[key-1]
		}
	}

	fmt.Println(fibbo)
}
