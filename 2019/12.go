package main

import "fmt"

func main() {
    type vec struct { x, y, z int }
    var moons, vel [4]vec
    for i := 0; i < 4; i++ {
        fmt.Scanf("<x=%d, y=%d, z=%d>%n", &moons[i].x, &moons[i].y, &moons[i].z)
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
    fmt.Println(energy)
}
