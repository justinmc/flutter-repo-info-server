package api

import (
  "encoding/json"
  "errors"
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

// Returns the SHA for the merge commit of the given PR.
//
// Returns an error if the PR doesn't exist or isn't merged.
func GetPrMergeCommit(prNumber int) (string, error) {
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

  if (pr.MergeCommitSha == "") {
    return "", errors.New("PR doesn't exist.");
  }
  if (pr.MergedAt == "") {
    return "", errors.New("PR not yet merged.");
  }

  return pr.MergeCommitSha, nil;
}
