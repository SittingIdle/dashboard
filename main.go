package main

import (
  "encoding/json"
  "net/http"
)

type Adriaan struct {
  Tsjah    string
  Deze []string}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  adriaan := map[string]interface{}{"Blah": "die", "Tsja": 4, "Kon": []string{"Blah", "die", "Tsja"}}
  

  js, err := json.Marshal(adriaan)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}