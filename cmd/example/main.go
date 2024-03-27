package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	// Оголошення змінних для параметрів командного рядка
	expressionPtr := flag.String("e", "", "Вираз для обчислення")
	inputPtr := flag.String("f", "", "Файл з виразом")
	outputPtr := flag.String("o", "", "Файл для результату")
)

func main() {
	
	flag.Parse()

	// Створення екземпляра ComputeHandler
	var handler ComputeHandler
	if *expressionPtr != "" {
		handler.Input = strings.NewReader(*expressionPtr)
	} else if *inputPtr != "" {
		file, err := os.Open(*inputPtr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Помилка відкриття файлу: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		handler.Input = file
	} else {
		fmt.Fprintln(os.Stderr, "Необхідно вказати вираз або файл з виразом")
		os.Exit(1)
	}

	if *outputPtr != "" {
		file, err := os.Create(*outputPtr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Помилка створення файлу: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		handler.Output = file
	} else {
		handler.Output = os.Stdout
	}

	// Обчислення виразу
	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Помилка обробки виразу: %v\n", err)
		os.Exit(1)
	}

	       handler := &lab2.ComputeHandler{
	           Input: {construct io.Reader according the command line parameters},
	           Output: {construct io.Writer according the command line parameters},
	       }
	       err := handler.Compute()

	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
