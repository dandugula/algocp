package main

import "fmt"
//import "sync"
//import "math/rand"
//import "time"
// Linear search
/*
func main () {
    a := []int{1,2,3,4}
    n := 6
    if linear_search (a, n) {
        fmt.Printf("Found needle (%v)\n", n)
    } else {
        fmt.Println("Needle not found")
    }
}
*/
// Binary search
/*
func main () {
    a := []int{1,2,3,4}
    n := 6
    if bs_list (a, n) {
        fmt.Printf("Found needle (%v)\n", n)
    } else {
        fmt.Println("Needle not found")
    }
}
*/

// Two Crystal balls
/*
func main () {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    idx := r.Intn(1000)
    fmt.Println("idx =", idx)
    var data [1000]bool
    for i := idx; i < len(data); i++ {
        data[i] = true
    }
    if idx == two_crystal_balls(data) {
        fmt.Println(idx)
    } else {
        fmt.Println(-1)
    }
}
*/

// Bubble sort
/*
func main() {
    arr := []int{5,3,7,4,2}
    fmt.Println(arr)
    bubble_sort(arr)
    fmt.Println(arr)
}
*/
/*
func main() {
    aList := &List {
        val : 90,
        next : &List{
            val: 100,
            next: nil,
        },
    }
    aList.enqueue(200)
    //aList.enqueue(100)
    fmt.Println(aList)
    fmt.Println(aList.dequeue())
    fmt.Println(aList)
    fmt.Println(aList.peek())
}
*/

/*
func main () {
    wg := sync.WaitGroup{}
    wg.Add(5)
    nums := []int{1,2,3,4,5}

    //for x := range 5 {
    for _, x := range nums {
        go func() {
            fmt.Println(x)
            wg.Done()
        }()
    }

    wg.Wait()
}
*/
/*
func sumOfInt(input []int) int {
    inc := 0
    for _, val := range input {
        inc += val
    }
    return inc
}

//Generic
type Number interface {
    int64 | float64
}

func sumOf[T Number](input []T) T {
    var inc T
    for _, val := range input {
        inc += val
    }
    return inc
}

func main () {
    fmt.Printf("%v\n", sumOf([]int64{1,2,3,4,5}))
    fmt.Printf("%v\n", sumOf([]float64{1.0,2.3,3.6,4.7,5.0}))
}
*/
/*
// Quicksort
func main() {
    arr := []int{5,3,7,4,2}
    fmt.Println(arr)
    quick_sort(arr)
    fmt.Println(arr)
}
*/
/*
func main() {
    var lruc LRUCache = NewCache(3)
    lruc.insert("abc", &CacheNode{v:10})
    lruc.insert("def", &CacheNode{v:11})
    lruc.insert("ghi", &CacheNode{v:12})
    lruc.insert("jkl", &CacheNode{v:13})
    if n := lruc.get("abc"); n != nil {
        fmt.Println(n.v)
    } else {
        fmt.Println("Value not found")
    }
}
*/

func main() {
    lru := Constructor(2)
    lru.Put(1, 1)
    lru.Put(2, 2)
    fmt.Println(lru.Get(1))
    lru.Put(3, 3)
    fmt.Println(lru.Get(2))
}
/*
type Node struct {
    v int
}

func (n *Node) init() *Node{
    if n == nil {
        n = &Node {
            v : 10,
        }
    }
    return n
}

func main () {
    var n *Node = nil
    n = n.init()
    if n != nil {
        fmt.Println(n.v)
    } else {
        fmt.Println("nil value")
    }
}
*/
