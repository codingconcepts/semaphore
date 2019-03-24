package semaphore

import (
	"fmt"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestRunWait(t *testing.T) {
	cases := []struct {
		name        string
		concurrency int
		expSum      int
	}{
		{
			name:        "concurrency 1 count 0",
			concurrency: 1,
			expSum:      0,
		},
		{
			name:        "concurrency 10 count 0",
			concurrency: 10,
			expSum:      0,
		},
		{
			name:        "concurrency 100 count 0",
			concurrency: 100,
			expSum:      0,
		},
		{
			name:        "concurrency 1000 count 0",
			concurrency: 1000,
			expSum:      0,
		},
		{
			name:        "concurrency 1 count 1",
			concurrency: 1,
			expSum:      1,
		},
		{
			name:        "concurrency 10 count 1",
			concurrency: 10,
			expSum:      1,
		},
		{
			name:        "concurrency 100 count 1",
			concurrency: 100,
			expSum:      1,
		},
		{
			name:        "concurrency 1000 count 1",
			concurrency: 1000,
			expSum:      1,
		},
		{
			name:        "concurrency 1 count 1000000",
			concurrency: 1,
			expSum:      1000000,
		},
		{
			name:        "concurrency 10 count 1000000",
			concurrency: 10,
			expSum:      1000000,
		},
		{
			name:        "concurrency 100 count 1000000",
			concurrency: 100,
			expSum:      1000000,
		},
		{
			name:        "concurrency 1000 count 1000000",
			concurrency: 1000,
			expSum:      1000000,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sem := New(c.concurrency)
			sum := int64(0)
			for i := 0; i < c.expSum; i++ {
				atomic.AddInt64(&sum, 1)
			}
			sem.Wait()

			equals(t, int64(c.expSum), sum)
		})
	}
}

func BenchmarkRunWait(b *testing.B) {
	cases := []struct {
		name        string
		concurrency int
	}{
		{
			name:        "concurrency 1",
			concurrency: 1,
		},
		{
			name:        "concurrency 10",
			concurrency: 10,
		},
		{
			name:        "concurrency 100",
			concurrency: 100,
		},
		{
			name:        "concurrency 1000",
			concurrency: 1000,
		},
		{
			name:        "concurrency 10000",
			concurrency: 10000,
		},
		{
			name:        "concurrency 100000",
			concurrency: 100000,
		},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			sem := New(c.concurrency)
			for i := 0; i < b.N; i++ {
				sem.Run(func() {})
			}
			sem.Wait()
		})
	}
}

func Example() {
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
}

func equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\n\texp: %#[1]v (%[1]T)\n\tgot: %#[2]v (%[2]T)\n", exp, act)
	}
}
