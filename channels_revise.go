package main

import (
	"fmt"
)

func main() {
	jobList := []string{"one", "two", "three"}
	jobs := make(chan string)
	done := make(chan bool, len(jobList))

	//create jobs
	go func() {
		for _, job := range jobList {
			jobs <- job
		}
		close(jobs)
	}()

	//execute jobs
	go func() {
		for job := range jobs {
			fmt.Println(job)
			done <- true
		}
	}()

	//wait for all of them to exit
	for i := 0; i < len(jobList); i++ {
		<-done
	}

}
