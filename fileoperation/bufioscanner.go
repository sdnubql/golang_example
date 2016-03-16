package main

import(
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {
	file,err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}


