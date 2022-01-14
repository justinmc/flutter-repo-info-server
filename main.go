package main

import (
  "fmt"
  "github.com/justinmc/flutter-repo-info/api"
  "log"
  "net/http"
)

func main() {
  /*
  fmt.Println("Starting server...")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
  */

  //var merged bool = prWasMerged(96323); // Not merged yet, draft.
  //var merged bool = prWasMerged(95948); // Merged.
  //fmt.Println("justin merged?", merged);

  /*
  mergeCommit, err := api.GetPrMergeCommit(95948);
  if (err != nil) {
    log.Fatalln(err);
  }

  fmt.Println("justin merge commit", mergeCommit);
  */

  isInStable, err := api.IsInStable("f3947ea"); // In stable.
  //isInStable, err := api.IsInStable("8b46014"); // Not in stable (yet).
  if (err != nil) {
    log.Fatalln(err);
  }
  fmt.Println("justin is in stable?", isInStable);
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
