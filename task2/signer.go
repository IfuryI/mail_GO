package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)


func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})

	for _, currJob := range jobs {
		out := make(chan interface{})

		wg.Add(1)
		go func(wg *sync.WaitGroup, currJob job, in, out chan interface{}) {
			currJob(in, out)

			close(out)
			wg.Done()
		}(wg, currJob, in, out)

		in = out
	}

	wg.Wait()
}


func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for val := range in {
		data := strconv.Itoa(val.(int))
		dataMd5 := DataSignerMd5(data)

		wg.Add(1)
		go func(data, dataMd25 string) {
			numberResults := 2
			results := make(chan orderedData, numberResults)

			lwg := &sync.WaitGroup{}
			lwg.Add(2)
			go handleCrc32(lwg, data, results, 0)
			go handleCrc32(lwg, dataMd5, results, 1)
			lwg.Wait()
			close(results)

			orderResults := getOrderValue(results)
			out <- orderResults[0] + "~" + orderResults[1]
			wg.Done()
		}(data, dataMd5)
	}
	wg.Wait()
}


func MultiHash(in, out chan interface{}) {
	hashCount := 6
	wg := &sync.WaitGroup{}

	for val := range in {
		lwg := &sync.WaitGroup{}
		results := make(chan orderedData, hashCount)

		for i := 0; i < hashCount; i++ {
			data := strconv.Itoa(i) + val.(string)
			lwg.Add(1)
			go handleCrc32(lwg, data, results, i)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			lwg.Wait()
			close(results)

			out <- strings.Join(getOrderValue(results), "")
		}()
	}
	wg.Wait()
}


func CombineResults(in, out chan interface{}) {
	var result []string
	for val := range in {
		result = append(result, val.(string))
	}
	sort.Strings(result)
	out <- strings.Join(result, "_")
}


type orderedData struct {
	number int
	value  string
}


func handleCrc32(wg *sync.WaitGroup, data string, results chan orderedData, order int) {
	results <- orderedData{order, DataSignerCrc32(data)}
	wg.Done()
}


func getOrderValue(out chan orderedData) []string {
	results := make([]orderedData, 0)
	for val := range out {
		results = append(results, val)
	}

	resultsStr := make([]string, len(results))
	for _, val := range results {
		resultsStr[val.number] = val.value
	}

	return resultsStr
}
