# Kowala Style Guide

This guide covers Kowala programming style from aesthetic issues to conventions and coding standard.

We try to follow as much as possible Clean Code from uncle Bob


## Formatting

Style is easiest due to go gofmt. We follow gofmt style so make sure you gofmt your files first

`gofmt -s -w file.go`


### Column limit: 100

Our line length should be limited to 100, CI should fail if above that level


### File size lines: 200

File size should have a soft limit of 200 lines, hard limit being 500 lines, but these should be rare  


### Dead/commented out code

Commented code will have no meaning to other developers, will go out of date really quickly.

http://www.informit.com/articles/article.aspx?p=1334908


### Comments

Prioritize good code over comments, code should be self explanatory

Consider refactor if you need to explain what the code does in a comment block

Only exception is package documentation and library usage examples

Avoid obvious comments:

``go
// MarshalJSON implements the json.Marshaller interface.
func (u URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}
``

