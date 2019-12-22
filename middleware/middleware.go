package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"cars/models"
	"github.com/gorilla/mux"
)

var carsDB []models.Car
var carIDTrack = 0

func init() {
	fmt.Println("initializing mock data")

	carsDB = append(carsDB, models.Car{ID: strconv.Itoa(carIDTrack), Model: "BMW X1M", Price: 50000})
	carIDTrack++
	carsDB = append(carsDB, models.Car{ID: strconv.Itoa(carIDTrack), Model: "BMW X2M", Price: 60000})
	carIDTrack++
	carsDB = append(carsDB, models.Car{ID: strconv.Itoa(carIDTrack), Model: "BMW X3M", Price: 70000})
	carIDTrack++
	carsDB = append(carsDB, models.Car{ID: strconv.Itoa(carIDTrack), Model: "BMW X4M", Price: 80000})
	carIDTrack++
	carsDB = append(carsDB, models.Car{ID: strconv.Itoa(carIDTrack), Model: "BMW X5M", Price: 90000})
	carIDTrack++
}

// ReadAllCars will return info of main page
func ReadAllCars(w http.ResponseWriter, r *http.Request) {
	// Set the way we will serve data between frontend and backend
	w.Header().Set("Content-Type", "application/json")
	// Allow cross origin connections making the routes accessible for everyone
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := readAllCars()
	json.NewEncoder(w).Encode(payload)
}
func readAllCars() []models.Car {
	// Create empty slice of cars
	var cars []models.Car
	// Loop through database cars and append each one to the empty slice of cars
	for _, element := range carsDB {
		cars = append(cars, element)
	}
	return cars
}

// ReadCar will return one car
func ReadCar(w http.ResponseWriter, r *http.Request) {
	// Set the way we will serve data between frontend and backend
	w.Header().Set("Content-Type", "application/json")
	// Allow cross origin connections making the routes accessible for everyone
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Get all params from url
	params := mux.Vars(r)
	payload := readCar(params["id"])
	json.NewEncoder(w).Encode(payload)
}
func readCar(id string) models.Car {
	var car models.Car
	for _, element := range carsDB {
		if element.ID == id {
			car = element
			break
		}
	}
	return car
}

// CreateCar will add a new car to the slice
func CreateCar(w http.ResponseWriter, r *http.Request) {
	// Set the way we will serve data between frontend and backend
	w.Header().Set("Content-Type", "application/json")
	// Allow cross origin connections making the routes accessible for everyone
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Allow the server to perform post operation
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// Allow the content type that is specified by client to be processed on server
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Declare an empty car
	var car models.Car
	// Take the car json from the client and decode it into car struct
	_ = json.NewDecoder(r.Body).Decode(&car)
	payload := createCar(car)
	json.NewEncoder(w).Encode(payload)
}
func createCar(car models.Car) models.Car {
	car.ID = strconv.Itoa(carIDTrack)
	carIDTrack++
	carsDB = append(carsDB, car)
	return car
}

// UpdateCar will modify a car chosen by id
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	// Set the way we will serve data between frontend and backend
	w.Header().Set("Content-Type", "application/json")
	// Allow cross origin connections making the routes accessible for everyone
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Allow the server to perform post operation
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// Allow the content type that is specified by client to be processed on server
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Get all params from url
	params := mux.Vars(r)
	// Declare an empty car
	var car models.Car
	// Take the car json from the client and decode it into car struct
	_ = json.NewDecoder(r.Body).Decode(&car)
	payload := updateCar(params["id"], car)
	json.NewEncoder(w).Encode(payload)
}
func updateCar(id string, car models.Car) models.Car {
	for index, element := range carsDB {
		if element.ID == id {
			element.Model = car.Model
			element.Price = car.Price
			carsDB[index] = element
			break
		}
	}
	return car
}

// DeleteCar will remove a car from db by its id
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	// Set the way we will serve data between frontend and backend
	w.Header().Set("Content-Type", "application/json")
	// Allow cross origin connections making the routes accessible for everyone
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Allow the server to perform post operation
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// Allow the content type that is specified by client to be processed on server
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Get all params from url
	params := mux.Vars(r)
	payload := deleteCar(params["id"])
	json.NewEncoder(w).Encode(payload)
}
func deleteCar(id string) models.Car {
	var car models.Car
	for index, element := range carsDB {
		if element.ID == id {
			carsDB = append(carsDB[:index], carsDB[index+1:]...)
			car = element
			break
		}
	}
	return car
}