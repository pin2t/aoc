package main

import "fmt"

func gcd(a, b int) int {
    for b != 0 { a, b = b, a%b }
    if a < 0 { return -a }
    return a
}

func lcm(a, b int) int {
    return a / gcd(a, b) * b
}

func main() {
    type vec struct { x, y, z int }
    var moons, vel [4]vec
    var xvs, yvs, zvs  [4][2]int
    for i := 0; i < 4; i++ {
        fmt.Scanf("<x=%d, y=%d, z=%d>%n", &moons[i].x, &moons[i].y, &moons[i].z)
        xvs[i][0] = moons[i].x
        yvs[i][0] = moons[i].y
        zvs[i][0] = moons[i].z
    }
    var gravity = func (moon vec) (res vec) {
        for i := 0; i < 4; i++ {
            if moons[i] == moon { continue }
            if moons[i].x > moon.x { res.x++ }
            if moons[i].x < moon.x { res.x-- }
            if moons[i].y > moon.y { res.y++ }
            if moons[i].y < moon.y { res.y-- }
            if moons[i].z > moon.z { res.z++ }
            if moons[i].z < moon.z { res.z-- }
        }
        return
    }
    var plus = func (v1, v2 vec) (res vec) {
        res.x = v1.x + v2.x
        res.y = v1.y + v2.y
        res.z = v1.z + v2.z
        return
    }
    var abs = func (a int) int {
        if a < 0 { return -a }
        return a
    }
    var sum = func (v vec) int { return abs(v.x) + abs(v.y) + abs(v.z) }
    for i := 0; i < 1000; i++ {
        for j := 0; j < 4; j++ {
            vel[j] = plus(vel[j], gravity(moons[j]))
        }
        for j := 0; j < 4; j++ {
            moons[j] = plus(moons[j], vel[j])
        }
    }
    var energy = 0
    for i := 0; i < 4; i++ {
        energy += sum(moons[i]) * sum(vel[i])
    }
    fmt.Print(energy)
    var loop = func (flat [4][2]int) (res int) {
        var states = make(map[[4][2]int]bool)
        var st = [4][2]int{ { flat[0][0], flat[0][1] }, { flat[1][0], flat[1][1] },
            { flat[2][0], flat[2][1] }, { flat[3][0], flat[3][1] } }
        for !states[st] {
            states[st] = true
            var nst = [4][2]int{ { st[0][0], st[0][1] }, { st[1][0], st[1][1] },
                { st[2][0], st[2][1] }, { st[3][0], st[3][1] } }
            for j := 0; j < 4; j++ {
                for i := 0; i < 4; i++ {
                    if j == i { continue}
                    if nst[j][0] > nst[i][0] { nst[j][1]-- }
                    if nst[j][0] < nst[i][0] { nst[j][1]++ }
                }
            }
            for j := 0; j < 4; j++ { nst[j][0] += nst[j][1] }
            st = nst
            res++
        }
        return
    }
    fmt.Println("", lcm(lcm(loop(xvs), loop(yvs)), loop(zvs)))
}
