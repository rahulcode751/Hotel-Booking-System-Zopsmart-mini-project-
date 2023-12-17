# Zopsmart Mini Project (Hotel Booking)

The Hotel Booking project, developed using the [**GoLang**](https://go.dev/) language and the [**Gofr**](https://gofr.dev/) framework, operates as a comprehensive API for managing hotel room reservations. With a single model featuring fields such as ID, hotel name, customer name, date, price, room number, and customer contact number, the API facilitates essential CRUD operations necessary for efficient hotel booking management. This backend solution offers seamless integration for hotel systems, optimizing the process of adding, retrieving, updating, and deleting booking entries.

## Functions :
* **Add Booking Details**: This API is basically used for adding booking details, i.e., customer level details, date for booking, and price for rooms.
* **Get Booking Details By ID**: This API is basically used for extracting the booking details of hotel rooms based on a unique ID issued for each booking.
* **Update Booking Entry Detail By ID**: This API is basically used for updating the booking details, i.e., the customer-level details or the details related to room number, price, and date.
* **Delete booking entry details. By ID**: This API is basically used to delete the booking details, or, in simple words, to say that the customer has checked out of the hotel.

## **Setup**

### Installation
#### Clone Repository
```bash
  git clone https://github.com/rahulcode751/Gofr.git
  cd Gofr
```
    
#### Initializing the go module
```bash
  go mod init zopsmart-mini-project
  go get gofr.dev
```

#### Go to configs folder and create .env file
```bash
  cd configs
  echo. > "ConfigsFolderPath\.env"
```

## Environment Variables
#### To run this project, you will need to add the following environment variables to your .env file
```bash
  APP_NAME = zopsmart-mini-project
  APP_VERSION = dev
  DB_HOST = localhost
  DB_PORT = 3306
  DB_USER = your_sql_username
  DB_PASSWORD = your_sql_password
  DB_NAME = Bookings
  DB_DIALECT = mysql
```
#### Create Database in mysql
```bash
  > CREATE DATABASE Bookings;
  > USE Bookings;
  > CREATE TABLE bookings(
      id int AUTO_INCREMENT PRIMAY KEY,
      hotelname varchar(255),
      customername varchar(255)
      date varchar(255)
      price int,
      customercontact varchar(10),
      roomno int
    );
```
#### Download and sync the required modules and start the project.
```bash
  go mod tidy
  go run main.go
```

## **Postman API Collection Documentation**

### Documentation link = [Postman API Collection Documentation click here](https://documenter.getpostman.com/view/21947736/2s9YkkgNpq )


### Endpoints Description
#### Get Booking Details By ID
```bash
  URL - http://localhost:8000/hotelbooking/:id
  METHOD - GET
  Path Variables [id:1]
```
#### Add Booking Details To System
```bash
  URL - http://localhost:8000/hotelbooking
  METHOD - POST
  Body - (content-type = application/json)
  {
    "HotelName":"Vedantam",
    "CustomerName":"Abhay Butola",
    "Date":"15-12-2024",
    "Price":1000,
    "CustomerContact":"6264959991",
    "roomno":22
  }
```
#### Update Booking Entry Detail By ID
```bash
  URL - http://localhost:8000/hotelbooking/:id
  METHOD - PUT
  Path Variables [id:1]
  Body - (content-type = application/json)
  {
    "HotelName":"Vedantam",
    "CustomerName":"Abhay Butola",
    "Date":"15-12-2024",
    "Price":1000,
    "CustomerContact":"6264959991",
    "roomno":24
  }
```
#### Delete Booking Entry Detail By ID
```bash
  URL - http://localhost:8000/hotelbooking/:id
  METHOD - DELETE
  Path Variables [id:1]
```

## **UML Diagram** 

### Sequence Diagram
![SequenceDiagram](https://github.com/rahulcode751/Gofr/assets/73958355/2046553c-5425-4353-b81b-053e95c3db01)
#### The above diagram depicts the sequence diagram for a hotel booking system.
```bash
1. Firstly, the customer provides details to the receptionist.
2. The receptionist then checks the available room.
3. The receptionist then asks for the payment of room available.
4. The customer then makes payment for the room.
5. Receptionists add details to the system.
6. The system generates a unique ID for each booking detail.
7. The receptionist then provides the key to the room to the customer.
```
### UseCase Diagram
![UsecaseDiagram](https://github.com/rahulcode751/Gofr/assets/73958355/1bdfdf6b-06fb-4a7e-814b-c5c4c75711f2)
#### The above diagram depicts the UseCase Diagram for a hotel booking system.
```bash
1. Here, our project has two actors: the receptionist and the customer.
2. Receptionists have access to all the features, like adding, deleting, updating, and creating bookings.
3. The customer has provided the details of the room after payment is done.
```


### 🔗 Profile Links
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/rahul3008/)   

[![twitter](https://img.shields.io/badge/twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/Rahulbairagi77)




