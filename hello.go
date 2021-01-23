package hellogomod

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

// Hello returns Hello, world.
func Hello() string {
	return quote.Hello()
}

// Proverb returns proverb
func Proverb() string {
	return quoteV3.Concurrency()
}
