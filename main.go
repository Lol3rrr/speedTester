package main

import (
  "fmt"
  "flag"

  "speedTester/tester"
)

func main() {
  urlPtr := flag.String("url", "http://localhost", "The URL to speedtest")
  triesPtr := flag.Int("tries", 1000, "The Amount of tries to get the average from")
  verbosePtr := flag.Bool("verbose", false, "Whether or not it should output each Request")
  flag.Parse()


  fmt.Printf("--------------------\n")
  fmt.Printf("Starting Test: \n")
  fmt.Printf("URL: '%s' \n", *urlPtr)
  fmt.Printf("Tries: %d \n", *triesPtr)
  fmt.Printf("Verbose: %v \n", *verbosePtr)
  fmt.Printf("--------------------\n\n")


  maxTime, avgTime, minTime := tester.MeasureTimes(*triesPtr, *urlPtr, *verbosePtr)

  fmt.Printf("\n")
  fmt.Printf("--------------------\n")
  fmt.Printf("Results: \n")
  fmt.Printf("Max: %v \n", maxTime)
  fmt.Printf("Min: %v \n", minTime)
  fmt.Printf("Average: %v \n", avgTime)
  fmt.Printf("--------------------\n")
}
