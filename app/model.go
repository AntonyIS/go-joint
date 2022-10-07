package app

type Attende struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"gte=3 & lte=25 & format=alnum_unicode"`
	Email    string `json:"email" validate:"empty=false | format=email"`
	About    string `json:"about" validate:"lte=140"`
	Ticket   Ticket `json:"ticket"`
}

type Event struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" validate:"gte=3 & lte=25 & format=alnum_unicode"`
	Location  string    `json:"location" `
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	Start     int64     `json:"start"`
	End       int64     `json:"end"`
	Attendes  []Attende `json:"attendes"`
}

type Ticket struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	Date   string
}
