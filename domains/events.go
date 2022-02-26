package domains

import "time"

type Event struct {
	ID        int64           `bson:"_id,omitempty"`
	Type      string          `bson:"type,omitempty"`
	Payload   interface{}     `bson:"payload,omitempty"`
	State     string          `bson:"state,omitempty"`
	History   []EventsHistory `bson:"history"`
	CreatedAt time.Time       `bson:"created_at,omitempty"`
	UpdatedAt time.Time       `bson:"updated_at,omitempty"`
}

type EventsHistory struct {
	ID        int64     `bson:"id,omitempty"`
	State     string    `bson:"state,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
}

type EventModel interface {
	Create(e Event) error
	GetById(id int64) (Event, error)
	UpdateStateById(id int64, state string) error
}
