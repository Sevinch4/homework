package main

import "fmt"

func main() {
	fmt.Println("homework-2")
	//task-1
	fmt.Println("Task-1")
	var (
		n1  = 123 //123... 2->%100/10....3->%100%10....1->/100
		rev int
	)
	rev = (n1/100)*100 + ((n1%100)%10)*10 + (n1%100)/10
	fmt.Println(n1, rev)

	fmt.Println("Task-2")
	var (
		n2  = 1235
		rem int
	)
	rem = (n2 / 100) * 100
	rem = (rem / 100) % 10
	fmt.Println(n2, ", yuz = ", rem)

	fmt.Println("Task-3")
	var (
		n3   = 3458
		rem1 int
	)
	rem1 = (n3 / 100) * 100
	rem1 = rem1 / 1000
	fmt.Println(n3, ", ming = ", rem1)

	fmt.Println("Task-4")
	fmt.Print("Input N: ")
	var (
		N     int
		minut int
	)
	fmt.Scan(&N)
	minut = N / 60
	fmt.Println("Min: ", minut)

	fmt.Println("Task-5")
	fmt.Print("Input N: ")
	var (
		N1 int
		h  int
	)
	fmt.Scan(&N1)
	h = N1 / 3600
	fmt.Println("Soat: ", h)

	fmt.Println("Task-6")
	fmt.Print("Input N: ")
	var (
		N2     int
		minut1 int
	)
	fmt.Scan(&N2)
	minut1 = N2 / 60
	fmt.Printf("Min: %d, secund: %d\n", minut1, N2)

	fmt.Println("Task-7")
	var (
		N3 int
		h1 int
	)
	fmt.Scan(&N3)
	h1 = N3 / 3600
	fmt.Printf("Soat: %d, secund: %d\n", h1, N3)

	fmt.Println("Task-8")
	var (
		N4   int
		min  int
		hour int
	)
	fmt.Scan(&N4)
	hour = N4 / 3600
	min = N4 / 60
	fmt.Printf("Soat: %d, Minut: %d, Secund: %d\n", hour, min, N4)

	fmt.Println("Task-9")
	var (
		n9   = 9458
		rem9 int
	)
	rem9 = (n9 / 100) * 100
	rem9 = rem9 / 1000
	fmt.Println(n9, ", ming = ", rem9)
}
