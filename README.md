go-statemachine
===============

About
-----

This package provides an implementation of a finite state machine in Go (http://golang.org/), inspired by David Mertz's article "Charming Python: Using state machines" (http://www.ibm.com/developerworks/library/l-python-state/index.html).

Usage
-----

Install the package in your environment:

```
go get github.com/dpapathanasiou/go-statemachine
```

To use it within your own code, import "github.com/dpapathanasiou/go-statemachine" and instantiate the state machine with an empty map of handler functions, along with the string name of the initial state handler function.

Example
-------

See the statemachine-test.go file included in this package.