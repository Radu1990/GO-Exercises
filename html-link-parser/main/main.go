package main

import (
	"fmt"
	link "github.com/Radu1990/html-link-parser"
	"strings"
)

/*
	In this exercise the goal is create a package
	that makes it easy to parse an HTML file and
	extract all of the links (<a href="">...</a> tags).
	For each extracted link you should return a data
	structure that includes both the href, as well as
	the text inside the link.
	Any HTML inside of the link can be stripped out,
	along with any extra whitespace including newlines,
	back-to-back spaces, etc.
	Links will be nested in different HTML elements, and
	it is very possible that you will have to deal with
	HTML similar to code below.

	<a href="/dog">
	  <span>Something in a span</span>
	  Text not in a span
	  <b>Bold text!</b>
	</a>

	In situations like these we want to get output that looks roughly like:

	Link{
	  Href: "/dog",
	  Text: "Something in a span Text not in a span Bold text!",
	}

 */
var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
    A link to another page
    <span> some span  </span>
  </a>
  <a href="/page-two">A link to a second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
