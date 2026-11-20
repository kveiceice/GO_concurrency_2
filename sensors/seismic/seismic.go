package seismic

import (
	structs "conc_2/Structs"
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func Seismic(ctx context.Context, wg *sync.WaitGroup, seismicTransfer chan<- structs.Seis, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я датчик сейсмической активности номер", id, "я завершаю свою работу")
			return
		default:
			fmt.Println("Я датчик сейсмической активности номер", id, "я узнаю давление влажность")
			time.Sleep(1 * time.Second)
			fmt.Println("Я датчик сейсмической активности номер", id, "я изучил какая сейчас  сейсмическая активность ")
			act := rand.Float64() * 10
			x := structs.NewSeis(id, act)
			seismicTransfer <- x
			fmt.Println("Я датчик сейсмической активности номер", id, "и я передал показания")
		}
	}
}

func SeismicPull(ctx context.Context, number int) <-chan structs.Seis {
	SeismucTransfer := make(chan structs.Seis)
	wg := &sync.WaitGroup{}
	for i := 1; i <= number; i++ {
		wg.Add(1)
		go Seismic(ctx, wg, SeismucTransfer, i)
	}

	go func() {
		wg.Wait()
		close(SeismucTransfer)
	}()

	return SeismucTransfer
}
