package main

import (
  "fmt"
  "math"
  "strings"
  "strconv"
)

var Alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func BaseConvert(number int, base int) []int {
   // Convert base-10 number to arbitrary base.

   digits := []int{}

   for (number > 0) {
    digits = append(digits, (number % base))
    number /= base
  }

  // NOTE: Reversing an array is much more complicated in go. Since we
  //       are simply interested in mapping to/from record ids it does
  //       not matter if the numbers are "backwards".

  return digits
}

func Encode(number int) string {
   // Convert a base-10 number to base-62.

   digits := []int{}
   digits = BaseConvert(number, 62)

   chars := []byte{}
   for _, c := range digits {
      chars = append(chars, Alphabet[c])
   }

   // Join the chars and return the base-62 encoded string.
   s := string(chars[:])
   return s
}

func Decode(key string) int {
  // Convert a base-62 number to base-10.

  tot := 0

  for i, _ := range key {
    var idx int = strings.IndexByte(Alphabet, key[i])
    tot += idx * int(math.Pow(62, float64(i)))
  }

  return tot
}


func main() {
  input  := 25848298124 // test value obtained from python test program
  expected := "qiZsnC"  // expected base-62 encoded value

  fmt.Println(BaseConvert(input, 62))

  encoded := Encode(input)

  fmt.Println("output  : " + encoded)
  fmt.Println("expected: " + expected)

  decoded := Decode(encoded)

  fmt.Println("input   : " + strconv.Itoa(input))
  fmt.Println("decoded : " + strconv.Itoa(decoded))
}
