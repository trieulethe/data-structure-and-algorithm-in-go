const climbStair = (n) => {
    let state1 = 1
    let state2 = 2
    if (n == state1) return state1
    if (n == state2) return state2
    for (let i = 3; i <= n; i+=2) {
        state1 = state1 + state2
        state2 = state1 + state2
    }
    if (n % 2 == 0) {
        return state2
    }
    return state1

}

climbStair(3)
climbStair(4)
climbStair(5)
climbStair(6)
climbStair(7)