package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}

	results := make(chan FreqMap, 10)

	for _, value := range s {
		go func(str string) {
			results <- Frequency(str)
		}(value)
	}

	for range s {
		for pl, value := range <-results {
			m[pl] += value
		}
	}

	return m
}
