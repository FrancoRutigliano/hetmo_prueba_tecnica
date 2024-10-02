package userEventsDto

type UserEventsListDto struct {
	EventID          string `db:"event_id" json:"event_id"`                   // ID del evento
	Title            string `db:"title" json:"title"`                         // Título del evento
	ShortDescription string `db:"short_description" json:"short_description"` // Descripción corta
	Date             string `db:"date" json:"date"`                           // Fecha del evento
	Organizer        string `db:"organizer" json:"organizer"`                 // Nombre del organizador
	Location         string `db:"location" json:"location"`                   // Ubicación
	IsPublished      bool   `db:"is_published" json:"is_published"`           // Estado de publicación
}
