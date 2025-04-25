package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
}

type student struct {
	name string
	age  int
}

type teacher struct {
	name string
	age  int
}

func (p *Person) greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.name, p.age)
}

type Speaker interface {
	speak() string
}

func (s *student) speak() string {
	return fmt.Sprintf("%s is a student and is %d years old.", s.name, s.age)
}

func (s *teacher) speak() string {
	return fmt.Sprintf("%s is a student and is %d years old.", s.name, s.age)
}

func sumNumbers(numbers []int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	
	sum := 0

	for _, num := range numbers {

		sum += num
	}

	result <- sum
}

func main() {

	var studentName []string

	studentScore := make(map[string]float64)

	studentName = append(studentName, "Nguyen Van A")
	studentScore["Nguyen Van A"] = 8.5

	studentName = append(studentName, "Nguyen Van B")
	studentScore["Nguyen Van B"] = 9.0

	studentName = append(studentName, "Nguyen Van C")
	studentScore["Nguyen Van C"] = 7.5

	for _, name := range studentName {
		fmt.Printf("- %s: %.2f điểm\n", name, studentScore[name])
	}

	greet := Person{name: "John", age: 30}
	fmt.Println(greet.greet())

	speak1 := student{name: "Alice", age: 20}
	fmt.Println(speak1.speak())

	speak2 := teacher{name: "Bob", age: 40}
	fmt.Println(speak2.speak())

	sumChannel := make(chan int)

	go func() {
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}

		sumChannel <- sum
	}()

	result := <-sumChannel
	fmt.Printf("Tong tu 1 den 10 la: %d\n", result)

	numbers := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup

	mid := len(numbers) / 2
	firstHalf := numbers[:mid]
	secondHalf := numbers[mid:]

	result1 := make(chan int)
	result2 := make(chan int)

	wg.Add(2)
	go sumNumbers(firstHalf, result1, &wg)
	go sumNumbers(secondHalf, result2, &wg)

	go func() {
		wg.Wait()     
		close(result1) 
		close(result2) 
	}()

	sum1 := <-result1 
	sum2 := <-result2 

	fmt.Println("Tong cac so trong day so: ", sum1+sum2)
}
