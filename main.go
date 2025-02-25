package main

import "fmt"
import "context"
import "time"
import "runtime"

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
/*
func main() {
    lru := Constructor(2)
    fmt.Println(lru.Get(2))
    lru.Put(2, 6)
    fmt.Println(lru.Get(1))
    lru.Put(1, 5)
    lru.Put(1, 2)
    fmt.Println(lru.Get(1))
    fmt.Println(lru.Get(2))
}
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
/*
func main() {
    fmt.Println("Adding ints: ", add(3, 4))
    fmt.Println("Adding floats: ", add(3.5, 4.57))
}
*/
/*
func main() {
    a,b := setUpViewIdItemIdMap(generateViewIds());
    /*
    fmt.Println(len(a))
    for k, v := range a {
        fmt.Println(k, ", ", len(v))
        fmt.Println(v[0])
        fmt.Println(v[1])
        break
    }
    fmt.Println(len(b))
    for i := range b {
        fmt.Println(i)
        fmt.Println(i+1)
        break;
    }
    count := 0
    c := make(map[ItemId]struct{})

    for k := range b {
        if count >= 5000 {
            break
        }
        count++
        c[k] = struct{}{}
    }
    d := setUpItemStatsMap(c)
    TransformAndProduce(d, a)
}
*/
/*
func main() {
	myList := &GenericList[int]{head: nil}
	//myList = nil
	myList.add(10)
	myList.add(20)
	myList.add(30)
	fmt.Println(myList.count())
	myList.add(40)
	myList.print()
	fmt.Println()
	myList.remove()
	myList.print()
}
func main() {
	a := []int{5, 3, 6, 7, 4}
	//a := []int{-8}
	//a := []int{-190, -189}
	res := maxDiffPriorElem(a)
	fmt.Println(res)
}

func main() {
    deferExp()
}

func Backward[E any](s []E) func(func(int, E) bool) {
    return func(yield func(int, E) bool) {
        for i := len(s)-1; i >= 0; i-- {
            if !yield(i, s[i]) {
                return
            }
        }
    }
}
func main() {
//    s := []string{"hello", "world"}
    in := []int{100, 99}
    for i, x := range Backward(in) {
        fmt.Println(i, x)
    }
}
*/
/*
func main() {
    f := fib()
    fmt.Println(f(), f(), f(), f(), f())
}
*/
type token struct{}

func consumer(ctx context.Context, in <-chan token) {
	for {
		select {
		case <-in:
			// do stuff
		case <-time.After(time.Hour):
			// log warning
		case <-ctx.Done():
			return
		}
	}
}

func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	tokens := make(chan token)
	go consumer(ctx, tokens)

	memBefore := getAlloc()

    for range 100000 {
        tokens <- token{}
    }

	memAfter := getAlloc()

    memUsed := memAfter - memBefore

    fmt.Printf("Memory used: %d KB\n", memUsed / 1024)
}
