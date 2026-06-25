package manager

type API struct {
	URL string
}

type Event struct {
	Name         string             `json:"name"`
	Capacity     int                `json:"capacity"`
	Start        string             `json:"start"`
	End          string             `json:"end"`
	Slug         string             `json:"slug"`
	Description  EventDescription   `json:"description"`
	Address      EventAddress       `json:"address"`
	URL          EventURL           `json:"url"`
	Participants []EventParticipant `json:"participants"`
	Tickets      []EventTicket      `json:"tickets"`
}

type EventDescription struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

type EventAddress struct {
	Line1    string `json:"line_1"`
	Line2    string `json:"line_2"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
	Country  string `json:"country"`
}

type EventURL struct {
	Base         string `json:"base"`
	Tickets      string `json:"tickets"`
	Participants string `json:"participants"`
	Timetables   string `json:"timetables"`
}

type EventParticipant struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Seat     string `json:"seat"`
}

type EventTicket struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type Timetable struct {
	Name    string           `json:"name"`
	Slug    string           `json:"slug"`
	Status  string           `json:"status"`
	Default bool             `json:"default"`
	Data    []TimetableEntry `json:"data"`
}

type TimetableEntry struct {
	StartTime   string `json:"start_time"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
