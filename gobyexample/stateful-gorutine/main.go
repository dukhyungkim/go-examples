package main

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps, writeOps uint64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		state := make(map[int]int)
		select {
		case read := <-reads:
			read.resp <- state[read.key]
		case write := <-writes:
			state[write.key] = write.val
			write.resp <- true
		}
	}()

	for r := 0; r < 100; r++ {

	}
}
