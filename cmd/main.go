package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  "os"
  _ "internal/config"
  // "github.com/gorilla/websocket"
  // "gorm.io/gorm"
  // "gorm.io/driver/postgres"
)

type Pipeline struct {
  Id int `json:"id"`
  Name string `json:"name"`
}

// get pipelines from db
func getPipelines (w http.ResponseWriter, r *http.Request) {
  var pipelines []Pipeline = []Pipeline {
    Pipeline {1,"Laravel-test"},
    Pipeline {2,"build"},
    Pipeline {3,"dockerise"},
  }
  b, err := json.Marshal(pipelines)
  
  if err != nil {
    log.Fatalf("failed to marshal json: %s", err)
  }
  fmt.Fprintf(w, `%s`, b)
  return
}

// get pipeline by id
func getPipeline (w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  fmt.Fprintf(w, `%s`, id)
  return
}

func authMiddleware (h http.HandlerFunc) {}

func getTestUser(w http.ResponseWriter, r *http.Request) {
  user, exists := os.LookupEnv("USER")
  if exists {
    fmt.Fprintf(w, "%v", user)
  }
  return
}

func main() {
  http.HandleFunc("GET /pipelines", getPipelines)
  http.HandleFunc("GET /pipelines/{id}", getPipeline)
  http.HandleFunc("GET /env/user", getTestUser)
  
  fmt.Printf("server is listening on port %v", 8080)
  http.ListenAndServe(":8080", nil)
}