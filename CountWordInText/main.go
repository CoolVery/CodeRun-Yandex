package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Создаем множество
	uniqueWord := make(map[string]struct{})	
	//Создаем сканер - обертку над потоком ввода
	scanner := bufio.NewScanner(os.Stdin)
	//Задаем, что сканер будет сплитить данные по словам
	scanner.Split(bufio.ScanWords)
	//Цикл, что пока сканер читает...
	for scanner.Scan() {
		//Проверяем его текст в множестве
		uniqueWord[scanner.Text()] = struct{}{}
	}
	//Выводим длину множества
	fmt.Println(len(uniqueWord))
}