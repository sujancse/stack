## Installation

```
go get github.com/sujancse/stack
```

## Usage

### Initialisation
```go
s := stack.Stack[int]{} // initialised stack with int data type
s1 := stack.Stack[float64]{} // initialised stack with float64 data type
```

### Push
Stack initialised with data type int put integer data and for float64 push float data
```go
s.Push(1)
s.Push(2)

s1.Push(1.00)
s1.Push(2.00)
```

### Pop
Get the top data and remove from stack
```go
s.Pop()
```

### Peek
Get the top element from the stack
```go
s.Peek()
```

### IsEmpty
Check whether the stack has data or empty
```go
s.IsEmpty()
```

## Usage in GoDoc ##

[See GoDoc](http://godoc.org/github.com/sujancse/stack) for
documentation and examples.

## Examples ##

See [GoDoc](http://godoc.org/github.com/sujancse/stack) or
`example_test.go` for examples.

## Contributing ##

Contributions welcome! Please fork the repository and open a pull request
with your changes.

## License ##

This is free software, licensed under the MIT License.
