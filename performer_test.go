package performer

import (
	"errors"
	"testing"
)

func alwaysSucceed() error {
	return nil
}

func alwaysFail() error {
	return errors.New("Boom!")
}

type Accumulator struct {
	Counter int
}

func (a *Accumulator) Inc() error {
	a.Counter += 1
	return nil
}

func (a *Accumulator) Dumb() error {
	return errors.New("I don't know how to increment")
}

func TestDo(t *testing.T) {
	t.Run("performs a single task", func(t *testing.T) {
		if err := Do(alwaysSucceed); err != nil {
			t.Errorf("error not expected, got %s", err)
		}
		if err := Do(alwaysFail); err == nil {
			t.Error("error expected and not thrown")
		}
	})

	t.Run("performs multiple succesful tasks in order", func(t *testing.T) {
		acc := &Accumulator{}
		err := Do(
			acc.Inc,
			acc.Inc,
			acc.Inc,
		)
		if err != nil {
			t.Errorf("error not expected, got %s", err)
		}
		expectedCounter := 3
		if acc.Counter != expectedCounter {
			t.Errorf("accumalator counter is wrong, got %d instead of %d", acc.Counter, expectedCounter)
		}
	})

	t.Run("halts execution when error returned", func(t *testing.T) {
		acc := &Accumulator{}
		err := Do(
			acc.Dumb,
			acc.Inc,
			acc.Inc,
			acc.Inc,
		)
		if err.Error() != "I don't know how to increment" {
			t.Errorf("error <%s> not expected", err)
		}
		expectedCounter := 0
		if acc.Counter != expectedCounter {
			t.Errorf("accumalator counter is wrong, got %d instead of %d", acc.Counter, expectedCounter)
		}
	})
}
