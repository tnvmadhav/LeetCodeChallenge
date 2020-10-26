func countAndSay(n int) string {
    var myAnswer [31][]string
    
    myAnswer[0] = []string{"1"}
    for i := 1; i < n; i++ {
        myAnswer[i] = calculateNext(myAnswer[i-1])
    }
    return strings.Join(myAnswer[n-1], "")
}

func calculateNext(A []string) []string {
    var count = 1
    var myAnswer []string
    for i := range A {
        if i == len(A) - 1 {
            myAnswer = append(myAnswer, []string{strconv.Itoa(count),string(A[i])}...)
            break
        }
        if A[i] == A[i+1] {
            count++
        } else {
            myAnswer = append(myAnswer, []string{strconv.Itoa(count),string(A[i])}...)
            count = 1
        }
    }
    return myAnswer
}
