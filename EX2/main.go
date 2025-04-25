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
	// Khi Goroutine hoan thanh, phai goi Done() de thong bao WaitGroup
	defer wg.Done()

	// Khoi tao bien sum de luu tru tong
	sum := 0

	// Duyet qua dai so va tinh tong
	for _, num := range numbers {

		sum += num
	}

	// Gui ket qua qua kenh result
	result <- sum
}

func main() {

	// Khai bao slice chua ten hoc sinh
	var studentName []string

	// Khai bao map chua diem hoc sinh, key la ten
	studentScore := make(map[string]float64)

	// Them hoc sinh
	studentName = append(studentName, "Nguyen Van A")
	studentScore["Nguyen Van A"] = 8.5

	studentName = append(studentName, "Nguyen Van B")
	studentScore["Nguyen Van B"] = 9.0

	studentName = append(studentName, "Nguyen Van C")
	studentScore["Nguyen Van C"] = 7.5

	// Hien thi danh sach hoc sinh va diem
	for _, name := range studentName {
		fmt.Printf("- %s: %.2f điểm\n", name, studentScore[name])
	}

	greet := Person{name: "John", age: 30}
	fmt.Println(greet.greet())

	speak1 := student{name: "Alice", age: 20}
	fmt.Println(speak1.speak())

	speak2 := teacher{name: "Bob", age: 40}
	fmt.Println(speak2.speak())

	// Tao channel de gui ket qua
	sumChannel := make(chan int)

	// Goroutine tinh tong tu 1 den 10
	go func() {
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}

		// Gui ket qua ve channel
		sumChannel <- sum
	}()

	// Nhan ket qua tu channel
	result := <-sumChannel
	fmt.Printf("Tong tu 1 den 10 la: %d\n", result)

	// Khoi tao 1 day so can tinh tong
	numbers := []int{1, 2, 3, 4, 5}

	// Khoi tao doi tuong WaitGroup de dong bo hoa cac goroutine
	var wg sync.WaitGroup

	// Chia day so thanh 2 phan
	mid := len(numbers) / 2
	firstHalf := numbers[:mid]
	secondHalf := numbers[mid:]

	// Khoi tao 2 kenh (channel) de nhan ket qua tu 2 goroutine
	result1 := make(chan int)
	result2 := make(chan int)

	// Khoi dong 2 Goroutine
	wg.Add(2) // Thong bao chp WaitGroup rang co 2 goroutine dang chay
	go sumNumbers(firstHalf, result1, &wg)
	go sumNumbers(secondHalf, result2, &wg)

	// Doi 2 Goroutine hoan thanh va nhan ket qua
	go func() {
		wg.Wait()      // Doi cho den khi 2 goroutine hoan thanh
		close(result1) // Dong kenh result1
		close(result2) // Dong kenh result2
	}()

	// Ket hop ket qua tu 2 Goroutine va in ra tong cuoi cung

	sum1 := <-result1 // Nhan ket qua tu goroutine 1
	sum2 := <-result2 // Nhan ket qua tu goroutine 2

	fmt.Println("Tong cac so trong day so: ", sum1+sum2) // In ra tong cuoi cung
}
