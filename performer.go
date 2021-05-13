// Performer intends to be an alternative to the multiple conditionals spread through the code, providing a cleaner way to handle errors.
package performer

// Task is the contract that should be implemented by your function(s) (or methods).
type Task func() error

// Do is the function that will perform the tasks. Given one or multiple tasks, they will be performed one by one. When faced with an error, the execution is halted and error is returned.
func Do(iterator ...Task) error {
	for _, fn := range iterator {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
