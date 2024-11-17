package repository

import (
	"RestProject/model"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient,err := mongo.Connect(context.Background(),
	 options.Client().ApplyURI("mongodb+srv://admin_oooppp:yourpass123@cluster0.hoqad94.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	 if err != nil{
		log.Fatal("Error connecting to MongoDB: ", err)
	 }

	 log.Println("Connected to MongoDB")

	 pingErr := mongoTestClient.Ping(context.Background(), readpref.Primary());

	 if pingErr != nil {
		log.Fatal("Error pinging MongoDB: ", pingErr)
	 }

	 log.Println("Ping to MongoDB successful")

	 return mongoTestClient	
}


func TestMongoOperations(t *testing.T){
	mongoTestClient := newMongoClient()

	defer mongoTestClient.Disconnect(context.Background())

	// dummy data
	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	// create employee
	empRepo := EmployeeRepo{MongoCollection: coll}


	t.Run("InsertEmployee 1", func(t *testing.T){
		emp := model.Employee{
			EmployeeID: emp1,
			Department: "IT",
			Name: "John Doe",
		}

		result , err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Errorf("Error inserting employee 1: %v", err)
		}

	t.Log("Inserted employee 1 with ID: ", result)
	})
	t.Run("InsertEmployee 2", func(t *testing.T){
		emp := model.Employee{
			EmployeeID: emp2,
			Department: "IT",
			Name: "John Doe",
		}

		result , err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Errorf("Error inserting employee 2: %v", err)
		}

	t.Log("Inserted employee 2 with ID: ", result)
	})

	t.Run("FindEmployeeById", func(t *testing.T){
		emp,err := empRepo.FindEmployeeById(emp2)

		if err != nil {
			t.Errorf("Error finding employee: %v", err)
		}


		t.Log("Found employee: ", emp)

	})


	t.Run("FindAllEmployees", func(t *testing.T){
		emp , err := empRepo.FindAllEmployees()

		if err != nil {
			t.Errorf("Error finding all employees: %v", err)
		}

		if len(emp) == 0 {
			t.Errorf("No employees found")
		}

		t.Log("Found all employees: ", emp)
	})


	t.Run("UpdateEmployeeById" , func(t *testing.T){

		updatedEmployee := model.Employee{
			EmployeeID: emp1,
			Department: "IT",
			Name: "John Doe",
		}


		result , err := empRepo.UpdateEmployeeById(emp1, &updatedEmployee)

		if err != nil{
			t.Errorf("Error updating employee: %v", err)
		}

		t.Log("Number of employees updated: ", result)
	})


	t.Run("DeleteEmployeeById" , func(t *testing.T){
		result , err  := empRepo.DeleteEmployeeById(emp1)


		if err != nil {
			t.Errorf("Error deleting employee: %v", err)
		}

		t.Log("Number of employees deleted: ", result)
	})


	t.Run("DeleteAllEmployees" , func(t *testing.T){
		result , err := empRepo.DeleteAllEmployees()  

		if err != nil {
			t.Errorf("Error deleting all employees: %v", err)
		}

		t.Log("Number of employees deleted: ", result)
	})


}