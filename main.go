package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"database/sql"
	"io/ioutil"
	"github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	"https://github.com/woeiyih/woeiyih_ETI_assg1.git"
)

type Passenger struct {
	Id string
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

func main() {

}

func EditPassenger (db *sql.DB, ID string, FN string, LN string, MN string, EA string) {
	query := fmt.Sprintf(
		"Update Drivers SET FirstName = '%s', Lastname = '%s', Mobileno='%s', Email='%s' WHERE ID = '%s'", FN, LN, MN, EA, ID)
		_, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
	)
}

func InsertPassenger (db *sql.DB, ID string, FN string, LN string, MN string, EA string) {
	query := fmt.Sprintf(
		"INSERT INTO Drivers VALUES ('%s', '%s', '%s', '%s', '%s' '%s')", ID, FN, LN, MN, EA)
		_, err := db.Query(query)

		if err != nil {
			panic(err.Error())
		}
	)
}

func GetPassengers(db *sql.DB) {
	results, err := db.Query("Select * FROM my_db.Passengers")

	if err != nil {
		panic(err.Error())

	}
	for results.next() {
		var passenger Passenger
		err = results.Scan(&passenger.Id, &passenger.Firstname, &passenger.Lastname, &passenger.Contact.Mobileno, &passenger.Contact.Email)

	}
}