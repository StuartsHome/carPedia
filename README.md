![Build Status](https://github.com/StuartsHome/carPedia/actions/workflows/go.yml/badge.svg)  
# wikipedia for cars

Under development, mostly on weekends

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
- add flags
- add context
- Logging needs improvement
    - log the port service is running on
    - log better errors
- More tests!!!
- cache (redis, elastic)



# Bug
- database - container uses different user and password credentials and required manual setup
this should be automated
- database - uri are different for local and container

Vscode:
Trigger suggest = cmd + i
Find with selection = cmd + e
Change focus = cmd + l

Atoi (string to int) and Itoa (int to string).