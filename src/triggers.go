package src

type JobTrigger struct {
	JobName string
}

type SchedulerTrigger struct {
	JobName         string
	IntervalType    TimeInterval
	IntervalPattern string
}
