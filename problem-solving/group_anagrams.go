import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	wordsMap := map[string]answer{}

	for _, word := range strs {
		sortedWord := sortWord(word)
		wordsMap[sortedWord] = answer{
			Freq:  wordsMap[sortedWord].Freq + 1,
			Words: append(wordsMap[sortedWord].Words, word),
		}
	}

	for _, val := range wordsMap {
		var arr []string
		for _, word := range val.Words {
			arr = append(arr, word)
		}
		result = append(result, arr)
	}

	return result
}

type answer struct {
	Freq  int
	Words []string
}

func sortWord(s string) string {
	bytes := []byte(s)

	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	return string(bytes)
}
