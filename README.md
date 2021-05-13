# performer

An alternative way of handling errors in Go.

A thorough text discussing the idea behind this library can be read on [dev.to](https://dev.to/hlmerscher/the-controversial-go-way-of-handling-errors-2ka1).

## Examples

A simple case:

```go
func answer42(number int) performer.Task {
	return func() error {
		if number == 42 {
			return nil
		}
		return fmt.Errorf("the answer for everything is 42, not %d", number)
	}
}

func main() {
	err := performer.Do(
		answer42(42),
		answer42(1),
		answer42(2),
		answer42(3),
	)
	if err != nil {
		panic(err)
	}
}

// the answer for everything is 42, not 1
```

Structs could be used to store context when dealing with dependent actions:

```go
type Accumulator struct {
	Value *Value
}

func (i *Accumulator) Inc() error {
	i.Value.Value++
	return nil
}

type Value struct{ Value int }

type Multiply struct {
	Value      *Value
	Multiplier int
}

func (m *Multiply) Do() error {
	m.Value.Value = m.Value.Value * m.Multiplier
	return nil
}

type Divide struct {
	Value   *Value
	Divisor int
}

func (m *Divide) Do() error {
	if m.Divisor == 0 {
		return errors.New("can't divide by zero")
	}
	m.Value.Value = m.Value.Value / m.Divisor
	return nil
}

func main() {
	acc := &Accumulator{&Value{0}}
	double := &Multiply{Multiplier: 2, Value: acc.Value}
	divide := &Divide{Divisor: 0, Value: double.Value}
	err := performer.Do(
		acc.Inc,
		acc.Inc,
		acc.Inc,
		double.Do,
		divide.Do,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(divide.Value)
}

// can't divide by zero
// exit status 1
```
