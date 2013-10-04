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
	"io/ioutil"
	"testing"
)

func assertParsed(
	t *testing.T, expected Rss2Document, testFile string) {
	xml, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	actual, err := Parse(xml)
	if err != nil {
		t.Errorf("%v", err)
	}
	var doctests = []struct {
		exp string
		act string
	}{
		{expected.Title, actual.Title},
		{expected.Link, actual.Link},
	}
	for i, doctest := range doctests {
		if doctest.exp != doctest.act {
			t.Errorf(
				"%v doctest %d: wanted %v got %v",
				testFile, i, doctest.exp, doctest.act)

		}
	}
	if len(expected.Items) != len(actual.Items) {
		t.Errorf(
			"%v wrong number of items: wanted %d got %d",
			testFile, len(expected.Items), len(actual.Items))
		return
	}
	for i, actualItem := range actual.Items {
		if expected.Items[i] != actualItem {
			t.Errorf(
				"%v item %d malformed: wanted %v got %v",
				testFile, i, expected.Items[i], actualItem)
		}
	}
}

func TestParse_Rss2(t *testing.T) {
	var expected = Rss2Document{
		Title: "Forget the Milk",
		Link:  "http://forgetthemilk.tumblr.com/",
		Items: []Rss2Item{
			{
				"Santigold sings in a carefree mix of a lot of musical styles....",
				"http://forgetthemilk.tumblr.com/post/53876831191",
				"<iframe width=\"400\" height=\"299\" src=\"http://www.youtube.com/embed/xQ94krsqFHg?wmode=transparent&autohide=1&egm=0&hd=1&iv_load_policy=3&modestbranding=1&rel=0&showinfo=0&showsearch=0\" frameborder=\"0\" allowfullscreen></iframe><br/><br/><p><a href=\"https://en.wikipedia.org/wiki/Santigold\">Santigold</a> sings in a carefree mix of a lot of musical styles. Reminds me of Neko Case, St. Vincent, or Ida Maria depending on the song. [<a href=\"https://www.youtube.com/watch?v=xQ94krsqFHg\">direct link to video</a>]</p>",
				"Tue, 25 Jun 2013 18:32:24 -0400",
			},
			{
				"First Aid Kit sings gentle blasphemous gospel, like the best of...",
				"http://forgetthemilk.tumblr.com/post/38646165851",
				"<iframe width=\"400\" height=\"225\" src=\"http://www.youtube.com/embed/gekHV9DIjHc?wmode=transparent&autohide=1&egm=0&hd=1&iv_load_policy=3&modestbranding=1&rel=0&showinfo=0&showsearch=0\" frameborder=\"0\" allowfullscreen></iframe><br/><br/><p><a href=\"http://en.wikipedia.org/wiki/First_Aid_Kit_(band)\">First Aid Kit</a> sings gentle blasphemous gospel, like the best of “Rabbit Fur Coat” by Jenny Lewis and the Watson Twins. [<a href=\"http://www.youtube.com/watch?v=gekHV9DIjHc\">direct link to video</a>]</p>",
				"Sun, 23 Dec 2012 14:36:27 -0500",
			},
			{
				"Hans-Joachim Roedelius writes piano-heavy instrumental music. Is...",
				"http://forgetthemilk.tumblr.com/post/37137021414",
				"<iframe width=\"400\" height=\"225\" src=\"http://www.youtube.com/embed/W70PJuoQXJs?wmode=transparent&autohide=1&egm=0&hd=1&iv_load_policy=3&modestbranding=1&rel=0&showinfo=0&showsearch=0\" frameborder=\"0\" allowfullscreen></iframe><br/><br/><p><a href=\"http://en.wikipedia.org/wiki/Hans-Joachim_Roedelius\">Hans-Joachim Roedelius</a> writes piano-heavy instrumental music. Is nice to work to. [<a href=\"https://www.youtube.com/watch?v=W70PJuoQXJs\">direct link to video</a>]</p>",
				"Mon, 03 Dec 2012 17:09:27 -0500",
			},
		},
	}
	assertParsed(t, expected, "testdata/rss20.rss")
}
