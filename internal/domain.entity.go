package internal

// Entity is the entity's father interface.
type Entity interface {
	// The unique ID of the entity.
	Identity() string
	// Returns the requested Worer
	Worker() Worker
	Marshal() ([]byte, error)
	// Add a publishing event.
	AddPubEvent(DomainEvent)
	// Get all publishing events..
	GetPubEvents() []DomainEvent
	// Delete all publishing events.
	RemoveAllPubEvent()
	// Add a subscription event.
	AddSubEvent(DomainEvent)
	// Get all subscription events..
	GetSubEvents() []DomainEvent
	// Delete all subscription events.
	RemoveAllSubEvent()
}
