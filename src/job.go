package src

import "time"

type Job struct {
	ID                  int16
	Name                string
	PID                 int
	Description         string
	Origin              JobOrigin
	Status              JobStatus
	Scheduler           *Scheduler
	LastStart           time.Time
	LastEnd             time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	JobDependencies     []Job
	OnSetupCallbacks    []func()
	OnFinishCallbacks   []func(job Job)
	FinishChannelReport chan<- Job
	StartChannelReport  chan<- Job
}
