// Package policy exposes interfaces and implementations of agent control.
package policy

import "learn/env"

// Controller types define a general policy of control for agents.
type Controller interface {
	// Returns the action to be taken from a state.
	Next(s env.State) env.Action
}

// Policy types define a behavioral policy of control for agents.
type Policy func(s env.State, l env.Valuator, b env.Behavior) env.Action
