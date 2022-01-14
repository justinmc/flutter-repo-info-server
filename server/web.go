package server

import (
  "fmt"
  "github.com/justinmc/flutter-repo-info/api"
  "net/http"
)


// TODO(justinmc):
//  * Handle only GET.
//  * Return all the info a web page would need:
//    - Merge SHA, if merged.
//    - What channels it's in.
//    - Return a response, don't just log.
//  * Currently takes a SHA, but it should take a PR number and find the SHA itself.
func RoutePR(w http.ResponseWriter, r *http.Request) {
  var prNumber string = r.URL.Path[len("/pr/"):]
  if (prNumber == "") {
    // TODO(justinmc): Error response code in error states like this one, and
    // 200's in success states.
    fmt.Fprintf(w, "No valid PR given.");
    return;
  }

  sha, err := api.GetPrMergeCommit(prNumber);
  if (err != nil) {
    fmt.Fprintf(w, "Couldn't find merge commit for given PR %s.", prNumber);
    return;
  }

  if (sha == "") {
    fmt.Fprintf(w, "Couldn't find merge commit for given PR %s.", prNumber);
    return;
  }

  isInStable, err := api.IsInStable(sha);
  if (err != nil) {
    fmt.Fprintf(w, "Couldn't determine if the PR %s's merge %s commit was in stable.", prNumber, sha);
    return;
  }

  if (isInStable) {
    fmt.Fprint(w, "Yes");
    return;
  }
  fmt.Fprint(w, "No");
  return;
}
