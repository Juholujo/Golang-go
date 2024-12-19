//Пример 2 (из WBTypes15.1.go)
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(4) // Можно менять на 1 или другой параметр для наблюдения результатов
    wg := sync.WaitGroup{}
    wg.Add(7)

    for i := 0; i < 9; i++ {
        i := i // Переназначаем i для правильной передачи в горутину
        go func(i int) {
            defer wg.Done()
            fmt.Println("Почему, КОЛЯ?", i)
        }(i)
    }

    wg.Wait()
    fmt.Println("Паника")
}

//Когда мы ставим runtime.GOMAXPROCS(4), это значит, что Go может использовать сразу 4 потока.
//Как-будто  есть четыре руки, и можно параллельно делать четыре задачи.
//Тогда то, какую задачу закончишь первой, будет зависеть от того, как быстро каждая рутина отработает.
//Поэтому порядок вывода сообщений может быть любым, всё происходит почти одновременно, и Go никак не гарантирует, в каком порядке горутины закончатся.
