# Golang Bank API

This is a simple Golang API that simulates login, logout and payment. This API uses json file to save its data, the JSON will automatically be made after the user uses the login endpoint.

### Endpoints
***
http://localhost:8080/login<br>
http://localhost:8080/logout<br>
http://localhost:8080/payment

### How to use endpoints
***
http://localhost:8080/login

Request
```json
{
    "username" : "john_doe",
    "password" : "12345"
}
```

if user not found then it will automatically register new it as a new user

Response
```json
User john_doe registered and logged in successfully
```
***
http://localhost:8080/payment

Request
```json
{
    "username" : "john_doe",
    "amount" : 100000
}
```

Response
```json
Payment of 100000.00 received from user john_doe
```

***
http://localhost:8080/logout

Request
```json
{
    "username" : "john_doe"
}
```

Response
```json
User john_doe logged out successfully
```

The data.json after all of the hits will look like this

data.json
```json
{
  "users": {
    "john_doe": {
      "username": "john_doe",
      "password": "12345",
      "logged_in": false
    }
  },
  "payments": [
    {
      "username": "john_doe",
      "amount": 100000
    }
  ]
}
```

### Different responses

/login endpoint

Wrong credentials
```json
Invalid credentials
```

Existing User
```json
User john_doe logged in successfully
```

/payment endpoint

User doesn't exist
```json
User not found
```

User not logged in
```json
User is not logged in
```

/logout enpoint

User doesn't exist
```json
User not found
```

User not logged in
```json
User is not logged in
```

