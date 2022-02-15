package main

import (
	"fmt"
	"time"
)

// Send the sequence 2, 3, 4, ... to returned channel
func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		println(<-primes)
		time.Sleep(1 * time.Second)
	}

	// slice and nil

	// 类型声明方式
	var v []int
	fmt.Println(len(v), cap(v)) // 0 0
	fmt.Println(v == nil)       // true

	v = append(v, 1)
	fmt.Println(len(v), cap(v)) // 1 1
	v = v[:0]
	fmt.Println(len(v), cap(v)) // 0 1
	fmt.Println(v == nil)       // false

	v = []int{}
	fmt.Println(len(v), cap(v)) // 0 0
	fmt.Println(v == nil)       // false

	// 字面量声明方式
	v2 := []int{}
	fmt.Println(len(v2), cap(v2)) // 0 0
	fmt.Println(v2 == nil)        // true

	// slice params
	a := []int{1, 2, 3}
	fmt.Println(a)
	changeValue(a)
	fmt.Println(a)

	a2 := []int{1, 2, 3}
	fmt.Println(a2)
	notChangeValue(a2)
	fmt.Println(a2)
}

func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}

func changeValue(a []int) {
	a[0] = 100
}

func notChangeValue(a []int) {
	// 扩容改变了引用
	a = append(a, 1024)
	a[0] = 100
}
