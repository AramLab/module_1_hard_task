package fan_in

import "sync"

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	// TODO: ваш код
	var wg sync.WaitGroup

	outCh := make(chan int)

	processChan := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			outCh <- v
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go processChan(ch)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}
