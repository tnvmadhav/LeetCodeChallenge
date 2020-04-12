func backspaceCompare(S string, T string) bool {
    
    value1, value2 := TextEditorReturn(S), TextEditorReturn(T)
    if len(value1) != len(value2) {
        return false
    }
    for i := range value1 {
        if value1[i] != value2[i] {
            return false
        }
    }
    return true
}



func TextEditorReturn(s string) []string {
    var myArray [101]string
    var index int
    
    for i := range s{
        myArray[index] = string(s[i])
        if string(s[i]) == string('#') {
            if index != 0 {
                index -= 2
            } else {
                index -= 1
            }
        }
        index += 1
    }
    return myArray[:index]
}
