package main

import (
	"fmt"
	// "time"
)

func main() {
	// o Go tem um conceito que se chama Zero Value,
	// onde variaveis declaradas sem valor recebem
	// um valor padrao, dependendo do seu tipo.
	var a int
	const b int = 1
	a = 1
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("Type: %T - Value: %v\n", a, a)

	total := sum(a, b)
	fmt.Printf("Total: %v\n", total)

	value1, value2 := returnValues("value1", "value2")
	fmt.Printf("Value1: %v - Value2: %v\n", value1, value2)

	// array tem tamanho fixo, e todo array e slice deve ter o mesmo tipo de dados
	var c [2]int = [2]int{1, 2}
	c[0] = 3
	c[1] = 4
	fmt.Printf("C: %v\n", c)

	// slice tem tamanho dinamico, mas deve ter valores inicializados
	// do contrario, deve-se usar a funcao make()
	var d []int
	// d[0] = 1 -> erro
	fmt.Printf("D: %v\n", d)

	var e = make([]int, 2)
	fmt.Printf("E: %v\n", e)

	f := []int{1, 2}
	fmt.Printf("F: %v\n", f)
	f = append(f, 3, 4)
	fmt.Printf("F: %v\n", f)

	// maps
	var g map[string]int = make(map[string]int)
	fmt.Printf("G: %v\n", g)

	ages := map[string]int{
		"Rafael": 23,
		"Batman": 30,
	}
	fmt.Printf("Ages: %v\n", ages)
	ages["Novo campo"] = 50
	fmt.Printf("Ages: %v\n", ages)

	// structs
	type Person struct {
		Name string
		Age  int
	}

	// struct com heranca
	type Professional struct {
		Person
		Job string
	}

	var p1 Person
	p1.Name = "Rafael"
	p1.Age = 23
	fmt.Printf("P1: %v\n", p1)

	var p2 Person = Person{
		Name: "Batman",
		Age:  30,
	}
	fmt.Printf("P2: %v\n", p2)

	p3 := Person{Name: "Chico Science", Age: 27}
	fmt.Printf("P3: %v\n", p3)

	professional := Professional{
		Person: p3,
		Job:    "Artist",
	}
	fmt.Printf("Professional: %v\n", professional)

	// map tipado com uma struct
	people := map[string][]Person{
		"P1": {
			{Name: "Rafael", Age: 23},
		},
	}
	fmt.Printf("People: %v\n", people)

	// slice tipado com uma struct
	peopleSlice := []Person{
		{Name: "Rafael", Age: 23},
	}
	peopleSlice = append(peopleSlice, Person{Name: "Batman", Age: 30})
	fmt.Printf("PeopleArray: %v\n", peopleSlice)

	// if
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else if 7%2 != 0 {
		fmt.Println("7 is odd")
	} else {
		fmt.Println("7 is neither even nor odd")
	}

	// switch
	switch 7 % 2 {
	case 0:
		fmt.Println("7 is even")
	case 1:
		fmt.Println("7 is odd")
	default:
		fmt.Println("7 is neither even nor odd")
	}

	// for loop
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	for i := range 2 {
		fmt.Println(i)
	}

	// for loop infinito
	for {
		fmt.Println("loop")
		// time.Sleep(1 * time.Second)
		break
	}

	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", i, fruit)
	}

	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := returnEvenNumbers(numbers[:])
	fmt.Println("Even numbers", evenNumbers)
}

func sum(a int, b int) int {
	return a + b
}

func returnValues(value1 string, value2 string) (string, string) {
	return value1, value2
}

func returnEvenNumbers(numbers []int) []int {
	var evenNumbers []int

	for i, n := range numbers {
		fmt.Println("index", i, "number", n)

		if n%2 == 0 {
			evenNumbers = append(evenNumbers, n)
		}
	}

	return evenNumbers
}
