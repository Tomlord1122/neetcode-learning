func evalRPN(tokens []string) int {
    stack := []int{}

    for _, val := range tokens{
        if val != "+" && val != "-" && val != "*" && val != "/"{
            num, _ := strconv.Atoi(val)
            stack = append(stack, num)
        } else {
            n := len(stack)
            b := stack[n-1]
            a := stack[n-2]
            stack = stack[:n-2]
            switch val{
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
