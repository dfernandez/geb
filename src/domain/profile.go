package domain

type Profile struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Locale  string `json:"locale"`
	Profile string `json:"link"`
	Picture string `json:"picture"`
}
