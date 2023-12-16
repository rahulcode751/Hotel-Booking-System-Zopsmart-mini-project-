# Zopsmart Mini Project (Hotel Booking)

The Hotel Booking project, developed using the Go language and the Gofr framework, operates as a comprehensive API for managing hotel room reservations. With a single model featuring fields such as ID, hotel name, customer name, date, price, room number, and customer contact number, the API facilitates essential CRUD operations necessary for efficient hotel booking management. This backend solution offers seamless integration for hotel systems, optimising the process of adding, retrieving, updating, and deleting booking entries.

## Setup

### Installation
#### Clone Repository
```bash
  git clone https://github.com/rahulcode751/Gofr.git
  cd Gofr
```
    
#### initializing the go module
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
```bash
  APP_NAME = zopsmart-mini-project
  DB_HOST = localhost
  DB_PORT = 3306
  DB_USER = your_sql_username
  DB_PASSWORD = your_sql_password
  DB_NAME = Database_name
  DB_DIALECT = mysql
```


#### Download and sync the required modules and start project.
```bash
  go mod tidy
  go run main.go
```

# Postman Documentation

#### Documentation link = [Postman Documentation click here](https://documenter.getpostman.com/view/21947736/2s9YkkgNpq )

# Sequence Diagram

![SequenceDiagram](https://github.com/rahulcode751/Gofr/assets/73958355/2046553c-5425-4353-b81b-053e95c3db01)
#### The above diagram depicts the sequence diagram for an hotel booking system.
```bash
1. Firstly, the customer provides details to the receptionist.
2. The receptionist then checks the available room.
3. The receptionist then asks for the amount of room available.
4. The customer then makes payment for the room.
5. Receptionists add details to the system.
6. The system generates a unique ID for each booking detail.
7. The receptionist then provides the key to the room to the customer.
```

### 🔗 Profile Links
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/)

[![twitter](https://img.shields.io/badge/twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/Rahulbairagi77)




