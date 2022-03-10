package app

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	fmt.Println("Application Start")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Scan stdin error", scanner.Err())
	}
}
