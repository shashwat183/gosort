# gosort

Go package with methods to sort int arrays using various different algorithms

## Description

A go package exporting sort function implemented using various different sorting algorithms.

The following sorting algorithms are covered :-

* bubble sort
* quick sort
* merge sort
* selection sort

## How to install

Before installing this project make sure you have Go Tools installed on your machine.

To install do

```bash
go get github.com/shashwat183/gosort
```

## How to use

```go
package main

import (
  "fmt"
  "math/rand"
  "github.com/shashwat183/gosort/sort"
)

func main() {
  array := []int{1, 3, 6, 3, 7 ,5}
  // use merge sort to sort array
  sorted = sort.Mergesort(array)
}
```

## Resources to learn Go

[A Tour of Go](https://tour.golang.org/welcome/1)

### Single youtube video to learn it all

[Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI)(This is what i used and then wrote this program)
Note - This is a good resource if you are not a beginner in programming.

[Great video covering all the important concepts](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=2430s)

### Github Resources

Very good github repo with details on how to structure your programs
[gostart](https://github.com/alco/gostart#canonical)
