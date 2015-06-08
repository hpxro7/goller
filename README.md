goller
=======
Reinforcement Learning methods in Go for optimal control via generalized policy iteration.

# Example
```go
greedy := policy.EpisilonGreedy(eps)
learner := learn.WithPolicy(greedy, behavior).Q(learnRate)
```
