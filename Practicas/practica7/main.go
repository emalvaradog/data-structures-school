package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	a "practica7/structs"
)

func main() {
	q := a.CreateDoubleQueue()

	files := make(map[string]string)
	var opc int
	for opc != 10 {
		fmt.Println("Choose an option")
		fmt.Println("0.- Create new file")
		fmt.Println("1.- Add file to left");
		fmt.Println("2.- Add file to right")
		fmt.Println("3.- Remove file from left")
		fmt.Println("4.- Remove file from right")
		fmt.Println("5.- Get first file")
		fmt.Println("6.- Get last file")
		fmt.Println("7.- Show all files in queue")
		fmt.Println("8.- Clear all queue")
		fmt.Println("9.- Show created files")
		fmt.Println("10.- Send files")
		fmt.Scanln(&opc)
		fmt.Printf("\n")
		switch opc {
			case 0:
				var fileName string
				var content string
				fmt.Println("-- Create new file --")
				fmt.Println("File name:")
				fmt.Scanln(&fileName)

				fileName += ".txt"
				
				f, err := os.Create(fileName );
				if err != nil {
					fmt.Println("Error creating file")
					break
				}

				fmt.Printf("\n")
				
				fmt.Println("File content:")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				content = scanner.Text()

				_, wE := f.WriteString(content)

				if wE != nil {
					fmt.Println("Error writing file")
					break;
				}

				fmt.Printf("\n")
				fmt.Println("File created succesfully")
				fmt.Printf("\n")
				files[fileName] = content
				f.Close()
				
			case 1: 
				fmt.Println("-- Add file to left --")

				fmt.Println("Choose file:")
				for i := range files {
					fmt.Println("\t", i);
				}

				var opt string
				_, err := fmt.Scanln(&opt)

				if err != nil {
					fmt.Println("Error getting input")
				}

				if _, found := files[opt]; found {
					fmt.Println("OPT", found)
					newData := a.Data{Value: files[opt], Title: opt}
					q.PushFront(newData)
					fmt.Println("File added succesfully")
				} else {
					fmt.Println("Invalid input")
				}
				fmt.Printf("\n")

			case 2:
				fmt.Println("-- Add file to right --")

				fmt.Println("Choose file:")
				for i:= range files {
					fmt.Println("\t", i);
				}

				var opt string
				_, err := fmt.Scanln(&opt)

				if err != nil {
					fmt.Println("Error getting input")
				}

				if _, found := files[opt]; found {
					newData := a.Data{Value: files[opt], Title: opt}
					q.PushBack(newData)
					fmt.Println("File added succesfully")
				} else {
					fmt.Println("Invalid input")
				}
					
				fmt.Printf("\n")

			case 3:
				fmt.Println("-- Remove file from left --")
				q.PopFront()
				fmt.Println("File removed succesfully")
				fmt.Printf("\n")

			case 4:
				fmt.Println("-- Remove file from right --")
				q.PopBack()
				fmt.Println("File removed succesfully")
				fmt.Printf("\n")

			case 5:
				fmt.Println("-- Get first file --")
				fmt.Println(q.GetHead())
				fmt.Printf("\n")

			case 6: 
				fmt.Println("-- Get last file --")
				fmt.Println(q.GetTail())
				fmt.Printf("\n")

			case 7:
				fmt.Println("-- Show all files in queue --")
				fmt.Println(q)
				fmt.Printf("\n")

			case 8:
				fmt.Println("-- Clear queue --")
				q.Clear()
				fmt.Printf("\n")

			case 9:
				fmt.Println("-- Show files --")
				for i, file := range files {
					fmt.Println("\t", i, "-" , file)
				}
				fmt.Printf("\n")

			case 10:
				fmt.Println("-- Send files --")
				q.Send()				
				
				for i := range files {
					e := os.Remove(i)
					if e != nil {
						log.Fatal(e)
					}
				}

		}
	}
}