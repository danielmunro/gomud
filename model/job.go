package model

import (
	"github.com/jinzhu/gorm"
	"strings"
)

type JobName string

const (
	MageJob = "mage"
	ClericJob = "cleric"
	WarriorJob = "warrior"
	ThiefJob = "thief"
	UninitializedJob = "uninitialized"
)

type Job struct {
	gorm.Model
	Name JobName
	Attributes *Attributes
}

func NewJob(name JobName, attr *Attributes) *Job {
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
		NewJob(MageJob, NewStats(0, 2, 1, 0, 0)),
		NewJob(WarriorJob, NewStats(2, 0, 0, 0, 1)),
		NewJob(ThiefJob, NewStats(1, 0, 0, 2, 0)),
		NewJob(ClericJob, NewStats(0, 1, 2, 0, 0)),
		NewJob(UninitializedJob, &Attributes{}),
	}
}
