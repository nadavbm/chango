package fanio

// fan-out fan-in
import (
	"fmt"
	"sync"
	"time"

	"github.com/nadavbm/chango/decorator"
	"github.com/nadavbm/chango/singleton"
	"github.com/nadavbm/chango/strategy"
)

const questions = 30

func MathClass(logger decorator.Logger, config *singleton.Config) {
	jobs := make(chan int, questions)
	results := make(chan int, questions)

	// Fan-Out: Start workers
	var wg sync.WaitGroup
	for w := 1; w <= config.Workers; w++ {
		wg.Add(1)
		go student(logger, w, jobs, results, &wg)
	}

	// Fan-In: Collect results
	go teacher(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Answer: ", res)
	}
}

func teacher(ch chan int) {
	for i := 1; i <= questions; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func student(logger decorator.Logger, id int, ch <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range ch {
		numberGenerator := strategy.Int{}
		num := numberGenerator.Integer(30, 1)
		logger.Info(fmt.Sprintf("student %d answer question %d * %d\n", id, job, num))
		results <- job * num
	}
}
