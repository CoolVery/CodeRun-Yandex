package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Переменная хранит количество пар
	var countPairWords int
	//Переменные хранят введенные синонимы
	var firstWord, secondWord string
	//Переменная хранит слово, по которому найдется его синоним
	var searchByWord string

	reader := bufio.NewReader(os.Stdin)
	dict := make(map[string]string)

	fmt.Fscan(reader, &countPairWords)
	for i := 0; i < countPairWords; i++ {
		//Читаем оба слова и делаем две записи в словарь с реверсом
		fmt.Fscan(reader, &firstWord, &secondWord)
		dict[firstWord] = secondWord
		dict[secondWord] = firstWord
	}
	fmt.Fscan(reader, &searchByWord)
	fmt.Println(dict[searchByWord])
}
