package hellogomod

import "rsc.io/quote/v3"

// Hello returns Hello, world.
func Hello() string {
	return quote.HelloV3()
}

// Proverb returns proverb
func Proverb() string {
	return quote.Concurrency()
}
