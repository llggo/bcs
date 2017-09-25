package math

import (
	"crypto/rand"
	"fmt"
	"sync"
)

// Source String used when generating a random identifier.
const idSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const idSourceUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const idSourceNumber = "0123456789"

// Save the length in a constant so we don't look it up each time.
const idSourceUpperLen = byte(len(idSourceUpper))
const idSourceNumberLen = byte(len(idSourceNumber))

// GenerateID creates a prefixed random identifier.
func RandString(prefix string, length int) string {
	// Create an array with the correct capacity
	arr := make([]uint8, length)
	// Fill our array with random numbers
	rand.Read(arr)
	id := make([]byte, length)

	// Replace each random number with an alphanumeric value
	for i, b := range arr {
		id[i] = idSource[b]
	}

	// Return the formatted id
	if "" == prefix {
		return fmt.Sprintf("%s", string(id))
	} else {
		return fmt.Sprintf("%s_%s", prefix, string(id))
	}
}
func RandStringUpper(prefix string, length int) string {
	// Create an array with the correct capacity
	id := make([]uint8, length)
	// Fill our array with random numbers
	rand.Read(id)

	// Replace each random number with an alphanumeric value
	for i, b := range id {
		id[i] = idSourceUpper[b%idSourceUpperLen]
	}

	// Return the formatted id
	return fmt.Sprintf("%s_%s", prefix, string(id))
}

func RandStringNumber(prefix string, length int) string {
	// Create an array with the correct capacity
	id := make([]uint8, length)
	// Fill our array with random numbers
	rand.Read(id)

	// Replace each random number with an alphanumeric value
	for i, b := range id {
		id[i] = idSourceNumber[b%idSourceNumberLen]
	}

	// Return the formatted id
	return fmt.Sprintf("%s_%s", prefix, string(id))
}

type RandStringMaker struct {
	Prefix string
	Length int
	sync.Mutex
}

func (m *RandStringMaker) Next() string {
	m.Lock()
	var N = RandString(m.Prefix, m.Length)
	m.Unlock()
	return N
}

var numbers = "0123456789"

func RandNumString(length int) string {
	// Create an array with the correct capacity
	id := make([]uint8, length)
	// Fill our array with random numbers
	rand.Read(id)

	// Replace each random number with an alphanumeric value
	for i, b := range id {
		id[i] = numbers[b%10]
	}
	return string(id)
}
