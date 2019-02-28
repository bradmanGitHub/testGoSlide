package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //รับค่าจาก console (short declaration ด้วย := )
	if scanner.Scan() {                   //หลัง if ไม่ต้องมีวงเล็บ
		input := scanner.Text() // return string

		/*
			go assignn ค่าให้ตัวแปรได้ทีละหลายค่า และ function ใน go return ได้ทีละหลายค่า
			เช่น function Atoi return 2 ค่านะ (num และ error เอามา chk ว่ามี error ไหมนั่นเอง)
		*/
		//num, err := strconv.Atoi(input) //convert string ที่ได้มาเป็น integer (Alphabet to Integer -> Atoi)
		//if err != nil {
		if num, err := strconv.Atoi(input); err != nil {
			fmt.Println("It's not a number")
		} else {
			if num%3 == 0 && num%5 == 0 {
				fmt.Println("FizzBuzz")
			} else if num%3 == 0 {
				fmt.Println("Fizz")
			} else if num%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(num)
			}
		}
	}
}
