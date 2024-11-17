package usecase

import (
	"RestProject/model"
	"RestProject/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}


type Response struct{
	Data  interface{} 
	Error string
}


func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type" , "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	_ , err = repo.InsertEmployee(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed while creating employee", err)
		res.Error = err.Error()
		return
	}

	log.Println("Employee create successfully")
	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-Type" , "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	empId := mux.Vars(r)["id"]

	if empId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid request")
		res.Error = "Invalid request"
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp , err := repo.FindEmployeeById(empId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error finding employee: ", err)
		res.Error = err.Error()
		return
	}

	log.Println("Employee found successfully")

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type" , "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)


	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emps , err := repo.FindAllEmployees()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error finding employee: ", err)
		res.Error = err.Error()
		return
	}

	log.Println("Employees found successfully")

	res.Data = emps
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type" , "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	empId := mux.Vars(r)["id"]

	if empId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid request")
		res.Error = "Invalid request"
		return
	}

	employee := model.Employee{}

	err := json.NewDecoder(r.Body).Decode(&employee)	

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}


	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	updatedCount , err := repo.UpdateEmployeeById(empId,&employee)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error finding employee: ", err)
		res.Error = err.Error()
		return
	}

	log.Println(updatedCount, " employees updated")

	res.Data = updatedCount
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type" , "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	empId := mux.Vars(r)["id"]

	if empId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid request")
		res.Error = "Invalid request"
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	deletedCount , err := repo.DeleteEmployeeById(empId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error finding employee: ", err)
		res.Error = err.Error()
		return
	}

	log.Println(deletedCount , " employees deleted")

	res.Data = deletedCount
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type" , "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)


	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	deletedCount , err := repo.DeleteAllEmployees()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error finding employee: ", err)
		res.Error = err.Error()
		return
	}

	log.Println(deletedCount," employees deleted")

	res.Data = deletedCount
	w.WriteHeader(http.StatusOK)

}
