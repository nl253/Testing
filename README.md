# Unit Testing Framework in Go

## Usage

```go
var Function = ut.Test("Worker")

func TestWorker_Start(t *testing.T) {
	should := Function("Start", t)
    expect := rand.Int()
    should("start & receive msg over input channel", expect, func() interface{} {
        worker := New(func(in *stream.Stream, out *stream.Stream) { out.PushBack(in.Pull()) })
        in, out := worker.Start()
        in.PushBack(expect)
        return out.Pull()
    })
}
```
