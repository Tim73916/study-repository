package main

import (
	"fmt"
	"study/feature1"
	"study/feature2"
	simple_connection "study/feature_postgres"
)

func main() {
	fmt.Println("Hello, Git!")
	feature1.Feature1()
	feature2.Feature2()
	simple_connection.CheckConnection()
}
