package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/okpalaChidiebere/protobuff-go/pb"
	"google.golang.org/protobuf/proto"
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

	chidi := pb.Employee{}
	chidi.Id = 1001
	chidi.Name = "Chidiebere"
	chidi.Salary = 1001
	//fmt.Println("My name is", chidi.GetName())

	ifeanyi := pb.Employee{
		Id: 1004,
		Name: "Ifeanyi",
		Salary: 20000,
	}
	chiamaka := pb.Employee{
		Id: 1004,
		Name: "Chiamaka",
		Salary: 30000,
	}

	employees2 := pb.Employees{ Employees: []*pb.Employee{} }
	employees2.Employees = []*pb.Employee{
		&chiamaka,
	}
	employees2.Employees = append(employees2.Employees, &ifeanyi, &chidi)
	//fmt.Println(employees2.String())

	data, err := proto.Marshal(&employees2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	ioutil.WriteFile("employeesbinary", data, 0666)

	employeesFromFile := pb.Employees{}
	if err := proto.Unmarshal(data, &employeesFromFile); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("employees from protobuf file %+v\n", employeesFromFile.String())
	
	body, _ := json.Marshal(employees)

	fmt.Println(string(body)) //prints: [{"id":1,"name":"Hussein","salary":1001},{"id":1002,"name":"Ahmed","salary":9000},{"id":1003,"name":"Rick","salary":5000}]
	ioutil.WriteFile("jsondata.json", body, 0666) //we write the file to hard drive and also set permissions on the file as well
}