package grains

import (
	"errors"
)

const squareCount uint = 64

func Total() uint64 {
	return 1<<squareCount - 1
}

func Square(index int) (uint64, error) {
	if index >= 1 && index <= int(squareCount) {
		return 1 << uint(index-1), nil
	}
	return 0, errors.New("Out of Bounce")
}
