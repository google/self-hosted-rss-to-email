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
package rss

import (
	"encoding/xml"
)

// RSS 2.0 Spec: http://cyber.law.harvard.edu/rss/rss.html
type Rss2Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type Rss2Document struct {
	XMLName xml.Name `xml:"rss"`

	Title string     `xml:"channel>title"`
	Link  string     `xml:"channel>link"`
	Items []Rss2Item `xml:"channel>item"`
}

func Parse(s []byte) (doc Rss2Document, err error) {
	err = xml.Unmarshal(s, &doc)
	return
}
