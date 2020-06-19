# semaphore
A very simple semaphore using channels.

[![Go Report Card](https://goreportcard.com/badge/github.com/codingconcepts/semaphore)](https://goreportcard.com/report/github.com/codingconcepts/semaphore) [![Build Status](https://travis-ci.org/codingconcepts/semaphore.svg?branch=master)](https://travis-ci.org/codingconcepts/semaphore)

## Installation
``` bash
go get -u -d github.com/codingconcepts/semaphore
```

## Usage
The following example kicks off 1,000,000 goroutines, each of which increment a `sum` variable atomically.  After all operations have been kicked off, we ask the semaphore to wait for them to complete.

``` go
sum := int64(0)

sem := semaphore.New(10)
for i := 0; i < 1000000; i++ {
	sem.Run(func() {
		atomic.AddInt64(&sum, 1)
	})
}
sem.Wait()

fmt.Println(sum)
// Prints: 1,000,000
```

Note that if you call `Wait()` before all operations have been kicked off (and don't call it again at the end), the result of `sum` is not guaranteed.
