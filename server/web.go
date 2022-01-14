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
  var sha string = r.URL.Path[len("/pr/"):]

  fmt.Println("justin sha", sha);
  if (sha == "") {
    // TODO(justinmc): Error response code in error states like this one, and
    // 200's in success states.
    fmt.Fprintf(w, "No valid SHA given.");
    return;
  }

  isInStable, err := api.IsInStable(sha); // In stable.
  //isInStable, err := api.IsInStable("8b46014"); // Not in stable (yet).
  if (err != nil) {
    fmt.Fprintf(w, "No valid SHA given.");
    return;
  }

  if (isInStable) {
    fmt.Fprint(w, "Yes");
    return;
  }
  fmt.Fprint(w, "No");
  return;
}
