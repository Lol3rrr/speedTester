package tester

import (
  "time"
  "net/http"
)

func makeGetRequest(url string) (time.Duration, error) {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return -1, err
  }

  start := time.Now()

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return -1, err
  }
  defer resp.Body.Close()

  return time.Since(start), nil
}
