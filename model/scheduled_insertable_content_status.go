package model


// ScheduledInsertableContentStatus : ScheduledInsertableContentStatus model
type ScheduledInsertableContentStatus string

// List of possible ScheduledInsertableContentStatus values
const (
    ScheduledInsertableContentStatus_CREATED ScheduledInsertableContentStatus = "CREATED"
    ScheduledInsertableContentStatus_SCHEDULED ScheduledInsertableContentStatus = "SCHEDULED"
    ScheduledInsertableContentStatus_TO_BE_DESCHEDULED ScheduledInsertableContentStatus = "TO_BE_DESCHEDULED"
    ScheduledInsertableContentStatus_DESCHEDULED ScheduledInsertableContentStatus = "DESCHEDULED"
    ScheduledInsertableContentStatus_ERROR ScheduledInsertableContentStatus = "ERROR"
)


