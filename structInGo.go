package main

import "fmt"


type car struct {
	gas_pedal uint16
	break_pedal uint16
	streering_wheel uint16
	top_speed float64
}

func main() {
	a_car := car {gas_pedal: 16535, break_pedal: 0, streering_wheel: 12562, top_speed: 225.0}
	fmt.Println("gas_pedal: ", a_car.gas_pedal, "break_pedal: ", a_car.break_pedal, "streering_wheel: ", a_car.streering_wheel, "top_speed: ", a_car.top_speed)
}