package gomud

import "strings"

type JobName string

const (
	MageJob = "mage"
	ClericJob = "cleric"
	WarriorJob = "warrior"
	ThiefJob = "thief"
	UninitializedJob = "uninitialized"
)

type Job struct {
	Name JobName
	Attributes *attributes
}

func NewJob(name JobName, attr *attributes) *Job {
	return &Job{
		Name: name,
		Attributes: attr,
	}
}

var jobs []*Job

func getJob(jobName JobName) *Job {
	for _, i := range jobs {
		if strings.HasPrefix(string(i.Name), string(jobName)) {
			return i
		}
	}

	return jobs[len(jobs)]
}

func init() {
	jobs = []*Job{
		NewJob(MageJob, newStats(0, 2, 1, 0, 0)),
		NewJob(WarriorJob, newStats(2, 0, 0, 0, 1)),
		NewJob(ThiefJob, newStats(1, 0, 0, 2, 0)),
		NewJob(ClericJob, newStats(0, 1, 2, 0, 0)),
		NewJob(UninitializedJob, &attributes{}),
	}
}
