package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "strconv"
)

const kAPI = "https://api.github.com/repos/flutter/flutter";

type PR struct {
  MergeCommitSha string `json:"merge_commit_sha"`
  MergedAt string `json:"merged_at"`
}

func main() {
  /*
  fmt.Println("Starting server...")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
  */

  //var merged bool = prWasMerged(96323); // Not merged yet, draft.
  //var merged bool = prWasMerged(95948); // Merged.
  //fmt.Println("justin merged?", merged);

  mergeCommit, err := getPrMergeCommit(95948);
  if (err != nil) {
    log.Fatalln(err);
  }

  fmt.Println("justin merge commit", mergeCommit);
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func getPrMergeCommit(prNumber int) (string, error) {
  resp, err := http.Get(kAPI + "/pulls/" + strconv.Itoa(prNumber));
  if (err != nil) {
    log.Fatalln(err);
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  bodyString := string(body)

  var pr PR;
  json.Unmarshal([]byte(bodyString), &pr)

  //fmt.Println("justin pr", pr);

  if (pr.MergedAt == "") {
    return "", errors.New("PR not yet merged.");
  }

  return pr.MergeCommitSha, nil;
}
