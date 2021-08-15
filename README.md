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


## Features
- MySQL Database
- Makefile w/ make commands
- Containerised
- Static web page (HTML, JS)
- Shell scripts
- Logging


### To do :wrench:
- Update readme w/make commands
- update makefile
    - create server with a controllable name, not "carpedia_server_run_e95dd6243e95"
    - make docker-start should start the service in detached mode
    - make docker-start, service doesn't run on a port
- update database w/ update and delete commands
    - stats for db
- update frontend to look better
    - frontend includes button to get all results from db
    - button dynamically populates a table
- service should use concurrency to query the db
    - frontend should display a graph of metrics
        - how many workers, how many free workers
        - average time per worker
- database should have more tables
    - Database tests should be integration tests
- add flags
- add context
- Logging needs improvement
    - log the port service is running on
    - log better errors
- More tests!!!
- cache (redis, elastic)
- Table driven tests for database (not handlers)



# Bug
- database - container uses different user and password credentials and required manual setup
this should be automated
- database - uri are different for local and container

##### Vscode:
Trigger suggest = cmd + i
Find with selection = cmd + e
Change focus = cmd + l
Search for function/method name = cmd + t

Atoi (string to int) and Itoa (int to string).

Composition to instantiate instead of Inheritance
Golang uses composition.

1. Object composition: Object composition is used instead of inheritance (which is used in most of the traditional languages). Object composition means that an object contains object of another object (say Object X) and delegate responsibilities of object X to it. Here instead of overriding functions (as in inheritance), function calls delegated to internal objects.
2. Interface composition: In interface composition and interface can compose other interface and have all set of methods declared in internal interface becomes part of that interface.


There is a book called:
Design Patterns: Elements of Reusable Object-Oriented Software
In that, it speaks about the gang of 4, and to favour composition over inheritance

Embedding provides automatic delegation. This in itself isn't enough to replace inheritance, as embedding provides no form of polymorphism. Go interfaces do provide polymorphism, they are a bit different than the interfaces you may be use to (some people liken them to duck typing or structural typing).

polymorphism via interfaces
There is no inheritance in Go, so leave those is-a relationships at the door. To write Go, we need to think about OO design in terms of composition.