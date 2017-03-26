package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"admin/data"

	//"github.com/go-sql-driver/mysql"
)

var (
	MySqlDB *sql.DB
	uid     string
	jres    string
)

func GetAdminRole(res http.ResponseWriter, req *http.Request) {

}

func AddUser(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	desc := req.FormValue("description")
	result, err := MySqlDB.Exec("INSERT INTO passengers (name, description) VALUES (?, ?)", name, desc)
	if err != nil {
		log.Println(result, err)
		return
	}

	log.Println("Table updated", result)

	jres = "Msg: Table updated"
	j_res, j_err := json.Marshal(jres)
	if j_err != nil {
		log.Println(j_err)
	}

	res.Write(j_res)

}

func AddDriver(w http.ResponseWriter, r *http.Request) {
	//vehicle driver details
	//driverId := strconv(r.FormValue("driver-id"))
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	dob := r.FormValue("dob")
	address := r.FormValue("address")
	mobile := r.FormValue("mobile")
	email := r.FormValue("email")
	nic := r.FormValue("nic")
	lang := r.FormValue("lang")
	licencePic := r.FormValue("licence-pic")                               //driver
	chauffeurLicencePic := r.FormValue("chauffeur-licence-pic")            //driver
	chauffeurGuideLicencePic := r.FormValue("chauffeur-guide-licence-pic") // driver

	//vehicle details
	owner := r.FormValue("owner") // bool
	ownerEmail := r.FormValue("owner-email")
	ownerContact := r.FormValue("owner-contact")

	vehicleRegDocumentPic := r.FormValue("vehicle-reg-doc-pic") //vehicle

	vehicleLocation := r.FormValue("vehicle-location")
	ac := r.FormValue("ac") // bool

	passengerInsurancePolicyPic := r.FormValue("passenger-insurance-policy-pic") //vehicle

	vehiclePic1 := r.FormValue("vehicle-pic1")
	vehiclePic2 := r.FormValue("vehicle-pic2")
	vehiclePic3 := r.FormValue("vehicle-pic3")

	driverPortrait := r.FormValue("driver-portrait")
}

func AddVehicle(res http.ResponseWriter, req *http.Request) {
	//add vehicle info
	vRegId := req.FormValue("reg-id")
	vType := req.FormValue("v-type")
	vMake := req.FormValue("v-make")
	vModel := req.FormValue("v-model")
	vYear := req.FormValue("v-year")
	vTransmission := req.FormValue("v-transmission")
	vEngineCapacity := req.FormValue("v-engine-cap")
	vAC := req.FormValue("v-ac")
	vFuel := req.FormValue("v-fuel")
	vNumPassengers := req.FormValue("v-num-passengers")
	vNumLuggages := req.FormValue("v-num-luggages")
	PassengersInsured := req.FormValue("passengers-insured")

}

func ViewPassengers(w http.ResponseWriter, r *http.Request) {
	//how many?

}

func EditDriver(w http.ResponseWriter, r *http.Request) {
	//update driver
}

func DeletePassenger(w http.ResponseWriter, r *http.Request) {
	passengerId, _ := strconv.Atoi(r.FormValue("passenger-id"))

	err := data.DeletePassenger(passengerId)
	if err != nil {
		log.Println(err)
		return
	}

}

func GetPassengersForRecentBookings(res http.ResponseWriter, req *http.Request) {
	//last 10 bookings by date in descending order
}

func GetPassengersForRecentRegistrations(res http.ResponseWriter, req *http.Request) {
	//last 10 registrations by date in descending order
}

func GetPassengerById(w http.ResponseWriter, r *http.Request) {
	passengerId, _ := req.FormValue("id")
	passenger, err := data.GetPassengerById(passengerId)
	if err != nil {
		log.Println("ERROR getting passenger data: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR getting passenger data"))
		return
	}
	w.WriteHeader(http.StatusFound)
	//get passenger By Id
}

func GetRecentBookings(res http.ResponseWriter, req *http.Request) {
	//get last 10 bookings by date in descending order
}

func GetAllBookings(res http.ResponseWriter, req *http.Request) {
	//Get all bookings
}

func GetAllDrivers(res http.ResponseWriter, req *http.Request) {
	//Get all drivers
}

func GetDriversForRecentBookings(res http.ResponseWriter, req *http.Request) {
	//Get Drivers for last 10 bookings in descending order
}

func GetDriversForRecentRegistrations(res http.ResponseWriter, req *http.Request) {
	//Get last 10 registered drivers
}

func GetDriverById(res http.ResponseWriter, req *http.Request) {
	driverId, _ := strconv.Atoi(req.FormValue("id"))

	err := data.GetDriverById(driverId)
	if err != nil {
		log.Println(err)
		return
	}

}

func SendInvoice(w http.ResponseWriter, r *http.Request) {
	//send invoice for driver id by end of month
	driverId, _ := strconv.Atoi(r.FormValue("driver-id"))
}

func EditBooking(res http.ResponseWriter, req *http.Request) {
	//same as add booking
	//check on how driver details get changed across multiple tables

	//unassign/reassign driver
	//change booking dates based on actual changes
}
