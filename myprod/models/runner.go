package models

type Runner struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age,omitemty"`
	IsActive     bool      `json:"is_active"`
	Country      string    `json:"country"`
	PersonalBest string    `json:"personal_best,omitempty"`
	SeasonBest   string    `json:"season_best,omitempty"`
	Results      []*Result `json:"results,omitempty"`
}

type IdPost struct {
	IdN string `json:"idn"`
}
