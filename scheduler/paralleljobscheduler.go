package scheduler

import (
	"sync"

	"github.com/ghosind/collection/set"
	"github.com/sandeepkv93/paralleljobschedulerwithdependentjobs/utils"
)

func ScheduleAllJobs(jobs []*Job) {
	// Get All starting jobs (jobs with no parents)
	var startingJobs *set.HashSet[*Job] = getStartingJobs(jobs)

	// Get All the children jobs in order
	childrenJobs := getChildrenJobsInOrder(startingJobs)

	wg := sync.WaitGroup{}
	for job, _ := range *startingJobs {
		wg.Add(1)
		go func(job *Job) {
			ProcessJob(job)
			wg.Done()
		}(job)
	}

	for _, job := range childrenJobs {
		wg.Add(1)
		go func(job *Job) {
			ProcessJob(job)
			wg.Done()
		}(job)
	}

	wg.Wait()
}

func getStartingJobs(jobs []*Job) *set.HashSet[*Job] {
	startingJobs := set.NewHashSet[*Job]()
	for _, job := range jobs {
		if len(job.ParentsJobs) == 0 {
			startingJobs.Add(job)
		}
	}
	return startingJobs
}

func getChildrenJobsInOrder(startingJobs *set.HashSet[*Job]) []*Job {
	var allChildrenJobs []*Job
	allChildrenJobsSet := set.NewHashSet[*Job]()
	var queue *utils.Queue = utils.NewQueue()
	for job, _ := range *startingJobs {
		queue.Enqueue(job)
	}

	for !queue.IsEmpty() {
		job := queue.Dequeue().(*Job)

		if !startingJobs.Contains(job) && !allChildrenJobsSet.Contains(job) {
			allChildrenJobsSet.Add(job)
			allChildrenJobs = append(allChildrenJobs, job)
		}

		for _, childJob := range job.ChildrenJobs {
			if !allChildrenJobsSet.Contains(childJob) {
				queue.Enqueue(childJob)
			}
		}
	}

	return allChildrenJobs
}

func ProcessJob(job *Job) {
	// Wait for all parent jobs to complete
	job.latch.Wait()

	// Run the job
	job.Run()

	// Notify all child jobs
	for _, childJob := range job.ChildrenJobs {
		childJob.latch.CountDown()
	}
}
