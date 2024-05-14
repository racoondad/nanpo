package event

type ChangePassword struct {
	prototypes  map[string]interface{}
	Code        string `json:"identity"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}

// GetPrototypes implements internal.DomainEvent.
func (c *ChangePassword) GetPrototypes() map[string]interface{} {
	panic("unimplemented")
}

// Identity implements internal.DomainEvent.
func (c *ChangePassword) Identity() string {
	panic("unimplemented")
}

// Marshal implements internal.DomainEvent.
func (c *ChangePassword) Marshal() ([]byte, error) {
	panic("unimplemented")
}

// SetIdentity implements internal.DomainEvent.
func (c *ChangePassword) SetIdentity(identity string) {
	panic("unimplemented")
}

// SetPrototypes implements internal.DomainEvent.
func (c *ChangePassword) SetPrototypes(map[string]interface{}) {
	panic("unimplemented")
}

// Topic implements internal.DomainEvent.
func (c *ChangePassword) Topic() string {
	panic("unimplemented")
}

// Unmarshal implements internal.DomainEvent.
func (c *ChangePassword) Unmarshal([]byte) error {
	panic("unimplemented")
}

// event ChangePassword nanpo.DomainEvent
