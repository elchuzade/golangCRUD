# RestAPI CRUD using Golang

A simple web application built using Golang to perform CRUD operations on a Mock Database (not persistent).

## Routes

- Show all cars - GET - api/cars
- Show one car - GET - api/cars/{id}
- Add new car - POST - api/cars
- Modify car - PUT - api/cars/{id}
- Remove car - DELETE - api/cars/{id}

## Libraries

I used mainly libraries that are included into Golang installation except for "mux" that can be installed from the command below
`go get github.com/gorilla/mux`

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
