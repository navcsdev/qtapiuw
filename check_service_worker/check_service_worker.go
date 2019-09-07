package check_service_worker

import (
	"fmt"
	"net/http"
)

func Worker(id int, jobs <-chan string, resc chan<- string, errc chan<- error) {
	for url := range jobs {
		fmt.Println("worker", id, "started  job", url)
		res, err := http.Get(url)
		if err != nil {
			errc <- err
		} else {
			resc <- "Get " + url + " " + res.Status
		}
	}
}

func CheckHealthUrlsWorkerPools(urls []string, pools int) {
	jobsc, resc, errc := make(chan string, 100), make(chan string), make(chan error)

	for i := 0; i < pools; i++ {
		go Worker(i, jobsc, resc, errc)
	}

	for _, url := range urls {
		jobsc <- url
	}
	close(jobsc)

	for i := 0; i < len(urls); i++ {
		select {
		case res := <-resc:
			fmt.Println(res)
		case err := <-errc:
			fmt.Println(err)
		}
	}
}
