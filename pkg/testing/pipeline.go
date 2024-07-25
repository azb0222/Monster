package testing

// change package to package testing
import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func generator[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	//as long as this goroutine runs (until message sent on done channel), will continue to do stream<-fn() even after generator() is done
	go func() {
		defer close(stream) //when goroutine ends, close stream
		for {
			//select lets a goroutine wait on multiple (channel communication or not) operations
			// ex. for { select { default: fmt.Println("will block forever") } }
			//select blocks until one if its cases run
			//chooses one at random if multiple messages are ready
			select {
			case <-done:
				return
				//send operation on unbuffered channel will block sending go routine, until value is received
			case stream <- fn(): //will send fn() to stream channel if ready to receive a value, otherwise will block until stream channel becomes ready
				//if default case was present, it would execute immediately if none of the other cases are ready to proceed
			}
		}
	}()
	return stream
}

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(random int) bool {
		for i := random - 1; i > 1; i-- {
			if random%1 == 0 {
				return false
			}
		}
		return true
	}
	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()

	return primes
}

// allows us to control amount of data we take in from the stream produced by the generator
// numItems = number of items to take in from stream produced by generator
func done[T any, K any](done <-chan K, stream <-chan T, numItems int) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for i := 0; i < numItems; i++ {
			select {
			case <-done:
				return
			case out <- <-stream: //put value from stream, put onto out
			}
		}
	}()
	return out
}

func main() {
	doneChan := make(chan int) //whenever we close this channel, this will be relayed to any goroutine with case done
	defer close(doneChan)

	randomNumberFetcher := func() int {
		return rand.Intn(234232)
	}

	//for num := range generator(done, randomNumberFetcher) {
	//	fmt.Println(num)
	//}

	stream := generator(doneChan, randomNumberFetcher)

	//primeStream := primeFinder(doneChan, stream)
	CPUCount := runtime.NumCPU() //checks available CPUs for system

	//fan out
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(doneChan, stream)
	}

	// synchronous channel communication from generator() to done()
	//for num := range done(doneChan, primeStream, 30) {
	//	fmt.Println(num)
	//}

	// fan in
	fannedInStream := fanIn(doneChan, primeFinderChannels...)

	for num := range done(doneChan, fannedInStream, 30) {
		fmt.Println(num)
	}

	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	//
	//numsChan := source(nums)
	//
	////fan out
	//squaredNumsChan := stage2(numsChan)
	//squaredNumsChan2 := stage2(numsChan)
	//
	////fan in
	////sink(squaredNumsChan)
	//for n := range merge(squaredNumsChan, squaredNumsChan2) {
	//	fmt.Println(n)
	//}

}

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	fannedInStream := make(chan T)

	//move data from each channel to fannedInStream
	transferData := func(c <-chan T) {
		defer wg.Done()
		for nums := range c {
			select {
			case <-done:
				return
			case fannedInStream <- nums:
			}
		}
	}

	for _, channel := range channels {
		wg.Add(1)
		go transferData(channel) //spins up a different goroutine for each Channel
	}

	//waits for all transferData()s to finish, then closes the fannedInStream
	go func() {
		wg.Wait()
		close(fannedInStream)
	}()
	return fannedInStream
}

// converts list of inbound channels to one outbound channel
// 1) starts goroutine for each inbound channel to copy value to outbound channel
// 2) once all outbound goroutines are started, merge starts one more goroutine to close outbound channel after all the sends are completed
//func merge(cs ...<-chan int) <-chan int {
//	var wg sync.WaitGroup
//}

// unbuffered channels: sender and receiver goroutines must be ready at same time, as it is synchronous and will block

// returns receive-only/read-only unbuffered chan
func source(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

// takes in receive-only chan, returns receive-only chan w/ all squared values
func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out) // close out after all values have been squared and sent downstream
	}()
	return out
}

func sink(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func wgExample() {
	fmt.Println("Hello, 世界")

	var wg sync.WaitGroup //WaitGroups used to wait for all launched goroutines to finish
	//note that if wait groups are passed into functions, they can only be done by pointer

	worker := func() {
		fmt.Println("running test")
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() //wrap worker() in closure that lets wg know when its done executing
			worker()
		}()

	}

	wg.Wait() //block until wg goes back to 0; ie. all worker() instances have finished executing

	time.Sleep(10)

}
