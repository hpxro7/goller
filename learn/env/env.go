// Package env prodivdes an interface for interacting with and querying state related to the
// environment of a problem modelled by a markov decision process(MDP).
package env

// Behavior types define the state and action space of a MDP.
type Behavior interface {
	// States returns all possible valid states of an MDP.
	States() []State
	// Actions returns all valid actions that can be taken from a given state.
	Actions(s State) []Action
	// ValidAction determines if Action a can be taken from State s.
	ValidAction(s State, a Action) bool
	// ValidState determies if State s is a valid state in this MDP.
	ValidState(s State) bool
}

type Action int

// State types define a state along with the reward and time-step that is associated with
// entering that state.
// Note that two State instances can have a different time-step and reward but be cannonically
// equivalent states.
type State interface {
	// Time returns the time-step of this State
	Time() int
	// Reward returns the reward gained when this state was entered.
	Reward() float32
	// Key returns a unique integer that determines equality for different states.
	Key() int
}

// Valuator defines types capable of computing state-value and action-value functions.
type Valuator interface {
	ActionValue(s State, a Action) float32
	StateValue(s State) float32
}
