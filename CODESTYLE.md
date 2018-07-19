# Kowala Style Guide

This guide covers Kowala programming style from aesthetic issues to conventions and coding standard.

## Test

All PRs should be accompanied by unit tests and end-to-end test.

### Tools

### Unit tests

We use :

[Testify]([https://github.com/stretchr/testify) "A toolkit with common assertions and mocks that plays nicely with the standard library".

[Mockery](https://github.com/vektra/mockery) "A mock code autogenerator for golang"

Mocks should go into a subpackage of the interface that it implements.

To generate a Mock for a interface run:

```
mockery -name AddressVote
```

And to use the mock: 

```go
addressVote := &mocks.AddressVote{}
addressVote.On("Vote").Return("yes")
addressVote.On("Address").Return("home")
```

### Test fixtures

We use Go idiomatic [Golden Files](https://medium.com/soon-london/testing-with-golden-files-in-go-7fccc71c43d3), to keep our test fixtures up to date.

All files should have a suffix `.golden` and update flag for tests should be `--update`.

example:

```go
var update = flag.Bool("update", false, "update .golden files")
func TestSomething(t *testing.T) {
  actual := functionUnderTest()
  golden := filepath.Join("testfiles", tc.Name+".golden")
  if *update {
    ioutil.WriteFile(golden, actual, 0644)
  }
  expected, _ := ioutil.ReadFile(golden)
  assert.Equal(t, actual, expected)
}
```


### E2E

[Godog](https://github.com/DATA-DOG/godog) "Cucumber for golang"


## General

### Valid Objects
 
Create valid objects at construction time, use several constructors to have variants of the object, but all should be valid.

Valid here means, that further call should be needed to use the object.

Don't do this

```go
email := NewEMail()
email.SetTo("kcoin@kowala.tech")
``` 
the developer might not me aware that it has to set de TO before using EMail.

If you require and IP why not:

```go
email := NewMailer("kcoin@kowala.tech")
```

### Avoid too many arguments

"Functions should have a small number of arguments. No argument is best, followed by one, two, and three. More than three is very questionable and should be avoided with prejudice."
(http://www.informit.com/articles/article.aspx?p=1375308)

Consider using struct of values, optional params or a builder.

Don't do this:

```go
postLetter(firstName string, lastName string, street string, city string, postcode string, flatNumber int)
```

why not:

```go
postLetter(personName PersonName, address Address)
```

### Readability over "smart" code

Write code for humans first, try to express the intent with function and variable names.

From:
```go
currentBlock := val.chain.CurrentBlock()
if currentBlock.Number().Cmp(big.NewInt(0)) == 0 {
    return
}
```

To:
```go
currentBlock := val.chain.CurrentBlock()
if isFirstBlock(currentBlock) {
    return
}

func isFirstBlock(block Block) bool {
    return block.Number().Cmp(big.NewInt(0)) == 0
}
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

### Interfaces

New concepts should always be expressed in interfaces first.

## Formatting


### Style

Style is easiest due to go gofmt. We follow gofmt style so make sure you gofmt your files first.

`gofmt -s -w file.go`


### Column limit: 100

Line length should be limited to 100, this is not enforced by CI at the moment.


### File size lines: 300

File size should have a soft limit of 300 lines, hard limit being 500 lines, but these should be rare.  

### Error messages

As per go standard error messages should start with lower case.

[Errors](https://github.com/golang/go/wiki/Errors)

User types should be used when you might expect the caller to do type assertion on type error example.

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

Don't do this:

```go
func logNotEmpty(message String) {
  if message != "") {
    log(message)
  }
//   else {
//      log("no log message")
//   }
}
```


### Comments

Prioritize good code over comments, code should be self explanatory.

Consider refactor if you need to explain what the code does in a comment block.

Only exception is package documentation and library usage examples.
