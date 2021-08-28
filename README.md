![Build Status](https://github.com/StuartsHome/carPedia/actions/workflows/go.yml/badge.svg)  
# wikipedia for cars

Under development, mostly on weekends

## How to use
### How to run
#### Local
- `go run main.go` for production
- `go run main.go -development=true` for development "testing"

#### Container
- `make docker-start`


### Logging
- logs are located in log/*

## Web Pages
- **home**: web page that provides html form for saving car details into db
- **allcars**: displays all the cars in the database in a html table
- **car**: displays all the cars in the database in JSON format

## Features
- MySQL Database
- Makefile w/ make commands
- Containerised
- Static web page (HTML, JS)
- Shell scripts
- Logging



# TO do
- all data endpoint (verify zone)
    - get request for all data in table