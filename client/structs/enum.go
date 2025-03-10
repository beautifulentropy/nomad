package structs

// AllocUpdatePriority indicates the urgency of an allocation update so that the
// client can decide whether to wait longer
type AllocUpdatePriority int

const (
	AllocUpdatePriorityNone AllocUpdatePriority = iota
	AllocUpdatePriorityTypical
	AllocUpdatePriorityUrgent
)
