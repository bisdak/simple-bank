package main

import (
	"fmt"

	"github.com/bisdak/simple-bank/util"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(util.RandomString(10))
	}

}
