# GO Rest API

![N|Solid](https://www.logolynx.com/images/logolynx/85/8591214be4b4d42924c438a9a4979947.png)

This is a Rest API made in Go in order to learn and practice more this great and cool language. Is a simple API that allows to manage supermarket products.

The current routes are enable:

| Route   | URL           | HTTP Verb | Description                                |
| ------- | ------------- | --------- | ------------------------------------------ |
| Index   | /products     | GET       | Returns all products                       |
| Create  | /products     | POST      | Add a new product and returns all products |
| Show    | /products/:id | GET       | Return a product of an specific id         |
| Update  | /products/:id | PUT       | Update a product by id and return it       |
| Destroy | /products/:id | DELETE    | Delete a product by id                     |

This API consumes a local slice that contains a couple products.

### 3rd Party libraries used:

- gorilla/mux
  `go get -u github.com/gorilla/mux`
