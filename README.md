# Zopsmart Mini Project (Hotel Booking)

The Hotel Booking project, developed using the Go language and the Gofr framework, operates as a comprehensive API for managing hotel room reservations. With a single model featuring fields such as ID, hotel name, customer name, date, price, room number, and customer contact number, the API facilitates essential CRUD operations necessary for efficient hotel booking management. This backend solution offers seamless integration for hotel systems, optimising the process of adding, retrieving, updating, and deleting booking entries.

## Setup

## Installation
#### Clone Repository
```bash
  git clone https://github.com/rahulcode751/Gofr.git
  cd Gofr
```
    
#### Install zopsmart-mini-project with go

```bash
  go mod init zopsmart-mini-project
```
#### Go to configs folder and create .env file
```bash
  cd configs
  echo. > "ConfigsFolderPath\.env"
```

## Environment Variables

#### To run this project, you will need to add the following environment variables to your .env file

`APP_NAME`= zopsmart-mini-project

`DB_HOST`= localhost

`DB_PORT`= 3306

`DB_USER`= your_sql_username

`DB_PASSWORD`= your_sql_password

`DB_NAME`= Database_name

`DB_DIALECT`= mysql


#### Download and sync the required modules and start project.
```bash
  go mod tidy
  go run main.go
```

# Postman Documentation

#### Documentation link = [Postman Documentation click here](https://documenter.getpostman.com/view/21947736/2s9YkkgNpq )




