package main

import (
	"github.com/sandeepkv93/paralleljobschedulerwithdependentjobs/scheduler"
)

func main() {
	jobA := scheduler.NewJob("Job A")
	jobB := scheduler.NewJob("Job B")
	jobC := scheduler.NewJob("Job C", jobA)
	jobD := scheduler.NewJob("Job D", jobB)
	jobE := scheduler.NewJob("Job E", jobC, jobD)
	jobF := scheduler.NewJob("Job F", jobE)
	jobG := scheduler.NewJob("Job G", jobE)
	jobH := scheduler.NewJob("Job H", jobE)
	jobI := scheduler.NewJob("Job I", jobF, jobG, jobH)

	allJobs := []*scheduler.Job{jobA, jobB, jobC, jobD, jobE, jobF, jobG, jobH, jobI}
	scheduler.ScheduleAllJobs(allJobs)
}
