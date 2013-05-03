
// This is an example of how to use the statemachine pkg, based on statemachine_test.py
// from http://www.ibm.com/developerworks/library/l-python-state/index.html

// To quote from the explanation on that page:

// "The cargo is just a number that keeps getting fed back into an iterative function.
// The next value of val is always simply do_math(), as long as val stays within a certain range. 
// Once the function returns a value outside that range, either the value is passed to a 
// different handler or the state machine exits after calling a do-nothing end-state handler. 
// One thing the example illustrates is that an event is not necessarily an input event. 
// It can also be a computational event (atypically). 
// The state-handlers differ from one another only in using a different marker when outputting 
// the events they handle. This function is relatively trivial, and does not require using a state machine. 
// But it illustrates the concept well."

package statemachine_test

import (
    "fmt"
    "math"
    "github.com/dpapathanasiou/go-statemachine"
	"testing"
)

func do_math(i float64) float64 {
    return math.Abs(math.Sin(i) * 31.0)
}

func ones_counter() statemachine.Handler {
    return func(val interface{}) (nextState string, nextVal interface{}) {
        nextState = ""
        nextVal = val.(float64)

        fmt.Printf("1s State:\t")
        for {
            switch {
            case (nextVal.(float64) <= 0 || nextVal.(float64) >= 30):
                nextState = "outofrange"
            case (nextVal.(float64) >= 20 && nextVal.(float64) < 30):
                nextState = "twenties"
            case (nextVal.(float64) >= 10 && nextVal.(float64) < 20):
                nextState = "tens"
            default:
                fmt.Printf(" @ %2.1f+", nextVal.(float64))
            }

            if len(nextState) > 0 {
                break
            }
            nextVal = do_math(nextVal.(float64))
        }
        fmt.Printf(" >>\n")
        return
    }
}

func tens_counter() statemachine.Handler {
    return func(val interface{}) (nextState string, nextVal interface{}) {
        nextState = ""
        nextVal = val.(float64)

        fmt.Printf("10s State:\t")
        for {
            switch {
            case (nextVal.(float64) <= 0 || nextVal.(float64) >= 30):
                nextState = "outofrange"
            case (nextVal.(float64) >= 20 && nextVal.(float64) < 30):
                nextState = "twenties"
            case (nextVal.(float64) >= 1 && nextVal.(float64) < 10):
                nextState = "ones"
            default:
                fmt.Printf(" #%2.1f+", nextVal)
            }

            if len(nextState) > 0 {
                break
            }
            nextVal = do_math(nextVal.(float64))
        }
        fmt.Printf(" >>\n")
        return
    }
}

func twenties_counter() statemachine.Handler {
    return func(val interface{}) (nextState string, nextVal interface{}) {
        nextState = ""
        nextVal = val.(float64)

        fmt.Printf("20s State:\t")
        for {
            switch {
            case (nextVal.(float64) <= 0 || nextVal.(float64) >= 30):
                nextState = "outofrange"
            case (nextVal.(float64) >= 10 && nextVal.(float64) < 20):
                nextState = "tens"
            case (nextVal.(float64) >= 1 && nextVal.(float64) < 10):
                nextState = "ones"
            default:
                fmt.Printf(" *%2.1f+", nextVal)
            }

            if len(nextState) > 0 {
                break
            }
            nextVal = do_math(nextVal.(float64))
        }
        fmt.Printf(" >>\n")
        return
    }
}

func TestStateMachine(t *testing.T) {
    m := statemachine.Machine{map[string]statemachine.Handler{}, "ones", map[string]bool{}}
    m.AddState("ones", ones_counter())
    m.AddState("tens", tens_counter())
    m.AddState("twenties", twenties_counter())
    m.AddEndState("outofrange")

    m.Execute(1.0)
}
