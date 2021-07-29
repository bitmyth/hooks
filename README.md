# Introduction
A web server which receive hook request and run pre-defined jobs accordingly

# Run
FILE=$(pwd)/jobs.yaml go run src/main.go

# Run Demo job
curl  localhost:8000/test

# Job definition file
There is a demo definition file jobs.yaml