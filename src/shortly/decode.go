package shortly

import (
  "math"
  "strings"
)

func Decode(key string) int {
  // Convert a base-62 number to base-10.

  tot := 0

  for i, _ := range key {
    var idx int = strings.IndexByte(Alphabet, key[i])
    tot += idx * int(math.Pow(62, float64(i)))
  }

  return tot
}

