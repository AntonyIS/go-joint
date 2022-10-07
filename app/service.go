package app

type AttendeeService interface {
	Create(attendee *Attendee) (*Attendee, error)
	Read(id string) (*Attendee, error)
	ReadAll() ([]*Attendee, error)
	Update(attendee *Attendee) (*Attendee, error)
	Delete(id string) error
}

type EventService interface {
	Create(Event *Event) (*Event, error)
	Read(id string) (*Event, error)
	ReadAll() ([]*Event, error)
	Update(Event *Event) (*Event, error)
	Delete(id string) error
}

type TicketService interface {
	Create(Ticket *Ticket) (*Ticket, error)
	Read(id string) (*Ticket, error)
	ReadAll() ([]*Ticket, error)
	Update(Ticket *Ticket) (*Ticket, error)
	Delete(id string) error
}
