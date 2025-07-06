package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/ritharam/logistics-api/database"
	"github.com/ritharam/logistics-api/functions"
	"github.com/ritharam/logistics-api/models"
)

func main() {
	godotenv.Load()
	database.ConnectDB()

	if len(os.Args) > 1 && os.Args[1] == "add" {
		s := bufio.NewScanner(os.Stdin)
		var ship models.Shipment
		fmt.Print("Enter Origin: ")
		s.Scan()
		ship.Origin = s.Text()
		fmt.Print("Enter Destination: ")
		s.Scan()
		ship.Destination = s.Text()
		fmt.Print("Enter Weight (kg): ")
		s.Scan()
		fmt.Sscan(s.Text(), &ship.Weight)
		fmt.Print("Enter Urgency (low/medium/high): ")
		s.Scan()
		ship.Urgency = strings.ToLower(s.Text())
		fmt.Print("Enter Your Email: ")
		s.Scan()
		email := s.Text()
		functions.HandleShipment(ship, email)
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/recommend", functions.RecommendHandler).Methods("POST")

	log.Println("🚀 Server running at :8080")
	http.ListenAndServe(":8080", r)
}
