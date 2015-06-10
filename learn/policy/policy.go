// Package policy exposes interfaces and implementations of agent control.
package policy

import (
	"errors"
	"math/rand"

	"github.com/hpxro7/goller/learn/env"
)

var (
	// No valid actions can be taken.
	NoActions error = errors.New("No valid actions can be taken.")
)

// Controller types define a general policy of control for agents.
type Controller interface {
	// Returns the action to be taken from a state.
	Next(s env.State) env.Action
}

// Policy types define a behavioral policy of control for agents.
type Policy func(s env.State, l env.Valuator, b env.Behavior) (env.Action, error)

// EpsilonGreedy returns an episilon-soft class Policy which chooses a non-greedy action
// with probabilty eps and a greedy action with probability 1 - eps.
// If there is only one valid action, that action is taken irrespective of greediness.
func EpsilonGreedy(eps float32, r rand.Rand) Policy {
	return func(s env.State, v env.Valuator, b env.Behavior) (next env.Action, err error) {
		greedy := r.Float32() >= eps

		as := b.Actions(s)
		if len(as) == 0 {
			err = NoActions
			return
		}

		// Find action that maximizes action value
		aMax, qMax := as[0], v.ActionValue(s, as[0])
		for _, a := range as[1:] {
			q := v.ActionValue(s, a)
			if q > qMax {
				aMax, qMax = a, q
			}
		}

		// Pick action with greatest expected return
		if greedy || len(as) == 1 {
			next = aMax
			return
		} else {
			// Choose a non-greedy action uniformly at random
			next = aMax
			return
		}
	}

}
