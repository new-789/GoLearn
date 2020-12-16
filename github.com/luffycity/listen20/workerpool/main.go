package main

import (
	"fmt"
	"math/rand"
)

type Job struct{
	Number int
	Id int
}

type Result struct{
	Job *Job
	Sum int
}

func calc(job *Job, result chan *Result) {
	var sum int
	//number := job.Number
	for job.Number != 0 {
		tmp := job.Number % 10  // 取 number 的个位数
		sum += tmp
		job.Number /= 10  // 每循环一次取 number 十位上的值
	}

	r := &Result{
		Job: job,
		Sum:sum,
	}

	result <- r  // 将计算的结果存入 result 管道中
}

func Worker(jobChan chan *Job, result chan *Result) {
	for job := range jobChan {
		calc(job, result)
	}
}

func startWorkerPool(num int, jobChan chan *Job, result chan *Result) {
	for i := 0; i < num; i++ {
		go Worker(jobChan, result)
	}
}

func printResult(resultChan chan *Result) {
	for result := range resultChan{
		fmt.Printf("job id:%v, number:%v, rsult:%v\n", result.Job.Id, result.Job.Number, result.Sum)
	}
}

func main() {
	jobChan := make(chan *Job, 1000)
	resultChan := make(chan *Result, 1000)

	startWorkerPool(128, jobChan, resultChan)

	go printResult(resultChan)
	var id int
	for {
		id++
		number := rand.Int()
		job := &Job{
			Id:     id,
			Number: number,
		}

		jobChan <- job  // 将需要计算的内容放入 jobChan 管道中
	}
}