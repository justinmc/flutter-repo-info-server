package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "strconv"
)

const kAPI = "https://api.github.com/repos/flutter/flutter";

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
  /*
  fmt.Println("Starting server...")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
  */

  //var merged bool = prWasMerged(96323); // Not merged yet, draft.
  var merged bool = prWasMerged(95948); // Merged.
  //var merged bool = prWasMerged(95948888); // Doesn't exist.
  fmt.Println("justin merged?", merged);
}

// Returns true if the PR has been merged, false otherwise.
//
// Even if the PR doesn't exist, will simply return false.
func prWasMerged(prNumber int) bool {
  resp, err := http.Get(kAPI + "/pulls/" + strconv.Itoa(prNumber) + "/merge");
  if (err != nil) {
    log.Fatalln(err);
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  bodyString := string(body)
  //fmt.Println("justin success", bodyString);

  return len(bodyString) == 0;
}
