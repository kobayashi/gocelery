package gocelery

import "time"

func Example_worker() {

	// initialize celery client
	cli, _ := NewCeleryClient(
		NewRedisCeleryBroker("redis://"),
		NewRedisCeleryBackend("redis://"),
		5, // number of workers
	)

	// task
	add := func(a, b int) int {
		return a + b
	}

	// register task
	cli.Register("add", add)

	// start workers (non-blocking call)
	cli.StartWorker()

	// wait for client request
	time.Sleep(10 * time.Second)

	// stop workers gracefully (blocking call)
	cli.StopWorker()
}
