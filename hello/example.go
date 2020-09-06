package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

func main() {
	test13()
}

func test10() {
	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
	fmt.Println(re.MatchString("slimshady99"))
	fmt.Println(re.MatchString("!asdf45dc"))
	fmt.Println(re.MatchString("roger"))
}

func test11() {
	fmt.Println("You have two seconds to calculate 19 *4")
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Time up!")
			return
		}
	}
}
func test12() {
	c := time.Tick(5 * time.Second)
	for t := range c {
		fmt.Printf("The time is now %v\n", t)
	}
}

func test13() {
	s1 := "2020-01-02T15:04:05+07:00"
	s2 := "2020-02-02T15:04:05+07:00"
	today, err := time.Parse(time.RFC3339, s1)
	if err != nil {
		log.Fatal(err)
	}

	tomorrow, err1 := time.Parse(time.RFC3339, s2)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(today.After(tomorrow))
	fmt.Println(today.Before(tomorrow))
	fmt.Println(today.Equal(tomorrow))
}
