package performer

type Task func() error

func Do(iterator ...Task) error {
	for _, fn := range iterator {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
