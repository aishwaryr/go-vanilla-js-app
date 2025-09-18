package models

type Actor struct {
	ID        int
	FirstName string
	LastName  string
	ImageUrl  *string // nullable
}
