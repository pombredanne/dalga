package dalga

import (
	"bytes"
	"time"
)

type Job struct {
	RoutingKey string
	Body       []byte
	Interval   time.Duration
	NextRun    time.Time
}

func NewJob(routingKey string, body []byte, interval uint32) *Job {
	job := Job{
		RoutingKey: routingKey,
		Body:       body,
		Interval:   time.Duration(interval) * time.Second,
	}
	job.SetNewNextRun()
	return &job
}

// Remaining returns the duration until the job's next scheduled time.
func (j *Job) Remaining() time.Duration {
	return -time.Since(j.NextRun)
}

// SetNewNextRun calculates the new run time according to current time and sets it on the job.
func (j *Job) SetNewNextRun() {
	j.NextRun = time.Now().UTC().Add(j.Interval)
}

func (j *Job) Equals(k *Job) bool {
	return (j.RoutingKey == k.RoutingKey) &&
		(bytes.Compare(j.Body, k.Body) == 0)
}
