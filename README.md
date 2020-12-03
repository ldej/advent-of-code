# Advent of Code

https://adventofcode.com/

2017, 2018 and 2019 have been moved from individual repos to this repo. They might need tinkering to get working again.

2020 is up and running :smile:

## Generate a custom grid

Define a struct

```go
type Person struct {
    Name string
    Age int
}
```

Generate a grid for type `Person`
```
$ go run tools/gen/grid.go -name=Person -type=Person > 2020/day04/1/person_grid.go
```