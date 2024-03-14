package api

type APIResponseEventUpcoming struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Desc     struct {
		Short string `json:"short"`
		Long  string `json:"long"`
	} `json:"desc"`
	Address struct {
		Line1    string `json:"line_1"`
		Line2    string `json:"line_2"`
		Street   string `json:"street"`
		City     string `json:"city"`
		Postcode string `json:"postcode"`
		Country  string `json:"country"`
	} `json:"address"`
	API struct {
		Base         string `json:"base"`
		Tickets      string `json:"tickets"`
		Participants string `json:"participants"`
		Timetables   string `json:"timetables"`
		Tournaments  string `json:"tournaments"`
	} `json:"api"`
	URL struct {
		Base         string `json:"base"`
		Tickets      string `json:"tickets"`
		Participants string `json:"participants"`
		Timetables   string `json:"timetables"`
		Tournaments  string `json:"tournaments"`
	} `json:"url"`
}

type APIResponseEventParticipants []struct {
	EventParticipant
}

type EventParticipant struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Seat     string `json:"seat"`
}
