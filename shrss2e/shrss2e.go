/*
Copyright 2013 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package shrss2e

import (
    "fmt"
    "net/http"
    rss "github.com/jteeuwen/go-pkg-rss"

    "appengine"
    "appengine/user"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }
    feed := rss.New(60, true, chanHandler, itemHandler)
    if err := feed.Fetch(c, "http://livemylief.com/rss", nil); err != nil {
        fmt.Fprintf(w, "[e]: %s", err)
        return
    }

    fmt.Fprintf(w, "Hello, %v! %d", u, feed.SecondsTillUpdate())
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
    fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
    fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
}
