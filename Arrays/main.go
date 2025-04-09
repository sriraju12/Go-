package main

import "fmt"

func main() {

	hobbies := [3]string{"Reading", "Traveling", "Cooking"}

	fmt.Println("Hobbies:", hobbies)

	fmt.Println("first element:", hobbies[0])

	fmt.Println("second and third elements:", hobbies[1:3])

	newHobbies := hobbies[:2]

	fmt.Println("newHobbies:", newHobbies)

	newHobbies = hobbies[1:3]

	fmt.Println("newHobbies:", newHobbies)

	courseGoals := []string{"Learn Go", "Build a web app"}

	fmt.Println("Course Goals:", courseGoals)

	courseGoals[1] = "Build a REST API"
	fmt.Println("Updated Course Goals:", courseGoals)

	courseGoals = append(courseGoals, "Deploy to AWS")
	fmt.Println("Updated Course Goals:", courseGoals)

	products := []Product{
		{"Laptop", 1, 999.99},
		{"Smartphone", 2, 499.99},
	}

	fmt.Println("Products:", products)

	products = append(products, Product{"Tablet", 3, 299.99})

	fmt.Println("Updated Products:", products)

}

type Product struct {
	title string
	id    int
	price float64
}
