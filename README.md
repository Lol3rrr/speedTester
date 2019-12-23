# speedTester
A small program to test the speed of single HTTP-Requests

## What is it?
The SpeedTester is basically just a small program that can send X amount of requests to a specific HTTP-Endpoint and measures the response times.
I wrote this tool initially to test the performance of my own webserver.

## Usage
### Running/Compiling the Code yourself
* Clone the repo into your own GOPATH
* Run the main.go file using `go run main.go` and it should just work

### Using the binary
* Download the newest Binary for your system
* Go to the Directory where you saved it
* Run the Binary

## Arguments
* url: The URL you want to test
* tries: The Amount of requests that should be send
* verbose: Whether or not it should output the result of every request
