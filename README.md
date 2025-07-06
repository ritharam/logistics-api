# Logistics Intelligence API

**Logistics Intelligence API** is a smart backend service built in Go that recommends the best courier service for shipments based on live traffic, urgency, and cost. It integrates with Google Maps API and optionally sends shipping confirmation emails. The system is designed to plug into e-commerce platforms at checkout for dynamic delivery recommendations.

---

## Features

- Intelligent courier recommendation (FedEx, DHL, UPS - mock)
- Real-time traffic-based ETA using Google Maps Directions API
- Shipment scoring based on urgency, time, and cost
- REST API for integration with external platforms
- CLI input support (`go run main.go add`) with map link + email output
- Auto-saving shipment history into MySQL database
- Email notification system using Gmail SMTP and HTML template
- .env-based configuration for secure key management

---

## Requirements

- Go 1.18+
- MySQL 5.7+ or 8.x
- Gmail account (with App Password)
- Google Maps API Key (Directions API enabled)

---

## Directory Structure

.
├── database/
│ ├── connection.go
│ ├── queries.go
│ └── repository.go
├── functions/
│ ├── mailer.go
│ └── shipment.go
├── models/
│ └── models.go
├── scraper/
│ └── maps.go
├── index.html
├── main.go
├── .env
└── README.md

---

## Configuration (.env)

Create a `.env` file in the root directory with the following keys:

DB_USER=your_mysql_user
DB_PASS=your_mysql_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=logistics

SENDER=your_gmail_address
APP_PASSWORD=your_gmail_app_password

MAPS_API_KEY=your_google_maps_api_key

> Ensure you have enabled the **Google Maps Directions API** from [Google Cloud Console](https://console.cloud.google.com/).

---

## How to Run the Project

1. **Clone the repository**

git clone https://github.com/yourusername/logistics-intelligence-api.git
cd logistics-intelligence-api
Create .env file

Refer to the above format and insert your actual credentials and API keys.

Start MySQL server

Ensure MySQL is running, and the database specified in DB_NAME exists. The tables will be created automatically.

Install dependencies (if any)

go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get github.com/go-sql-driver/mysql
Run the server

go run main.go
The server will start at http://localhost:8080.

Tables will be auto-created if not present.

---
## REST API Usage
POST /api/recommend
Request Body (JSON):

{
  "origin": "Chennai",
  "destination": "Delhi",
  "weight": 5.0,
  "urgency": "high"
}
Response:

[
  {
    "provider": "FedEx",
    "estimated_time": 120,
    "cost": 180,
    "score": 75.3
  },
  ...
]

---
## CLI Usage
You can also use the terminal to input a shipment and trigger scoring and email sending:

go run main.go add
You will be prompted to enter:

Origin
Destination
Weight
Urgency (low/medium/high)
Receiver's Email

A map link will be shown in the terminal and an email will be sent to the recipient.

---
## Email System

HTML template: index.html used to send formatted emails
Template is populated with shipment and recommendation info
Uses Gmail SMTP for delivery
No extra database column for email—sent based on input only

---
## License

This project is licensed under the Apache License 2.0.
See the LICENSE file for more information.

---
## Contributing

Feel free to fork this project and submit a pull request. For major changes, please open an issue first to discuss what you’d like to change.