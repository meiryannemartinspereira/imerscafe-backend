package domain

type Round struct {
	ID       string
	Customer Customer
	Recipe   Recipe
	Score    int
}
