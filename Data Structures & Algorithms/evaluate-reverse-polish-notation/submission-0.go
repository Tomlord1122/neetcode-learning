func evalRPN(tokens []string) int {
    stack := []int{}

    for i := 0; i < len(tokens); i++{
        if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/"{
            val, _ := strconv.Atoi(tokens[i])
            stack = append(stack, val)
        } else {
            n := len(stack)
            val2, val1 := stack[n-1], stack[n-2]
            stack = stack[:n-2]
            switch tokens[i]{
                case "+":
                    stack = append(stack, val1+val2)
                case "-":
                    stack = append(stack, val1-val2)
                case "*":
                    stack = append(stack, val1*val2)
                case "/":
                    stack = append(stack, val1/val2)
            }
        }
    }
    return stack[0]
}
