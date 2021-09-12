package main
 
import (
    "fmt"
    "unicode"
    "bufio"
    "os"
    "strings"
  )

var count int = 0

func checkLower(stringArray []rune, n int) int {
    if n-1 == -1 {
      return count
    } else if n == 0 && !unicode.IsLower(stringArray[0]) {
      count+=0
      return count
    } else if unicode.IsLower(stringArray[n-1]) == true {
      count +=1
      return checkLower(stringArray, n-1)
    } else {
      return checkLower(stringArray, n-1)
    }     
}


func checkNumber(stringArray []rune, n int) int {
    if n-1 == -1 {
      return count
    } else if n == 0 && !unicode.IsNumber(stringArray[0]) {
      count+=0
      return count
    } else if unicode.IsNumber(stringArray[n-1]) == true {
      count +=1
      return checkNumber(stringArray, n-1)
    } else {
      return checkNumber(stringArray, n-1)
    }
}

func main(){
//driver code
  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Println("Input some string :  ")
  input, _ := consoleReader.ReadString('\n')
  input = strings.TrimSuffix(input, "\n")
  stringArray := []rune(input)
  var n int = len(input)
  fmt.Println("List of non-capital alphabet(s) : ", checkLower(stringArray, n))
  fmt.Println("List of number(s) : ", checkNumber(stringArray, n))
}