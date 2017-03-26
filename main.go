package main

import (
	//"fmt"
	"log"
	//"os"
	"admin/db"
	"admin/handlers"
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/getAdminRole", handlers.GetAdminRole)
	http.HandleFunc("/addDriver", handlers.AddDriver)
	http.HandleFunc("/viewPassengers", handlers.ViewPassengers)
	http.HandleFunc("/editDriver", handlers.EditDriver)
	http.HandleFunc("/deleteDriver", handlers.DeleteDriver)
	http.HandleFunc("/deletePassenger", handlers.DeleteUser)
	http.HandleFunc("/banDriver", BanDriver)

	http.HandleFunc("/editBooking", handlers.EditBooking)

	//admin driver searches
	http.HandleFunc("/getPassengers/recentBooking", handlers.GetPassengersForRecentBookings)
	http.HandleFunc("/getPassengers/recentRegister", handlers.GetPassengersForRecentRegistrations)
	http.HandleFunc("/getPassenger/byId", handlers.GetPassengerById)
	http.HandleFunc("/getRecentBookings", handlers.GetRecentBookings)
	http.HandleFunc("/getAllBookings", handlers.GetAllBookings)
	http.HandleFunc("/getDrivers", handlers.GetAllDrivers)
	http.HandleFunc("/getDrivers/recentBooking", handlers.GetDriversForRecentBookings)
	http.HandleFunc("/getDrivers/recentRegister", handlers.GetDriversForRecentRegistrations)
	http.HandleFunc("/getDriver/byId", handlers.GetDriverById)
	http.HandleFunc("/sendInvoice", handlers.SendInvoice)
	log.Fatal(endless.ListenAndServe("localhost:8081", nil))
}
