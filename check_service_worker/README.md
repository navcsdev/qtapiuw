# Check the service is up or stop with worker pools
- The function given a list of urls and number workers to check if the service is up and running.

# Example

```
func main() {
	urls := []string{"http://google.com/", "https://facebook.com", "https://abc.com", "https://test.com", "https://test2.com"}
	wokerPools := 3
	checkHealthUrlsWorkerPools(urls, wokerPools)
}
```

Output 
```
worker 1 started  job https://abc.com
worker 2 started  job https://facebook.com
worker 0 started  job http://google.com/
worker 0 started  job https://test.com
Get http://google.com/ 200 OK
worker 2 started  job https://test2.com
Get https://facebook.com 200 OK
Get https://test2.com: dial tcp: lookup test2.com: no such host
Get https://test.com 200 OK
Get https://abc.com: dial tcp 34.216.127.34:443: i/o timeout
```
