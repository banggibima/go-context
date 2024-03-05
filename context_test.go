package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

type myKeyType string

const (
	keyB myKeyType = "b"
	keyC myKeyType = "c"
	keyD myKeyType = "d"
	keyE myKeyType = "e"
	keyF myKeyType = "f"
	keyG myKeyType = "g"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, keyB, "B")
	contextC := context.WithValue(contextA, keyC, "C")

	contextD := context.WithValue(contextB, keyD, "D")
	contextE := context.WithValue(contextB, keyE, "E")

	contextF := context.WithValue(contextC, keyF, "F")
	contextG := context.WithValue(contextF, keyG, "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextF.Value("a"))
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}
