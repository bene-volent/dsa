package random // Package for random number generation functions

import (
	"math/rand" // Import the math/rand package for random number generation
)

// Seed sets the seed for the random number generator
func Seed(seed int64) {
	// Set the seed using the provided value
	rand.Seed(seed)
}

// RandInt generates a random integer between a (inclusive) and b (exclusive)
func RandInt(a, b int) int {
	// Ensure a is less than or equal to b
	if a > b {
		a, b = b, a // Swap values if a is greater than b
	}
	// Handle the case where a and b are equal
	if a == b {
		return a // Return the common value
	}
	// Generate a random number between 0 and (b-a) (inclusive)
	return rand.Intn(b-a+1) + a
}

// RandFloat32 generates a random float32 between a (inclusive) and b (exclusive)
func RandFloat32(a, b float32) float32 {
	// Ensure a is less than or equal to b
	if a > b {
		a, b = b, a // Swap values if a is greater than b
	}
	// Handle the case where a and b are equal
	if a == b {
		return a // Return the common value
	}
	// Generate a random float32 between 0 and 1
	randomValue := rand.Float32()
	// Scale and offset the random value to fit the desired range
	return randomValue*(b-a) + a
}

// RandFloat64 generates a random float64 between a (inclusive) and b (exclusive)
func RandFloat64(a, b float64) float64 {
	// Ensure a is less than or equal to b
	if a > b {
		a, b = b, a // Swap values if a is greater than b
	}
	// Handle the case where a and b are equal
	if a == b {
		return a // Return the common value
	}
	// Generate a random float64 between 0 and 1
	randomValue := rand.Float64()
	// Scale and offset the random value to fit the desired range
	return randomValue*(b-a) + a
}

// Shuffle shuffles the elements of a slice based on the provided swap function
func Shuffle(length int, swap func(i, j int)) {
	// Use the Rand.Shuffle function from the math/rand package to shuffle
	rand.Shuffle(length, swap)
}
