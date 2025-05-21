# E-commerce GraphQL API

A GraphQL API for an e-commerce application built with Go, featuring user authentication, product management, and order processing.

## Features

- GraphQL API with gqlgen
- JWT Authentication
- Role-based access control (Admin/Customer)
- Product management
- Order processing
- Environment configuration
- Dockerization

## Prerequisites

- Go 1.24
- Git
- Docker (optional, for containerized deployment)

## Setup

### Option 1: Local Development

1. Clone the repository:
```bash
git clone https://github.com/yourusername/ecommerce-graphql-go.git
cd ecommerce-graphql-go
```

2. Install dependencies:
```bash
go mod download
```

3. Create environment file:
```bash
cp .env.example .env
```

4. Update the `.env` file with your configuration:
```env
PORT=8080
JWT_SECRET=your-secret-key-here
JWT_EXPIRATION=24h
```

5. Run the server:
```bash
go run server.go
```

The server will start at `http://localhost:8080/`

### Option 2: Docker Deployment

1. Clone the repository:
```bash
git clone https://github.com/yourusername/ecommerce-graphql-go.git
cd ecommerce-graphql-go
```

2. Create and configure environment file:
```bash
cp .env.example .env
```

3. Update the `.env` file with your configuration:
```env
PORT=8080
JWT_SECRET=your-secret-key-here
JWT_EXPIRATION=24h
```

4. Build and run with Docker:

Using docker-compose (recommended):
```bash
# Build and start the container in detached mode
docker-compose up --build
```

The server will start at `http://localhost:8080/`

## API Documentation

### Authentication

The API uses JWT tokens for authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-token>
```

#### Sample Tokens (Development)
- Admin Token: Generated on server start
- Customer Token: Generated on server start

### GraphQL Queries

#### Products

1. Get all products:
```graphql
query {
  products {
    id
    name
    price
    inStock
    description
    category
  }
}
```

2. Get product by ID:
```graphql
query {
  product(id: "1") {
    id
    name
    price
    inStock
    description
    category
  }
}
```

3. Get products with pagination and category filter:
```graphql
query {
  products(limit: 10, offset: 0, category: "electronics") {
    id
    name
    price
    inStock
    description
    category
  }
}
```

#### Orders

1. Get all orders (requires authentication):
```graphql
query {
  orders {
    id
    products {
      id
      name
      price
    }
    total
    createdAt
    status
    user {
      id
      name
    }
  }
}
```

2. Get order by ID (requires authentication):
```graphql
query {
  order(id: "1") {
    id
    products {
      id
      name
      price
    }
    total
    createdAt
    status
    user {
      id
      name
    }
  }
}
```

### GraphQL Mutations

#### Products (Admin only)

1. Create product:
```graphql
mutation {
  createProduct(input: {
    name: "New Product"
    price: 99.99
    inStock: 100
    description: "Product description"
    category: "electronics"
  }) {
    id
    name
    price
  }
}
```

2. Update product:
```graphql
mutation {
  updateProduct(
    id: "1"
    input: {
      name: "Updated Product"
      price: 89.99
      inStock: 50
      description: "Updated description"
      category: "electronics"
    }
  ) {
    id
    name
    price
  }
}
```

#### Orders (Customer only)

1. Place order:
```graphql
mutation {
  placeOrder(productIds: ["1", "2"]) {
    id
    total
    status
    products {
      id
      name
      price
    }
  }
}
```

## Role-based Access Control

- `ADMIN`: Can manage products (create, update)
- `CUSTOMER`: Can place orders and view their own orders


### Project Structure

```
.
├── graph/              # GraphQL schema and resolvers
├── data/              # Data layer
├── middleware/        # HTTP middleware
├── util/             # Utility functions
├── server.go         # Main server file
├── .env              # Environment variables
└── .env.example      # Example environment variables
```