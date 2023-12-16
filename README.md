# Zopsmart Mini Project (Hotel Booking)

It's a simple API mini-project for Zopsmart that contains the CRUD (Get,Post,Update,Delete) operation for hotel booking.

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
  echo. > "filepath\.env"
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




