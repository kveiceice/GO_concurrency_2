package humidity

import (
	structs "conc_2/Structs"
	"context"
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
)

func Humidity(ctx context.Context, wg *sync.WaitGroup, HumTransfer chan<- structs.Hum, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я датчик влажности номер", id, "я завершаю свою работу")
			return
		default:
			fmt.Println("Я датчик влажности номер", id, "я узнаю давление влажность")
		}
		select {
		case <-ctx.Done():
			fmt.Println("Я датчик влажности номер", id, "я завершаю свою работу")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Я датчик влажности номер", id, "я изучил какая сейчас влажность")
			hum := strconv.FormatFloat(10+rand.Float64()*70, 'f', 2, 64) + "%"
			x := structs.NewHum(id, hum)
			HumTransfer <- x
			fmt.Println("Я датчик влажности номер", id, "и я передал показания")
		}
	}
}

func HumidityPull(ctx context.Context, number int) <-chan structs.Hum {
	HumTransfer := make(chan structs.Hum)
	wg := &sync.WaitGroup{}
	for i := 1; i <= number; i++ {
		wg.Add(1)
		go Humidity(ctx, wg, HumTransfer, i)
	}

	go func() {
		wg.Wait()
		close(HumTransfer)
	}()

	return HumTransfer
}
