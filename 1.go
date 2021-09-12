package main
 
import (
    "fmt"
  )
var n, x, i,divide float64

func suma(n float64) float64 {
  if n == 1 {
    i = 20
    fmt.Println(i)
    return i
  } else {
    divide = 20/n
    fmt.Print(divide, "+")
    return divide + suma(n-1)
  }
}

func sumb(n float64) float64 {
  if n == 1 {
    i = 100
    fmt.Println(i)
    return i
  } else {
    divide = 100/n
    fmt.Print(divide, "+")
    return divide + sumb(n-1)
  }
}

func sumc(n float64) float64 {
  if n == 1 {
    i = 1
    fmt.Println(i)
    return i
  } else {
    divide = 1/n
    fmt.Print(divide, "+")
    return divide + sumc(n-1)
  }
}

func sumd(n float64, x float64) float64 {
  if n == 1 {
    i = 1/(2*x)
    fmt.Println(i)
    return i
  } else {
    divide = n/((n*n)*x)
    fmt.Print(divide, "+")
    return divide + sumd(n-1,x)
  }
}

func sume(n float64, x float64) float64 {
  if n == 1 {
    i = 2*x/1
    fmt.Println(i)
    return i
  } else {
    divide = (n*n)*x/n
    fmt.Print(divide, "+")
    return divide + sume(n-1,x)
  }
}

func main(){
//driver code
 n = 5
 x = 2
 fmt.Println("Result series a : ", suma(n))
 fmt.Println("Result series b : ", sumb(n))
 fmt.Println("Result series c : ", sumc(n))
 fmt.Println("Result series d : ", sumd(n,x))
 fmt.Println("Result series e : ", sume(n,x))
}