package eventsRepository

type Repository interface {
	CreateEvent() error                // recibe dto
	GetAllEvents()                     //  reciba objeto con los filtros
	GetEventById(eventId string) error // dto única devolución
	UpdateEvent(eventId string) error
	DeleteEvent(eventId string) error
	GetPublishedEvents() //dto []arreglo
	GetCompletedEvents() //dto []
}
