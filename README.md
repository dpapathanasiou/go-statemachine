go-statemachine
===============

About
-----

This package provides an implementation of a finite state machine in Go (<a href="http://golang.org/">http://golang.org/</a>), inspired by David Mertz's article "<a href="http://www.ibm.com/developerworks/library/l-python-state/index.html">Charming Python: Using state machines</a>" (<a href="http://www.ibm.com/developerworks/library/l-python-state/index.html">http://www.ibm.com/developerworks/library/l-python-state/index.html</a>).

Usage
-----

Install the package in your environment:

```
$ go get github.com/dpapathanasiou/go-statemachine
```

Make sure the <a href="http://golang.org/doc/code.html#tmp_2">GOPATH</a> environment variable is defined correctly and test it:

```
$ go test github.com/dpapathanasiou/go-statemachine
ok  	github.com/dpapathanasiou/go-statemachine	0.014s
```

To use it within your own code, import "github.com/dpapathanasiou/go-statemachine" and instantiate the state machine with an empty map of handler functions, along with the string name of the initial state handler function.

Example
-------

See the statemachine_test.go file included in this package.