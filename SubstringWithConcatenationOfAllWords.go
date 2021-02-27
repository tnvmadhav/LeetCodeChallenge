func findSubstring(s string, words []string) []int {
    answer := []int{}
    
    if len(words) == 0 || len(s) == 0 {
        return []int{}
    }
    
    // Get number of occurances of all words
    wordCount := make(map[string]int)
    for i := range words {
        wordCount[words[i]]++
    }
    
    // Get the word concatenationWindow using below formula (since each word in the words slice is of same length)
    concatWindow := len(words) * len(words[0])
    wordLength := len(words[0])
    
    for i := 0; i < len(s) - concatWindow + 1; i++ {
        instanceCount := make(map[string]int)
        for j := i; j < i + concatWindow; j += wordLength {
            instanceCount[s[j:j+wordLength]]++
        }
        if reflect.DeepEqual(wordCount, instanceCount) {
            answer = append(answer, i)
        }
    }
    return answer
}
