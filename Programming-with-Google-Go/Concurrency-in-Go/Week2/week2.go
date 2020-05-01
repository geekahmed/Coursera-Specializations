package main

import "time"

var count int

func race() {
	count++
}

func main() {
	go race()
	go race()
	time.Sleep(1 * time.Second)
}

/*

Explaination:

First Output from race detector:
==================
WARNING: DATA RACE
Read at 0x00000056cc78 by goroutine 7:
  main.race()
      /home/geekahmed/Desktop/Coursera-Specializations/Programming-with-Google-Go/Concurrency-in-Go/Week2/week2.go:8 +0x3a

Previous write at 0x00000056cc78 by goroutine 6:
  main.race()
      /home/geekahmed/Desktop/Coursera-Specializations/Programming-with-Google-Go/Concurrency-in-Go/Week2/week2.go:8 +0x56

Goroutine 7 (running) created at:
  main.main()
      /home/geekahmed/Desktop/Coursera-Specializations/Programming-with-Google-Go/Concurrency-in-Go/Week2/week2.go:13 +0x5a

Goroutine 6 (finished) created at:
  main.main()
      /home/geekahmed/Desktop/Coursera-Specializations/Programming-with-Google-Go/Concurrency-in-Go/Week2/week2.go:12 +0x42
==================
Found 1 data race(s)


-- What is race  condetion?
A data race happens when two goroutines access the same variable concur­rently, and at least one of the accesses is a write.


*/
