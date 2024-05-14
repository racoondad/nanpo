package internal

type DomainEvent interface {
	// The only noun used to identify an event.
	Topic() string
	// Set the property.
	SetPrototypes(map[string]interface{})
	// Read the property.
	GetPrototypes() map[string]interface{}
	//
	Marshal() ([]byte, error)
	//
	Unmarshal([]byte) error
	// The unique ID of the event.
	Identity() string
	// Set a unique ID.
	SetIdentity(identity string)
}
