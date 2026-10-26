// pass_fail сообщает, сдал ли пользователь экзамен
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade == 100 {
		status = "perfect"
	} else if grade >= 60 {
		status = "good"
	} else if grade < 60 {
		status = "bad"
	}
	fmt.Println(status)
}
