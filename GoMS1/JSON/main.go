package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	// "https://github.com/woeiyih/woeiyih_ETI_assg1.git"
)

type Passenger struct {
	Id        string
	Firstname string
	Lastname  string

	Contact struct {
		Mobileno string
		Email    string
	}
}

type Driver struct {
	Firstname        string
	Lastname         string
	Mobileno         string
	Email            string
	Identificationno int
	Licenceno        string
}

func EditPassenger(db *sql.DB, ID string, FN string, LN string, MN string, EA string) {
	query := fmt.Sprintf(
		"Update Drivers SET FirstName = '%s', Lastname = '%s', Mobileno='%s', Email='%s' WHERE ID = '%s'", FN, LN, MN, EA, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertPassenger(db *sql.DB, ID string, FN string, LN string, MN string, EA string) {
	query := fmt.Sprintf(
		"INSERT INTO Drivers VALUES ('%s', '%s', '%s', '%s', '%s')", ID, FN, LN, MN, EA)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func GetPassengers(db *sql.DB) {
	results, err := db.Query("Select * FROM my_db.Passengers")

	if err != nil {
		panic(err.Error())

	}
	for results.Next() {
		var passenger Passenger
		err = results.Scan(&passenger.Id, &passenger.Firstname, &passenger.Lastname, &passenger.Contact.Mobileno, &passenger.Contact.Email)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(passenger.Id, passenger.Firstname, passenger.Lastname, passenger.Contact.Mobileno, passenger.Contact.Email)
	}
}

type passengerInfo struct {
	Id        string `json:"Id"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Mobileno  string `json:"Mobileno"`
	Email     string `json:"Email"`
}

// for storing info of passengers
var passengers map[string]passengerInfo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Oobor!")
}

func displayAllPassengers(w http.ResponseWriter, r *http.Request) {
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v)
	}

	json.NewEncoder(w).Encode(passengers)
}

func passenger(w http.ResponseWriter, r *http.Request) {

	parameter := mux.Vars(r)

	if r.Header.Get("Content-Type") == "application/json" {
	}

	// creating new Passengers using POST
	if r.Method == "POST" {
		var nPassenger passengerInfo
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			json.Unmarshal(reqBody, &nPassenger)

			if nPassenger.Id == "" || nPassenger.Firstname == "" || nPassenger.Lastname == "" || nPassenger.Mobileno == "" || nPassenger.Email == "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte(
					"422 - Please supply passenger's information in JSON format"))
				return
			}

			//validate existing passengers
			if _, ok := passengers[parameter["passengerid"]]; !ok {
				passengers[parameter["passengerid"]] = nPassenger
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("201 - Passenger added successfully: " + parameter["passengerid"]))
			} else {
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte("409 - Duplicate of passenger ID found"))
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("422 - Please supply passenger's information in JSON format"))
		}
	}
}
func main() {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	fmt.Println("Database opened")

	GetRecords(db)

	passengers = make(map[string]passengerInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/passengers", displayAllPassengers())
	router.HandleFunc("/api/v1/passengers/{passengerid}", passenger).Methods(
		"GET", "PUT", "POST")
	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))

	defer db.Close()

}

func GetRecords(db *sql.DB) {
	panic("unimplemented")
}
