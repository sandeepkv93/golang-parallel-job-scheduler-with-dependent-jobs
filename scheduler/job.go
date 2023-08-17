package scheduler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sandeepkv93/paralleljobschedulerwithdependentjobs/concurrencyutils"
)

type Job struct {
	// string name
	Name string
	// list of Jobs struct childrenJobs
	ChildrenJobs []*Job
	// list of jobs struct parentsJobs
	ParentsJobs []*Job
	latch       *concurrencyutils.CountDownLatch
}

func NewJob(name string, parentJobs ...*Job) *Job {
	job := &Job{
		Name: name,
	}

	for _, parentJob := range parentJobs {
		job.ParentsJobs = append(job.ParentsJobs, parentJob)
		parentJob.ChildrenJobs = append(parentJob.ChildrenJobs, job)
	}

	job.latch = concurrencyutils.NewCountDownLatch(len(parentJobs))
	return job
}

func (job *Job) Run() {
	fmt.Println(job.Name + " started")

	// Sleeping for a random time between 4 and 8 seconds to simulate work
	time.Sleep(time.Duration(rand.Intn(4)+4) * time.Second)

	fmt.Println(job.Name + " completed")
}
