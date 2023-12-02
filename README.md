# Go Practices

Try the "new" stuff available in Go.


## Go 1.18

### Go Workspaces
[Go Dev Tutorial](https://go.dev/doc/tutorial/workspaces)

Create a new Go workspace with `go work init <path>`
```shell
go work init .
```

Add a new Go module to workspace:
```shell
go work use ./hello
go work use ./generics
```
Run the go module with `go run ./hello`
Run the go module with `go run ./generics`


### Generics
[Go Dev Tutorial](https://go.dev/doc/tutorial/generics)

1. Generic function which handles multiple types
```go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V
```

2. Generic function which handles type constraint as interface
```go
type Number interface {
    int64 | float64
}
func SumNumbers[K comparable, V Number](m map[K]V) V
```

### Fuzzing (for testing)
[Go Dev Tutorial](https://go.dev/doc/tutorial/fuzz)

Fuzz testing:
```go
func FuzzReverse(f *testing.F) {
    ...
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        ...
    })
}
```

Run test with fuzzing. Note it will not stop unless it encounter error or Ctrl+C
```shell
go test -fuzz=Fuzz
```
Creates fuzz test corpus entry in `testdata/fuzz/FuzzReverse`

Run with specific fuzz test case using:
```shell
go test -run=FuzzReverse/f6de80301b525abbb75c6df22db8b76ffac55fa3784e52c28e102b94e2a33bed
```

Run with time limit:
```shell
go test -fuzz=Fuzz -fuzztime 30s
```

## Go 1.20
```shell
go work edit -go=1.20
```

### DateFormat
[Medium Ref](https://betterprogramming.pub/changes-in-go-1-20-b0a82d4b6c44)

Instead of `.Format("2006–01–02")`, which is confusing, the date formatting is clearer now:

```go
now := time.Now()
fmt.Println(now.Format(time.DateOnly))
```

```go
now := time.Now()
fmt.Println(now.Format(time.TimeOnly))
```

### error.Join
[Medium Ref](https://betterprogramming.pub/changes-in-go-1-20-b0a82d4b6c44)

It's finally here! Better error grouping using `errors.Join`! No more errgroups!

```go
for i := 0; i < 10; i++ {
    err = errors.Join(err, writeTofile(), doOtherThing())
}
```

Check for specific error without joining!
```go
func do() error {
   err := doManyThings()
   if errors.Is(err, fs.ErrClosed) {
      return err
   }

   return nil
}
```
