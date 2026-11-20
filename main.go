package main

import (
	structs "conc_2/Structs"
	"conc_2/sensors/humidity"
	"conc_2/sensors/pressure"
	"conc_2/sensors/seismic"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	mtxP := sync.Mutex{}
	mtxH := sync.Mutex{}
	mtxS := sync.Mutex{}
	wg := sync.WaitGroup{}
	pressureContext, pressureCancel := context.WithCancel(context.Background())
	humidityContext, HumidityCancel := context.WithCancel(context.Background())
	seismicContext, SeismicCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer SeismicCancel()

	PressTr := pressure.PressurePull(pressureContext, 150)
	HumTr := humidity.HumidityPull(humidityContext, 150)
	SeisTr := seismic.SeismicPull(seismicContext, 150)
	InitTime := time.Now()
	PressureSl := make([]structs.Press, 0, 10)
	HumiditySl := make([]structs.Hum, 0, 10)
	SeismicSl := make([]structs.Seis, 0, 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range PressTr {
			mtxP.Lock()
			PressureSl = append(PressureSl, v)
			// МОЖНО ДЕЛАТЬ ВСЕ ВЫВОДЫ ЗДЕСЬ С УЖЕ СОЗДАННОЙ СТРКУКТУРОЙ А НЕ В ДАТЧИКАХ
			// ТАКЖЕ МЬЮТЕКСЫ ЗДЕСЬ НЕ НУЖНЫ ТАК КАК WAIT НО ПРОБЛЕМА В ТОМ ЧТО ЧТЕНИЕ LEN БЕЗ МЬЮТЕКСА - ТАМ ОН КАК РАЗ НУЖЕН
			// ТАК ЖЕ ПО ТЗ НУЖНО БЫЛО СДЕЛАТЬ КООРДИНАТЫ - ЭТО ЛЕГКО ДОБАВИТЬ ПРОСТО РАНДОМ ГЕНЕРИТЬ
			// ТАКЖЕ SELECT ВЕЗДЕ НУЖЕН ЧЕРЕЗ <-ctx.DONE()
			// В САМОЙ СТРУКТУРЕ НУЖНО ID VALUE НЕ СКРЫТЫЕ ЧТОБЫ МОЖНО БЫЛО ИЗУЧАТЬ ЭТИ ДАННЫЕ ДАЛЬШЕ
			mtxP.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range HumTr {
			mtxH.Lock()
			HumiditySl = append(HumiditySl, v)
			mtxH.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range SeisTr {
			mtxS.Lock()
			SeismicSl = append(SeismicSl, v)

			mtxS.Unlock()
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("-------------PressureCancel---------------")
		pressureCancel()

	}()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("-------------HumidityCancel---------------")
		HumidityCancel()
	}()
	wg.Wait()
	FinishTime := time.Since(InitTime)
	fmt.Println()
	fmt.Println("Всего было проведено исследований датчиками: type-{id, value}")
	fmt.Println("Давление:", len(PressureSl))
	fmt.Println(PressureSl)
	fmt.Println("Влажность:", len(HumiditySl))
	fmt.Println(HumiditySl)
	fmt.Println("Сейсмическая активность:", len(SeismicSl))
	fmt.Println(SeismicSl)
	fmt.Println("Итоговое время выполнения -", FinishTime)
}
