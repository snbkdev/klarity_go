// 4.2.1 Функция append
package main

import "fmt"

func main() {
	// var runes []rune
	// for _, r := range "Hello, Может быть" {
	// 	runes = append(runes, r)
	// }

	// fmt.Printf("%q\n", runes)
	// ['H' 'e' 'l' 'l' 'o' ',' ' ' 'М' 'о' 'ж' 'е' 'т' ' ' 'б' 'ы' 'т' 'ь']

	// var x, y []int
	// for i := 0; i < 10; i++ {
	// 	y = appendInt(x, i)
	// 	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }
	// 0 cap=1 [0]
	// 1 cap=2 [0 1]
	// 2 cap=4 [0 1 2]
	// 3 cap=4 [0 1 2 3]
	// 4 cap=8 [0 1 2 3 4]
	// 5 cap=8 [0 1 2 3 4 5]
	// 6 cap=8 [0 1 2 3 4 5 6]
	// 7 cap=8 [0 1 2 3 4 5 6 7]
	// 8 cap=16        [0 1 2 3 4 5 6 7 8]
	// 9 cap=16        [0 1 2 3 4 5 6 7 8 9]

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)
	fmt.Println(x)
	// [1 2 3 4 5 6 1 2 3 4 5 6]
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// Имеется место для роста. Расширяем срез
		z = x[:zlen]
	} else {
		// места для роста нет. Выделяем новый массив. Увеличиваем в два раза оинейной амортизированной сложности
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // Встроенная функция
	}

	z[len(x)] = y
	return z
}

func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	copy(z[len(x):zlen], y)
	return z
}

type IntSlice struct {
	ptr *int
	len, cap int
}