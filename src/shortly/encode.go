package shortly

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
