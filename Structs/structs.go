package structs

type Press struct {
	Id    int
	Value int
}

func NewPress(id int, value int) Press {
	return Press{
		Id:    id,
		Value: value,
	}
}

type Hum struct {
	id    int
	value string
}

func NewHum(id int, value string) Hum {
	return Hum{
		id:    id,
		value: value,
	}
}

type Seis struct {
	id    int
	value float64
}

func NewSeis(id int, value float64) Seis {
	return Seis{
		id:    id,
		value: value,
	}
}
