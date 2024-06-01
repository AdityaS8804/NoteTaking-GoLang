package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func viewMeta() []string {
	dat, err := os.ReadFile("./notes/meta.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	info := strings.Split(content, "\n")[1:]
	return info

}
func multilineReader(scanner *bufio.Scanner) []string {
	input := []string{}
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			input = append(input, text)
		} else {
			break
		}
	}
	fmt.Println(" ")
	return input
}
func CreateFile() {

	var name string
	var content string
	fmt.Printf("Enter file name : ")
	fmt.Scan(&name)
	file, err := os.Create("./notes/" + name + ".txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter file content :")
	scanner := bufio.NewScanner(os.Stdin)

	content = strings.Join(multilineReader(scanner)[:], "\n")
	fmt.Println(content)
	io.WriteString(file, content)
	file.Close()
	details := viewMeta()
	id2 := "0"
	if len(details)-1 >= 0 {
		i := strings.Split(details[len(details)-1], "\t")[0]

		id, err := strconv.Atoi(i)
		if err != nil {

			panic(err)
		}
		id = id + 1
		id2 = fmt.Sprint(id)
	}
	metaFile, err := os.OpenFile("./notes/meta.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	finStr := "\n" + id2 + "\t" + name + ".txt"
	if _, err := metaFile.WriteString(finStr); err != nil {
		panic(err)
	}
	metaFile.Close()

}
func ViewNote(id uint8) string {
	info := viewMeta()
	idStr := fmt.Sprint(id)
	temp := ""
	for _, s := range info {
		if strings.Contains(s, idStr) {
			temp = s
		}
	}
	if temp == "" {
		panic("No such Note present")
	}
	noteName := strings.Split(temp, "\t")[1]
	files, err := os.ReadDir("./notes")
	if err != nil {
		panic(err)
	}
	content := ""
	for _, f := range files {
		if f.Name() == noteName {
			dat, err := os.ReadFile("./notes/" + noteName)
			if err != nil {
				panic(err)
			}
			content = string(dat)
		}
	}
	if content == "" {
		panic("No file with that name")
	}
	return content

}
func createMeta() {
	files, err := os.ReadDir("./notes")
	if err != nil {
		panic(err)
	}
	check := false
	for _, f := range files {
		check = strings.Contains(f.Name(), "meta.txt")
		if check {
			break
		}
	}
	if !check {
		f, err := os.Create("./notes/meta.txt")
		if err != nil {
			panic(err)
		}
		id := 0
		heading := "ID\tName"
		for _, f := range files {
			heading += "\n" + fmt.Sprint(id) + "\t" + f.Name()
			id += 1
		}
		io.WriteString(f, heading)
	}
}
func main() {
	fmt.Println("Welcome to the best CLI note taking experience\nSelect an appropriate option")
	createMeta()
	var choice uint8
	for {
		fmt.Printf("1 - Create new note\n2 - View a particular note\n3 - Update a note\n4 - Delete a note\n5 - Exit\nChoose : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			CreateFile()
		case 2:
			var id uint8
			fmt.Printf("Enter the note ID : ")
			fmt.Scan(&id)
			content := ViewNote(id)
			fmt.Println(content)
		default:
			os.Exit(1)
		}
	}
}
