# semaphore
A very simple semaphore using channels.

## Installation
``` bash
go get -u -d github.com/codingconcepts/semaphore
```

## Usage
The following example creates a semaphore capable of running 10 concurrent operations.  A loop spawns 1,000,000 goroutines, that are executed with the semaphore with 10 running at any one time.  At the end of the function, the semaphore waits for all outstanding operations to complete.

``` go
sum := int64(0)

sem := New(10)
for i := 0; i < 1000000; i++ {
	sem.Run(func() {
		atomic.AddInt64(&sum, 1)
	})
}
sem.Wait()

fmt.Println(sum)
// Output: 1000000
```