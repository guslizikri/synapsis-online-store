<div align='center' style="text-align: center;">

<h1 style="border:0;margin:1rem">Online Store Rest API</h1>

Backend for Online Store

<hr>
<br>

</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Postman Collection](#postman-collection)
- [Build Program](#Build-Program-with-Docker)
- [Contributors](#contributors)
- [Suggestion](#suggestion)

## Overview

It is an online store application that provides features for user management, product management, shopping cart, and transactions.

## Features

1. User Management:

   - Login
   - Register

2. Product Management:

   - Add products
   - Fetch products by category

3. Cart Management:

   - Add to cart
   - Remove from cart
   - Fetch cart

4. Transaction Management:
   - Create transactions

## Technologies Used

- Programming Language: Go (Golang)
- Framework: Go Fiber
- Database: PostgreSQL
- ORM/SQL Library: SQLx (jmoiron)

## Getting Started

1. Clone this repo

   ```bash
   git clone https://github.com/ninja1cak/coffeshop-be
   ```

2. Enter the directory

   ```bash
   cd ./
   ```

3. Install all dependencies

   ```bash
   go mod tidy
   ```

4. Start the local server

   ```bash
   go run cmd/main.go
   ```

## Postman Collection

cd ./docs

## Build and Run Instructions

# Build Program with Docker

    1. Build Go binary
    go build -o "./build/onlinestore.exe" ./cmd/main.go

    2. Build Docker Image

Use the following command to build the Docker image:
docker build -t zikrigusli/coffeeshopbe:1 .
(Replace zikrigusli/coffeeshopbe:1 with your desired image name and tag.)

# Run and Create Container from Image

To run the application and create a container:
docker compose up -d

# Delete Container

To stop and delete the container:
docker compose down

## Contributors

Currently, there are no contributors to this project. If you would like to contribute, you can submit a pull request.

## Suggestion

If you find bugs / find better ways / suggestions you can pull request.
