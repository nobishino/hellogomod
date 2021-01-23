package hellogomod

import (
	quoteV3 "rsc.io/quote/v3"
)

// Hello returns Hello, world.
func Hello() string {
	return quoteV3.HelloV3()
}

// Proverb returns proverb
func Proverb() string {
	return quoteV3.Concurrency()
}
