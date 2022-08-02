# go-graphql

Simple example of using GraphQL with Go

```
go run server.go
```

open browser and go to localhost:8080

paste following queries inside graphiql playground and use them to create/retrieve records:

```
mutation createCustomer {
  createCustomer(input: { name: "João", email: "joao@gmail.com" }) {
    id
    name
    email
  }
}

mutation createItem {
  createItem(input: { name: "GeForce RTX 3060 TI", price: 599.99 }) {
    id
    name
    price
  }
}

mutation createOrder {
  createOrder(input: { 
    customerId: "CUSTOMER_UUID_FROM_CREATE_CUSTOMER_REQUEST", 
    items: [ {itemId: "ITEM_UUID_FROM_CREATE_ITEM_REQUEST", quantity: 2} ] 
  }) {
    id
    total
    customer {
      id
      name
      email
    }
    items {
      id
      quantity
      item {
        id
        name
        price
      }
    }
  }
}

query findCustomers {
  customers {
    id
    name
    email
  }
}

query findItems {
  items {
    id
    name
    price
  }
}

query findOrders {
  orders {
    id
    total
    items {
      id
      quantity
      item {
        id
        name
        price
      }
    }
    customer {
      id
      name
    }
  }
}

query findOrderItems {
  orderItems {
    id
    order {
      id
      total
    }
    item {
      id
      name
      price
    }
    quantity
  }
}
```
