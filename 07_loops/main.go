package main

import "fmt"

func main ()  {
 i := 1

 for i < 10 {
	 fmt.Println(i)

	 i = i + 1;
 }


 for i := 1; i <= 10; i++ {
 fmt.Printf("Number %d\n", i)
 }



 //FizzBuzz
  for i := i; i <= 100; i++ {
	  if i % 3 == 0 {
		  fmt.Println("fizz")
	  } else if i % 5 == 0 {
		  fmt.Println("buzz")
	  } else if i % 15 == 0 {
		  fmt.Println("Fizzbuzz")
	  } else {
		  fmt.Println(i)
	  }
  }


}