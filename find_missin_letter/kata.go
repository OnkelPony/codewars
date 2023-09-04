package kata

func FindMissingLetter(chars []rune) rune {
  for i := 1; i < len(chars); i++{
    if chars[i] - chars[i - 1] != 1 {
      return chars[i] - rune(1)
    }
  }
  return 108
}