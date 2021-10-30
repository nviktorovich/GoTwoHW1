package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
//Дополните функцию из п.1 возвратом собственной ошибки в случае возникновения
//панической ситуации. Собственная ошибка должна хранить время обнаружения
//панической ситуации. Критерием успешного выполнения задания является наличие
//обработки созданной ошибки в функции main и вывод ее состояния в консоль

func Div(firstNun, secNum int) (ans int, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("деление на ноль невозможно")
			return
		}
	}()
	ans = firstNun / secNum
	return
}

func main()  {
	var a, b int
	fmt.Println("введите через пробел два целых числа")
	_, err := fmt.Scanln(&a, &b)
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	ans, err := Div(a, b)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(ans)

	fmt.Println("Программа успешно завершила работу")
}