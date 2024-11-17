package repository

import (
	"RestProject/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)

	if err != nil {
		log.Fatal("Error inserting employee: ", err)
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *EmployeeRepo) FindEmployeeById(empId string) (*model.Employee, error) {
	var emp model.Employee
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "employee_id", Value: empId}}).Decode(&emp)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployees() ([]model.Employee, error) {
	var emps []model.Employee
	cursor, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &emps)
	if err != nil {
		return nil, fmt.Errorf("result decoding failed: %v", err)
	}

	return emps, nil
}


func (r *EmployeeRepo) UpdateEmployeeById(empId string, updatedEmployee *model.Employee) (int64, error){
	result , err := r.MongoCollection.UpdateOne(
		context.Background(),
		bson.D{{Key:"employee_id", Value: empId}},
		bson.D{{Key:"$set", Value: updatedEmployee}},
	);

	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}


func(r *EmployeeRepo) DeleteEmployeeById(empId string) (int64, error){
	result, err := r.MongoCollection.DeleteOne(
		context.Background(),
		bson.D{{Key: "employee_id", Value: empId}},
	)

	if err != nil {
		return 0 , err
	}

	return result.DeletedCount, nil
}


func(r *EmployeeRepo) DeleteAllEmployees() (int64, error){
	result, err := r.MongoCollection.DeleteMany(
		context.Background(),
		bson.D{},
	)

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}