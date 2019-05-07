package testbatch

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/example/stringutil"
)

func TestBatch(t *testing.T) {
	bufCh := make(chan int, 20)
	go func() {
		for {
			var items []int
			items = append(items, <-bufCh)

			// As we want to batch get maximum 10 items, 9 items remain to get.
			remains := 9

		Remaining:
			for i := 0; i < remains; i++ {
				select {
				case item := <-bufCh:
					items = append(items, item)
				default:
					break Remaining
				}
			}

			// The batch processing. Here we just log output.
			t.Log("Items:", items)
			fmt.Println("Items:", items)
		}
	}()

	for i := 0; i < 50; i++ {
		bufCh <- i
		t.Log("Push:", i)
	}

	time.Sleep(time.Second)
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func ExampleReverse() {
	fmt.Println(stringutil.Reverse("hello"))
	// Output: olleh
}
