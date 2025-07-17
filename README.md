# Key-Value Timestamp Store (Golang + MySQL)

This project is a backend service that allows storing and retrieving values for keys at specific timestamps. It is built using Golang, Gin, GORM, and MySQL. You can insert multiple values for the same key with different timestamps and retrieve the latest value for a given key at or before a specific timestamp.


## Technology Used

- Golang 1.24.3
- Gin (HTTP web framework)
- GORM (ORM library)
- MySQL (Relational Database)
- dotenv (for environment configuration)


## Setup Instructions

### 1. Clone the Repository

git clone https://github.com/Hariharan3008/tryg.git
cd tryg


### 2. Set up the environment

  ## Create a .env file in the root directory and add the following

   DB_USER=root
   DB_PASSWORD=your_mysql_password (use your mysql password)
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=trygdb

### 3. Running the application

go run main.go

### 4. API end points and the curl commands 

  ## Note : The expected response attached are as per the date and time the tests were carried out.

  ## 4.1 Store value,key and timestamp

  ## Resquest :

  curl -X PUT http://localhost:8080/api/v1/value \
  -H "Content-Type: application/json" \
  -d '{"key": "item1", "value": "v1", "timestamp": 100}'

  ## Response :

  {
  "message": "Entry saved successfully",
  "data": {
    "ID": 9,
    "Key": "item1",
    "Value": "v1",
    "Timestamp": 100,
    "CreatedAt": "2025-07-15T21:48:31.109+05:30"
  }
}

## 4.2 Retrieves the latest value for a key at or before a given timestamp.

## Request :

curl "http://localhost:8080/api/v1/value/at?key=item1&timestamp=100"

## Response :

{
  "message": "Value retrieved successfully",
  "data": {
    "ID": 9,
    "Key": "item1",
    "Value": "v1",
    "Timestamp": 100,
    "CreatedAt": "2025-07-15T16:18:31.109Z"
  }
}

## 4.3 Insert newer timestamp

## Request :

curl -X PUT http://localhost:8080/api/v1/value \
  -H "Content-Type: application/json" \
  -d '{"key": "item1", "value": "v2", "timestamp": 200}'


## 4.4 Fetch between two values

## Request :

curl "http://localhost:8080/api/v1/value/at?key=item1&timestamp=150"

## Response :

{
  "message": "Value retrieved successfully",
  "data": {
    "ID": 9,
    "Key": "item1",
    "Value": "v1",
    "Timestamp": 100,
    "CreatedAt": "2025-07-15T16:18:31.109Z"
  }
}

## 4.5 Fetch after all timestamps

## Request :

curl "http://localhost:8080/api/v1/value/at?key=item1&timestamp=999"

## Response :

{
  "message": "Value retrieved successfully",
  "data": {
    "ID": 10,
    "Key": "item1",
    "Value": "v2",
    "Timestamp": 200,
    "CreatedAt": "2025-07-15T16:18:45.552Z"
  }
}

## 4.6 Fetch before any entry

## Request :

curl "http://localhost:8080/api/v1/value/at?key=item1&timestamp=50"

## Response :

{
  "error": "No value found for the given key and timestamp"
}

## 4.7 If timestamp is zero (Edge case scenario)

## Request (Add entry):

curl -X PUT http://localhost:8080/api/v1/value \
  -H "Content-Type: application/json" \
  -d '{"key": "edgeCase", "value": "zero", "timestamp": 0}'

## Request (Retrieve entry)

  curl "http://localhost:8080/api/v1/value/at?key=edgeCase&timestamp=0"


## Response :

{
  "message": "Value retrieved successfully",
  "data": {
    "ID": 11,
    "Key": "edgeCase",
    "Value": "zero",
    "Timestamp": 0,
    "CreatedAt": "2025-07-15T16:20:17.175Z"
  }
}



       