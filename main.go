package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Employee struct {
	Id          int32 `json:"id"`
	Name        string `json:"name"`
	Salary int `json:"salary"`
}

func main() {
	employees := []Employee{}

	employees = append(employees, Employee{
		Name: "Hussein",
		Salary: 1001,
		Id: 1,
	})

	ahmed := Employee{
		Name: "Ahmed",
		Salary: 9000,
		Id: 1002,
	}

	employees = append(employees, ahmed)

	employees = append(employees, Employee{
		Name: "Rick",
		Salary: 5000,
		Id: 1003,
	})

	body, _ := json.Marshal(employees)

	fmt.Println(string(body)) //prints: [{"id":1,"name":"Hussein","salary":1001},{"id":1002,"name":"Ahmed","salary":9000},{"id":1003,"name":"Rick","salary":5000}]
	ioutil.WriteFile("jsondata.json", body, 0666) //we write the file to hard drive and also set permissions on the file as well
}