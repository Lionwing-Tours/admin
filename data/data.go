package data

import (
	"booking/db"
	"log"
)

const (
	deletePassengerQuery = `DELETE * FROM passengers WHERE id=?`
	getPassengerQuery    = `SELECT salutation, name, lastname, email, phone, country FROM passengers WHERE id=? INNER JOIN country ON passengers.country_id=country.country_id`
	getDriverQuery       = `SELECT * FROM driver WHERE driver_id=?`
	getRecentBookingsPassengersQuery = ``
	getRecentRegistationsPassengersQuery = ``
	getRecentBookings = ``
	getRecentBookingsDriversQuery = ``
	getRecentRegistationsDriversQuery = ``
	
)

type Passenger struct {
	Salutation string
	FirstName  string
	LastName   string
	Email      string
	Phone      string
	Country    string
}

func AddVehicleDetails() {

}

func UpdateDriverDetails() {

}

func BanDriver(driverId int){
	
}

func DeletePassenger(passengerId int) error {
	_, err := db.MySqlDB.Exec(deletePassengerQuery, passengerId)
	if err != nil {
		log.Println("Unable to delete passenger: ", err)
		return err
	}
	return nil
}

func GetPassengerById(passengerId int) (p Passenger{}, err error) {
	row := db.MySqlDB.QueryRow(getPassengerQuery, passengerId)
	//p := Passenger{}
	err := row.Scan(&p.Salutation, &p.FirstName, &p.LastName, &p.Email, &p.Phone, &p.Country)
	if err != nil {
		log.Println("ERROR getting passeger data: ", err)
		return p, err
	}
	return p, nil

}

func GetPassengersForRecentBookings() {
	
}

func GetDriverById(driverId int) (err error) {
	//row := db.MySqlDB.QueryRow(getDriverQuery, driverId)

	return err
}

func InsertDriverAndVehicleInfo() {
	//insert driver and vehicle info
}
