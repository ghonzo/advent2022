# advent2022
Advent of Code 2022 in Go

Each day is its own subdirectory, in the file `main.go`. Usually, the input for the day will be in a file called `input.txt`. Additionally,
the example data is usually in a file called `testdata/example.txt` and is referenced from the unit tests.

There is also a set of common utility functions and types in the `common` package.

There is a command `leaderboard.go` to print out the time it took everyone on a private leaderboard to get the first and second part of a particular day.

```
Usage of leaderboard.go:
  -day int
    	day to display, or most recent if not provided
  -endpoint string
    	URL of the leaderboard JSON endpoint. Can also set the LEADERBOARD_URL env variable.
  -session string
    	session cookie value. Can also set the LEADERBOARD_SESSION env variable.
```

Example:

```
$ go run leaderboard.go -day 3 -endpoint https://adventofcode.com/2022/leaderboard/private/view/123456.json -session 53616c7465645f5f89e40247fe925752de8a696d76c2f03f361ed6ea24283

Day 3                |      Part 1 |      Part 2
------------------------------------------------
Buddy The Elf        |       7m38s |      14m10s
Bonzo                |      12m15s |      17m23s
Stephen Peach        |      11m14s |      18m27s
Tony Tiger           |       19m9s |      27m34s
kenny                |      29m22s |      55m46s
Bender               |     9h35m0s |    10h7m49s
Mouse                |   11h41m55s |   12h33m28s
```
