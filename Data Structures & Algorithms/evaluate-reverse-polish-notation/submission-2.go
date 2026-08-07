func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range tokens{
        if token != "+" && token != "-" && token != "*" && token != "/"{
            val, _ := strconv.Atoi(token)
            stack = append(stack, val)
        } else {
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            switch token{
                case "+":
                    stack = append(stack, a+b)
                case "-":
                    stack = append(stack, a-b)
                case "*":
                    stack = append(stack, a*b)
                case "/":
                    stack = append(stack, a/b)
            }
        }
    }
    return stack[0]
}
