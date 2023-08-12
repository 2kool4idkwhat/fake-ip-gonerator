package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func usage() {
	fmt.Println(
		`Usage: fake-ip-gonerator [number]
	
for example, to generate 5 IPs you would run:
	
./fake-ip-gonerator 5`)

	os.Exit(0)

}

func generateIp(a int) {
	for i := 0; a > i; a-- {
		fmt.Print(rand.Intn(255), ".", rand.Intn(255), ".", rand.Intn(255), ".", rand.Intn(255), "\n")
	}
	return
}

func main() {

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	arg := flag.Arg(0)

	if len(args) == 0 {
		generateIp(1)

	} else if len(args) >= 2 {
		fmt.Println("You can only input 1 number at once")
		os.Exit(1)

	} else {
		convertedArgs, err := strconv.Atoi(arg)

		if err != nil {
			fmt.Println("Uhhh... you know you can only input numbers, right?")
			os.Exit(1)
		} else {
			generateIp(convertedArgs)
		}

	}

}
