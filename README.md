## AWS CDK in Go

Introductory Infrastructure as Code project


Stack:
- **DynamoDB**. Stores user table.
- **Lambda**. Handles register, login and protected routes.
  - `/register`. Receives username and plain text password. Hashes the password and inserts user row into DynamoDB.
  - `/login`. Validates credentials, generates access token and sends it back.
  - `/protected`. Has a middleware that checks for bearer token and, if it's valid, sends back protected data.
- **API Gateway**. Acts as proxy that validates client requests and passes them to Lambda.

Visual representation, because why not:

```mermaid
sequenceDiagram
    participant Client
    participant API Gateway
    participant Lambda
    participant DynamoDB
    Client->>API Gateway: request to /register
    API Gateway->>Lambda: validates request
    Lambda->>DynamoDB: inserts user with hashed password
    Client->>API Gateway: request to /login
    API Gateway->>Lambda: validates request
    Lambda->>API Gateway: generates access token
    API Gateway->>Client: sends access token
    Client->>API Gateway: request to /protected
    API Gateway->>Lambda: validates Authorization header
    Lambda->>Lambda: middleware validates bearer token
    Lambda->>DynamoDB: gets protected data
    Lambda->>API Gateway: transforms protected data
    API Gateway->>Client: sends protected data
```

Big thanks to [Melkey](https://github.com/Melkeydev)
