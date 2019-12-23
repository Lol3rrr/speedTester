package tester

import (
  "fmt"
  "time"
)

func MeasureTimes(tries int, url string, verbose bool) (max, average, min time.Duration) {
  var total time.Duration = 0

  var working time.Duration = 0
  for i := 0; i < tries; i++ {
    duration, err := makeGetRequest(url)
    if err != nil {
      continue
    }

    if verbose {
      fmt.Printf("[Request] Took %v \n", duration)
    }

    total += duration
    working++

    if duration > max {
      max = duration
    }
    if duration < min || min <= 0 {
      min = duration
    }
  }

  if working == 0 {
    return -1, -1, -1
  }

  average = total / working

  return
}
