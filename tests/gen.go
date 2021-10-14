package tests

import "generator/generator"

// TestGen is test struct to generate some values non-concurrent safe
type TestGen struct {
	i int
}

// Value is a generation function
func (tg *TestGen) Value() interface{} {
	tg.i++
	return tg.i
}

func TestGenerator() *generator.Generator {
	tg := TestGen{}
	return generator.New(&tg)
}
