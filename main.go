package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func getLine() string {
	if len(os.Args) < 2 {
		sc.Scan()
		return sc.Text()
	}

	line := ""
	var args []string
	for i, arg := range os.Args {
		if i == 1 {
			line = os.Args[1]
			continue
		}
		args = append(args, arg)
	}
	os.Args = args
	return line
}

func main() {
	n, _ := strconv.Atoi(getLine())

	// TODO : Something code
	sum := 0
	factorial := 1
	for i := 0; i < n; i++ {
		m, _ := strconv.Atoi(getLine())
		sum += m
		factorial = factorial * m
	}

	// Output
	fmt.Println(sum)
	fmt.Println(factorial)
}
