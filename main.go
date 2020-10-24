package main

import (
	"fmt"
)

func main() {
	var lap int = 4
	var engine int = 254
	var wheels int = 93
	var steeringWheel int = 49
	var wind int = 21
	var rain int = 17

	var speed int = engine + wheels + steeringWheel - wind - rain

	fmt.Println("===================")
	fmt.Print("Супер гонки. Круг ", lap, "\n")
	fmt.Println("===================")
	fmt.Print("Шумахер (", speed, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Print("Скорость: ", speed, "\n")
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колеса: +", wheels, "\n")
	fmt.Print("Руль: +", steeringWheel, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: -", wind, "\n")
	fmt.Print("Дождь: -", rain, "\n")
}
