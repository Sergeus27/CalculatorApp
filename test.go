// TODO написать прогу которая 10000 сложит строку со строкой,
//
//	строку в стринг билдере и проверить что будет быстрее по времени исполнения
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Андрюша поридж настоящий не знает что название файла должно быть main")
	fmt.Println("")
	aToi := "aToi_"
	iterations := 5000

	for gi := iterations; gi < 100000; gi += 5000 {
		poridge := ""
		start := time.Now()
		for i := 0; i < gi; i++ {
			poridge += aToi + strconv.Itoa(i) + " "
		}
		elapsed := time.Since(start)
		fmt.Printf("%5d итераций | += : %.2f ms\n", gi, float64(elapsed)/float64(time.Millisecond))
	}

	fmt.Println("")

	for gi := iterations; gi < 100000; gi += 5000 {
		// 🔹 Тест 2: strings.Builder
		start := time.Now()
		var builder strings.Builder
		// Опционально: предварительный резерв памяти (ускоряет)
		// builder.Grow(iterations * 12)
		for i := 0; i < gi; i++ {
			builder.WriteString(aToi)
			builder.WriteString(strconv.Itoa(i))
			builder.WriteByte(' ')
		}
		elapsed := time.Since(start)
		fmt.Printf("%5d итераций | Builder: %.2f ms\n", gi, float64(elapsed)/float64(time.Millisecond))
	}
}
