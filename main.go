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
}
