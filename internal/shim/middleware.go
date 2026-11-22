package shim

type Middleware func(name string, args []string) error

var middlewares = []Middleware{checkCountry}

func runMiddlewares(name string, args []string) error {
	for _, m := range middlewares {
		if err := m(name, args); err != nil {
			return err
		}
	}
	return nil
}
