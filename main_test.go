package main

import (
	"awesomeProject/operations"
	"awesomeProject/staticTests"
	"math/rand"
	"testing"
	"time"
)

func TestFIPS140_3(t *testing.T) {
	// Ініціалізуємо генератор випадкових чисел з використанням seed
	rand.Seed(time.Now().UnixNano())

	// Розмір бітової послідовності
	bitSequenceSize := 20000

	// Генеруємо бітову послідовність
	bitSequence := operations.GenerateRandomBitSequence(bitSequenceSize)

	if staticTests.CheckMonobit(bitSequence) != true {
		t.Errorf("Monobit test failed")
	}
	if staticTests.CheckMaxRun(bitSequence) != true {
		t.Errorf("MaxRun test failed")
	}
	if staticTests.CheckRunLengthOne(bitSequence) != true || staticTests.CheckRunLengthZero(bitSequence) != true {
		t.Errorf("RunLengthOne or RunLengthZero failed")
	}
	if staticTests.CheckPoker(bitSequence) != true {
		t.Errorf("Poker failed")
	}
}
