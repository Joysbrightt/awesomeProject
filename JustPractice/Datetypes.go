package main

import "fmt"

func main() {
	x := 5
	x += 1
	fmt.Println(x)

	//write a program that converts from Fahrenheit into Celcius
	fahrenheit := 273
	var convertTemp = (fahrenheit - 32) * 5 / 9
	fmt.Println(convertTemp)

	// write a  program that converts from feet into meters

	feet := 0.3048
	var contertMeter = feet * 6.7
	fmt.Println(contertMeter)

	// using loop to perform iteration

	for i := 1; i <= 20; i++ {
		//fmt.Println(i + 2)

	}

	// todo write a program that print out all the number evenly divisible by 3
	// todo between 1 and 100

	for evenly := 3; evenly <= 100; evenly++ {
		if evenly%3 == 0 {
			//fmt.Println(evenly)
		}

	}

	// todo write a program that prints the numbers from to 100
	// todo the multiple of 3 print "Fizz"
	// todo for the multiple of 5 print 'FizBuzz'

	for iteration := 3; iteration <= 100; iteration++ {

		if iteration%15 == 0 {
			fmt.Print("Fizzbuzz\n")
		} else if iteration%5 == 0 {
			fmt.Print("Buzz\n")
		} else if iteration%3 == 0 {
			fmt.Print("Fizz")
		} else {
			fmt.Println(iteration)
		}

		var x [5]float64
		x[0] = 98
		x[1] = 93
		x[2] = 77
		x[3] = 82
		x[4] = 83
		//var total float64 = 0
		for i := 0; i < 5; i++ {
			//total += x[i]
		}
		//fmt.Println(total / 5)

		//// this code gives an undefined error
		//var total float64 = 0
		//for _, i := 0; i < len(x); i++ {
		//	total += x[i]
		//}
		//fmt.Println(total / float64(len(x)))

	}

	//var x [5]float64
	//x[0] = 98
	//x[1] = 93
	//x[2] = 77
	//x[3] = 82
	//x[4] = 83
	//var total float64 = 0
	//for i := 0; i < 5; i++ {
	//	for _, value := range x {
	//		total += value
	//	}
	//	fmt.Println(total / float64(len(x)))
	//}

	//	todo SLICING
	//	var x []float64
	// todo slicing using makr
	//x := make([]float64, 5,10.13, 134,123,1)

	//	todo slicing using [low: high]
	//arr :=[]float64{1,2,3,4,5}
	//x :=[0:5]
}

//for _, value := range x {
//total += value
//}
//fmt.Println(total / float64(len(x)))for _, value := range x {
//total += value
//}
//fmt.Println(total / float64(len(x)))
