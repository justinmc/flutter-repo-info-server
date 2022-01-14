package main

import (
  "fmt"
  "github.com/justinmc/flutter-repo-info/server"
  "log"
  "net/http"
)

const kPort string = "8080";

func main() {
  fmt.Println("Server listening on", kPort, "...");
  http.HandleFunc("/pr/", server.RoutePR)
  log.Fatal(http.ListenAndServe(":" + kPort, nil))
}
