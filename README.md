# Unit Testing Framework in Go

## Usage

```go
func TestConcurrentList_MapParallelInPlace(t *testing.T) {
	should := fCon("MapParallelInPlace", t)
	should("apply func to each item and modify list in place", New(2, 3, 4), func() interface{} {
		xs := New(1, 2, 3)
		xs.MapParallelInPlace(func(x interface{}, idx uint) interface{} {
			return x.(int) + 1
		})
		return xs
	})
	should("do nothing for empty lists", New(), func() interface{} {
		xs := New()
		xs.MapParallelInPlace(func(x interface{}, idx uint) interface{} {
			return x.(int) + 1
		})
		return xs
	})
}
```
