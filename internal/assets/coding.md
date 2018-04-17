# Kowala Style Guide

This guide covers Kowala programming style from aesthetic issues to conventions and coding standard.

## Test

All PR should be accompanied by unit tests

### Tools

### Unit tests
We use 

[Testify]([https://github.com/stretchr/testify) "A toolkit with common assertions and mocks that plays nicely with the standard library"

[Mockery](https://github.com/vektra/mockery) "A mock code autogenerator for golang"

Mocks should go into a subpackage of the interface that implements
to generate a Mock for a interface run:

```go
mockery -name InterfaceName
```

### E2E

[Godog](https://github.com/DATA-DOG/godog) "Cucumber for golang"


## General


### Avoid too many arguments

"Functions should have a small number of arguments. No argument is best, followed by one, two, and three. More than three is very questionable and should be avoided with prejudice."
(http://www.informit.com/articles/article.aspx?p=1375308)

This does not apply to constructor, but even when a constructor has to many arguments consider optional params or a builder class

Don't do this

```go
postLetter(string country, string town, string postcode, string streetAddress, int appartmentNumber, string careOf)
```

why not

```go
postLetter(Address address)
```

### Named returns

Go supports named returns, but they are discouraged from standard go style. They should be used on very small functions only.
Some are well known problems like shadowing. 

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
``` 

## Formatting


### Style

Style is easiest due to go gofmt. We follow gofmt style so make sure you gofmt your files first

`gofmt -s -w file.go`


### Column limit: 100

Line length should be limited to 100, CI should fail if above that level


### File size lines: 300

File size should have a soft limit of 300 lines, hard limit being 500 lines, but these should be rare  

### Error messages

As per go standard error messages should start with lower case

[Errors](https://github.com/golang/go/wiki/Errors)

User types should be used when you might expect the caller to do type assertion on type error example

```go
var ErrNotFound = errors.New("not found")

func findById(id int) (string, error) {
    name, err := searchPerson(id) 
	if err != nil {
		return "", err
	}
	
	if name = "" {
		return "", ErrNotFound
	}
	
	return name, nil
}
``` 

### Dead/commented out code

Commented code will have no meaning to other developers, will go out of date really quickly.

http://www.informit.com/articles/article.aspx?p=1334908

Don't do this

```go
function logNotEmpty(message String) {
  if message != "") {
    log(message)
  }
//      else {
//        log("no log message")
//      }
}
```


### Comments

Prioritize good code over comments, code should be self explanatory

Consider refactor if you need to explain what the code does in a comment block

Only exception is package documentation and library usage examples

