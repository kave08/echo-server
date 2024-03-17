# echo-server
Echo Server is a simple web server built with Go, utilizing the Echo framework for handling HTTP requests and MongoDB for data storage. This project demonstrates basic user management functionalities, including creating, retrieving, and listing users.

# Features
 . User Management: Create, retrieve, and list users.
 . MongoDB Integration: Utilizes MongoDB for data persistence.
 . Echo Framework: Leverages the Echo framework for handling HTTP requests.

 # Getting Started

Prerequisites
 - Go 1.16 or later
 - MongoDB

 ## Installation
1. Clone the repository:
   git clone https://github.com/kave08/echo-server.git
2. Navigate to the project directory:
    go mod download
3. Install dependencies:
   go mod download
4. Start the server:
    go run main.go

# Usage
The server runs on port 8080 by default. You can change the port in the config.yml file.

# Endpoints
- GET /users/getlist: Retrieves a list of all users.
- POST /users/create: Creates a new user.

# Contributing
Contributions are welcome. Please feel free to submit a pull request or open an issue.

