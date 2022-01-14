package api

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "log"
  "net/http"
)

const kAPI string = "https://api.github.com/repos/flutter/flutter";
const kBranch string = "master";

// TODO(justinmc): This should be scraped from https://docs.flutter.dev/development/tools/sdk/releases.
const kLatestStable string = "77d935a";
const kLatestBeta string = "628f0e3";


type PR struct {
  MergeCommitSha string `json:"merge_commit_sha"`
  MergedAt string `json:"merged_at"`
  Base Branch `json:"base"`
}

type Branch struct {
  Ref string `json:"ref"`
}

type Comparison struct {
  Status string `json:status`
}

// Returns true iff the given sha is on the current stable channel.
//
// This doesn't consider reverts.
func IsInStable(sha string) (bool, error) {
  //  http get https://api.github.com/repos/flutter/flutter/compare/77d935a...f3947ea
  resp, err := http.Get(kAPI + "/compare/" + kLatestStable + "..." + sha);
  if (err != nil) {
    return false, err;
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return false, err;
  }
  bodyString := string(body)

  var comparison Comparison;
  json.Unmarshal([]byte(bodyString), &comparison)

  if (comparison.Status == "") {
    return false, errors.New("Unable to find the given SHA.");
  }

  return comparison.Status == "behind" || comparison.Status == "identical", nil;
}

// Returns the SHA for the merge commit of the given PR.
//
// Returns an error if the PR doesn't exist, isn't merged, or isn't based on
// master.
func GetPrMergeCommit(prNumber string) (string, error) {
  resp, err := http.Get(kAPI + "/pulls/" + prNumber);
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
  if (pr.Base.Ref != kBranch) {
    return "", errors.New("PR's base isn't master");
  }

  return pr.MergeCommitSha, nil;
}
