package model

type DogFactResponse struct {
	Facts   []string // Field names are capitalized to be exported
	Success bool
}
