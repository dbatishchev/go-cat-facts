package service

import (
	"testing"
	"errors"
	"fmt"
)

func TestGetLocalRandomFact(t *testing.T) {
	fact := getLocalRandomFact()

	if len(fact) == 0 {
		errors.New("Failed to get random cat fact")
	}
}

func TestGetFact(t *testing.T) {
	fact, err := GetFact()

	if err != nil {
		t.Errorf("Failed to get cat fact: %v", err)
	}

	if len(fact) == 0 {
		t.Error("Failed to get cat fact")
	}
}