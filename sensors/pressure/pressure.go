package pressure

import (
	structs "conc_2/Structs"
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func pressure(ctx context.Context, wg *sync.WaitGroup, pressureTransfer chan<- structs.Press, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я датчик давления воздуха номер", id, "я завершаю свою работу")
			return
		default:
			fmt.Println("Я датчик давления воздуха номер", id, "я узнаю давление воздуха")
			time.Sleep(1 * time.Second)
			fmt.Println("Я датчик давления воздуха номер", id, "я изучил какое сейчас давление воздуха")
			press := 700 + rand.IntN(80)
			x := structs.NewPress(id, press)
			pressureTransfer <- x
			fmt.Println("Я датчик давления воздуха номер", id, "и я передал показания")
		}
	}
}

func PressurePull(ctx context.Context, number int) <-chan structs.Press {
	pressTrans := make(chan structs.Press)
	wg := &sync.WaitGroup{}
	for i := 1; i <= number; i++ {
		wg.Add(1)
		go pressure(ctx, wg, pressTrans, i)
	}

	go func() {
		wg.Wait()
		close(pressTrans)
	}()

	return pressTrans
}
