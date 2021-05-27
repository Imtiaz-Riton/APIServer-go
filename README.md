# Go Clean Architecture
Example that shows core principles of the Clean Architecture in Golang projects.

More on Clean Architecture can be found <a href="https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html">here</a>.

Here is an implimentaion of Clean Architecture <a href="https://github.com/bxcodec/go-clean-arch">here</a>.

### Project Description&Structure:
REST API with custom JWT-based authentication system. Core functionality is about creating and managing products

#### Structure:
4 Domain layers:

- Models layer
- Repository layer
- UseCase layer
- Delivery layer

## API:



### POST /auth/sign-in

Request to get JWT Token based on user credentials

##### Example Input:
```
{
	"username": "username",
	"password": "password"
} 
```

##### Example Response:
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

### POST /api/product

Creates new product

##### Example Input:
```
{
	"id" : "12345",
	"title" : "Walton",
	"amount" : 234,
	"price" : 23000
} 
```

### GET /api/products

Returns all products

##### Example Response:
```
{
	"products": [
            {
                "id": "5da2d8aae9b63715ddfae856",
                "title": "Walton",
                "amount": 234,
                "price": 23000
            }
    ]
} 
```

### DELETE /api/product

Deletes bookmark by ID:

##### Example Input:
```
{
	"id": "12345"
} 
```


## Requirements
- go 1.13

