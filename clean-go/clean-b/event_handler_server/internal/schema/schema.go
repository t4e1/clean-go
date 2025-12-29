package schema

import (
	"time"
)

type EventSchema struct {
	event_id       string
	timestamp      time.Time
	group_id       string
	group_order_id int
	event_type     string // SEQUENCIAL, CONCURRENCE  둘중 택1
}

type EventSubSchema struct {
	event_id       string
	group_id       string
	group_order_id int
}
