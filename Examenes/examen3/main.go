package main

import (
	"bufio"
	"examen3/structs"
	"log"
	"os"
)


type chars struct {
	value string
	freq int
}

func getValue(s []chars, str string) int {
	for i := 0; i < len(s); i++ {
		if s[i].value == str {
			return i
		}
	}

	return -1
}

func increaseFreq(s []structs.TData, str string) {
	for i := 0; i < len(s); i++ {
		if s[i].Char == str {
			s[i].Freq++
		}
	}

}

func main() {
	
	treeList := structs.CreateList()

	file, err := os.Open("TextoPlano.txt")
	if err != nil {
		log.Fatal(err)
	}
	fScan := bufio.NewScanner(file)
	fScan.Split(bufio.ScanLines)
	var lines []string
	for fScan.Scan() {
		lines = append(lines, fScan.Text())
	}

	file.Close()

	str := []rune(lines[0])
	var char[]string
	var charList []structs.TData	

	for i := 0; i < len(str); i++ {
		char = append(char, string(str[i]))
	}

	for i := 0; i < len(char); i++{
		if !treeList.Contains(char[i])  {
			treeList.AddRight(structs.DataList{Valor: char[i], Freq: 1})
			charList = append(charList, structs.TData{Char: char[i], Freq: 1})
		} else {
			treeList.IncreaseFreq(char[i])
			increaseFreq(charList, char[i])
		}
	}

	treeList.Huffman(charList)

	// fmt.Println(charList)

}
