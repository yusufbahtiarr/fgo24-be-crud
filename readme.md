# Project Go - User Management (CRUD) API

A simple User Management RESTful API built with Go, Gin framework, and PostgreSQL.
This project demonstrates how to build a modular and scalable backend with key features such as:

- CRUD operations (Create, Read, Update, Delete)
- File upload functionality
- PostgreSQL integration using pgx
- API documentation using Swagger (with annotations)
- Redis caching for improved performance on selected endpoints

Designed for clarity, maintainability, and as a solid foundation for more advanced applications.

---

## API Endpoints Documentation

The following table lists the available API endpoints for managing users:

| Method | Endpoint     | Description    | Query Parameters             |
| ------ | ------------ | -------------- | ---------------------------- |
| GET    | `/users`     | Get List Users | ?search=[keyword] (optional) |
| POST   | `/users`     | Create User    | -                            |
| GET    | `/users/:id` | Detail User    | -                            |
| PATCH  | `/users/:id` | Update User    | -                            |
| DELETE | `/users/:id` | Delete User    | -                            |

## How to Clone and Use

Make sure you have Golang installed on your device.

#### 1. Clone the repository

```
git clone https://github.com/yusufbahtiarr/fgo24-be-crud.git
```

#### 2. Navigate into the project directory

```
cd fgo24-be-crud
```

#### 3. Setup .env

Create a .env file in the root folder with the following variables:

```
APP_PORT= (co:8800)
PGHOST=
PGPORT=
PGDATABASE= (co:postgres)
PGUSER=
PGPASSWORD=
REDIS_ADDR= (co:localhost:6379)
REDIS_PASSWORD=
REDIS_DB=

```

#### 4. Run the program

```
go run main.go
```

## 📄 License

This project is licensed under the **MIT License**.

## ©️ Copyright

&copy; 2025 Kodacademy
