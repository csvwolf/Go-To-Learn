package letter

func ConcurrentFrequency(wordList []string) FreqMap {
  m := FreqMap{}
  channel := make(chan FreqMap, len(wordList))
  for _, words := range wordList {
    go func(words string) { channel <- Frequency(words) }(words)
  }

  for range wordList {
    for key, value := range <-channel {
      m[key] += value
    }
  }
  return m
  }
