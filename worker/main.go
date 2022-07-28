package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Inspired by:
// https://callistaenterprise.se/blogg/teknik/2019/10/05/go-worker-cancellation/
// https://www.opsdash.com/blog/job-queues-in-go.html

const MAX_WORKERS = 10

var jobList sync.Map

type Job struct {
	Id     int
	Status int // 0 - created, 1 - inFlight, 2 - completed, 9 - failed
}

type Consumer struct {
	jobsChan  chan Job
	proxyChan chan Job
}

func readNewJobs(limit int) []Job {
	count := limit

	time.Sleep(700 * time.Millisecond)

	jobs := make([]Job, 0, count)
	for i := 0; i < 200; i++ {
		item, _ := jobList.Load(i)
		job := item.(Job)
		if job.Status != 0 {
			// job is in-progress
			continue
		}

		jobs = append(jobs, job)
		if len(jobs) == count {
			break
		}
	}

	// set jobList statuses to InFlight
	readIds := ""
	for _, job := range jobs {
		job.Status = 1
		jobList.Store(job.Id, job)
		readIds += fmt.Sprintf(" %d ", job.Id)
	}

	fmt.Printf("Read %d jobs (%s)\n", count, readIds)

	return jobs
}

func (c Consumer) start(ctx context.Context) {
	for {
		select {
		case job := <-c.proxyChan:
			select {
			case <-ctx.Done():
				return
			default:
			}
			c.jobsChan <- job
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.jobsChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}

func (c Consumer) workerFunc(wg *sync.WaitGroup, index int) {
	defer wg.Done()

	for job := range c.jobsChan {
		fmt.Printf("Worker %d started job id: %d, status: %d\n", index, job.Id, job.Status)
		process(job)
		fmt.Printf("Worker %d completed job: id %d, status: %d\n", index, job.Id, job.Status)
	}

	fmt.Printf("Worker %d interrupted\n", index)
}

func process(job Job) {
	time.Sleep(time.Duration(rand.Intn(9)+1) * time.Second)
	// ...
	// if fails => set "Failed" status
	// otherwise set status to completed
	job.Status = 2
	jobList.Store(job.Id, job)
}

func waitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()
	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 200; i++ {
		jobList.Store(i, Job{Id: i, Status: 0})
	}

	consumer := Consumer{
		jobsChan:  make(chan Job, MAX_WORKERS),
		proxyChan: make(chan Job, 1),
	}

	// produce jobs
	go func() {
		for {
			jobs := readNewJobs(MAX_WORKERS) // status = New limit MAX_WORKERS
			for _, job := range jobs {
				consumer.proxyChan <- job // send one job at a time
			}
		}
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	go consumer.start(ctx)

	wg.Add(MAX_WORKERS)
	for i := 0; i < MAX_WORKERS; i++ {
		go consumer.workerFunc(wg, i)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	fmt.Println("*****\n Shutdown signal received\n*****")
	cancelFunc()
	waitWithTimeout(wg, 10*time.Second)
	fmt.Println("All workers done, shutting down!")
}
