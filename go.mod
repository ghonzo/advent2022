module github.com/ghonzo/advent2022

go 1.19

// For generics constraints
require golang.org/x/exp v0.0.0-20221126150942-6ab00d035af9

// Easier JSON parsing for leaderboard.go
require (
	github.com/tidwall/gjson v1.14.4
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
)

// Priorirty queue (Used in day 12)
require github.com/oleiade/lane/v2 v2.0.0
