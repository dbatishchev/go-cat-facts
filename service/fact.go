package service

import (
	"math/rand"
	"go-cat-facts/api"
)

var facts = []string{
	"When a go-cat-facts chases its prey, it keeps its head level. Dogs and humans bob their heads up and down.",
	"Unlike dogs, cats do not have a sweet tooth. Scientists believe this is due to a mutation in a key taste receptor.",
	"A go-cat-facts can’t climb head first down a tree because every claw on a go-cat-facts’s paw points the same way. To get down from a tree, a go-cat-facts must back down",
	"Cats make about 100 different sounds. Dogs make only about 10.",
	"A group of cats is called a “clowder.”",
}

// get fact from "Cat Facts" API even service don't reachable
func GetFact() (fact string, err error) {
	fact, err = api.GetFact()
	if err != nil {
		return getLocalRandomFact(), nil
	}

	return fact, nil
}

// get random fact from slice
func getLocalRandomFact() string {
	n := rand.Int() % len(facts)
	return facts[n]
}