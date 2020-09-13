package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	defaultParallelJobs = 5
	defaultTotalJobs    = 0
)

var (
	parallelJobs int
	limitJobs    int
	countJobs    int
)

func main() {
	fmt.Println("This is a golang docker sample!")

	// concurrentJobs is a primitive example how parallel jobs work in go
	concurrentJobs(getIntEnv("PARALLEL_JOBS", defaultParallelJobs), getIntEnv("LIMIT_JOBS", defaultTotalJobs))
}

// randomNumber returns a random non-negative number
func randomNumber(maxInt int) int {
	source := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(source)

	return gen.Intn(maxInt)
}

// getIntEnv will lookup environment variables by key
// and try to convert it to int type
func getIntEnv(key string, defaultInt int) int {
	if val, ok := os.LookupEnv(key); ok {
		ret, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("Unable to receive int env variable %s", key)
			os.Exit(2)
		}
		return ret
	}
	fmt.Printf("Environment variable %s is unset using default value (%d). \n", key, defaultInt)
	return defaultInt
}

// concurrentJobs creates a buffered channel used as job queue,
// runs concurrent jobs, exit if job limit is set and reached
func concurrentJobs(jobsQueueSize int, limitJobs int) {
	jobQueue := make(chan struct{}, jobsQueueSize)
	for {
		if limitJobs != 0 && countJobs >= limitJobs {
			for len(jobQueue) > 0 {
				fmt.Printf("Stop processing new jobs limit %v reached, wait until jobqueue is empty (jobqueue: %v) \n", limitJobs, len(jobQueue))
				time.Sleep(5 * time.Second)
			}
			fmt.Println("All jobs processed! Good bye!")
			os.Exit(0)
		}
		jobQueue <- struct{}{}
		countJobs++
		go func() {
			defer func() { _ = <-jobQueue }()
			randSleepSeconds := randomNumber(10)
			time.Sleep(time.Duration(randomNumber(10)) * time.Second)
			fmt.Printf("This is a concurrency job that was running for %v seconds! \n", randSleepSeconds)
		}()
	}
}
