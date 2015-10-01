# Templates for Gen

This repository contains a set of templates that can be substitute by [gen](http://clipperhouse.github.io/gen/).
command line tool. Presently, it contains the following templates:

- [Stack](http://bit.ly/1Pwvd5W) data struct. Go documentation [here](https://godoc.org/github.com/svett/gen/stack).

## Requirments

- Go
- [Gen](http://clipperhouse.github.io/gen/)

## Installation

```
$ go get clipperhouse.github.io/gen
$ gen add github.com/svett/gen/stack
```

## Usage

```
// +gen * stack
type Student struct {
	FirstName string
	LastName  string
	BirthDate time.Time
}
```

```
stack := &StudentStack{}
stack.Push(&Student{FirstName: "John", LastName: "Smith"})
```

```
$ gen
```

## Author

Svett Ralchev - http://blog.ralch.com
