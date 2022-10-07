package app

import "errors"

var (
	ErrItemNotFound = errors.New("item not found")
	ErrInvalidItem  = errors.New("item not valid")
)

type attendeeService struct {
	attendeeRepo AttendeeRepository
}
type eventService struct {
	eventRepo EventRepository
}
type ticketService struct {
	ticketRepo TicketRepository
}

func NewAttendeeService(attendeeRepo AttendeeRepository) AttendeeService {
	return &attendeeService{
		attendeeRepo,
	}
}
func NewEventService(eventRepo EventRepository) EventService {
	return &eventService{
		eventRepo,
	}
}

func NewTicketService(ticketRepo TicketRepository) TicketService {
	return &ticketService{
		ticketRepo,
	}
}

// Attendee Repository
func (s attendeeService) Create(attendee *Attendee) (*Attendee, error) {
	return s.attendeeRepo.Create(attendee)
}
func (s attendeeService) Read(id string) (*Attendee, error) {
	return s.attendeeRepo.Read(id)
}
func (s attendeeService) ReadAll() ([]*Attendee, error) {
	return s.attendeeRepo.ReadAll()
}
func (s attendeeService) Update(attendee *Attendee) (*Attendee, error) {
	return s.attendeeRepo.Update(attendee)
}
func (s attendeeService) Delete(id string) error {
	return s.attendeeRepo.Delete(id)
}

// Event Repository
func (s eventService) Create(event *Event) (*Event, error) {
	return s.eventRepo.Create(event)
}
func (s eventService) Read(id string) (*Event, error) {
	return s.eventRepo.Read(id)
}
func (s eventService) ReadAll() ([]*Event, error) {
	return s.eventRepo.ReadAll()
}
func (s eventService) Update(event *Event) (*Event, error) {
	return s.eventRepo.Update(event)
}
func (s eventService) Delete(id string) error {
	return s.eventRepo.Delete(id)
}

// Ticket Repository
func (s ticketService) Create(ticket *Ticket) (*Ticket, error) {
	return s.ticketRepo.Create(ticket)
}
func (s ticketService) Read(id string) (*Ticket, error) {
	return s.ticketRepo.Read(id)
}
func (s ticketService) ReadAll() ([]*Ticket, error) {
	return s.ticketRepo.ReadAll()
}
func (s ticketService) Update(ticket *Ticket) (*Ticket, error) {
	return s.ticketRepo.Update(ticket)
}
func (s ticketService) Delete(id string) error {
	return s.ticketRepo.Delete(id)
}
