package main

import (
	"RestProject/usecase"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init(){
	// load .env file

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Loaded .env file")


	mongoClient , err = mongo.Connect(context.Background(),
		options.Client().ApplyURI(os.Getenv("MONGO_URL")));

	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}


	err = mongoClient.Ping(context.Background(), readpref.Primary());

	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}

	log.Println("Connected to MongoDB")

}


func main(){

	defer mongoClient.Disconnect(context.Background())

	r := mux.NewRouter()

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	empService := usecase.EmployeeService{MongoCollection: coll}

	r.HandleFunc("/health",healthHandler).Methods(http.MethodGet);
	r.HandleFunc("/employee",empService.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}",empService.GetEmployeeByID).Methods(http.MethodGet)
	r.HandleFunc("/employee",empService.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}",empService.UpdateEmployeeByID).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}",empService.DeleteEmployeeByID).Methods(http.MethodDelete)
	r.HandleFunc("/employee" , empService.DeleteAllEmployee).Methods(http.MethodDelete)

	log.Println("Server running on port 4444")

	http.ListenAndServe(":4444" , r);
	
}



func healthHandler(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running...."))
}