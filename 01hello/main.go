package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("This is", input)
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("thbgurg is:", num+1)
	}

	///handling time

	time := time.Now()
	fmt.Println(time)
	fmt.Println(time.Format("01-02-2006 Saturday"))
}
