package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const LOG_FILE = "/Users/eadna/week-3-homework-2-denizedizcan/GoLog.txt"

func init() {
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type Rectangle struct {
	r float64
	l float64
}

func main() {
	err1 := strategy1a()
	if err1 != nil {
		fmt.Printf("Error: %s\n", err1)
	}
	
	err2 := strategy1b()
	if err2 != nil {
		fmt.Printf("Error: %s\n", err2)
	}

	strategy4()
}

func strategy1a() error {
	fmt.Println("\nstrategy1a")
	c1 := Rectangle{r: 12, l:5}
	circ1, err := area(c1)
	if err != nil {
		return err
	} else {
		fmt.Printf("area is %f\n", circ1)
		return nil
	}
}

func strategy1b() error {
	fmt.Println("\nstrategy1b")
	c1 := Rectangle{r: 12, l:-5}
	circ1, err := circumference(c1)
	if err != nil {
		return fmt.Errorf("hey, we've a problem, this is what I've got: %v", err)
	} else {
		fmt.Printf("Circumference is %f\n", circ1)
		return nil
	}
}


func strategy4() {
	fmt.Println("\nstrategy4")
	c1 := Rectangle{r: 12, l:-5}
	circ1, err := circumference(c1)
	if err != nil {
		log.Println("Calculation failed: ", err)
	} else {
		fmt.Printf("Circumference is %f\n", circ1)
	}
}


func circumference(c Rectangle) (float64, error) {
	if c.r < 0 || c.l < 0 {
		return -1, errors.New("Negative longitude: " + fmt.Sprintf("%f, %f", c.r, c.l))
	} else {
		circ := 2*c.l + 2*c.r
		return circ, nil
	}
}

func area(c Rectangle) (float64, error){
	if c.r < 0 || c.l < 0 {
		return -1, errors.New("Negative longitude: " + fmt.Sprintf("%f, %f", c.r, c.l))
	} else {
		circ := c.l*c.r
		return circ, nil
	}
}
