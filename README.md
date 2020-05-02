## err

```go
import "github.com/qq1060656096/err"

err := err.NewErr("test", 500100100)
err.SetLocation()

webErr := err.NewWebErr(503, "http 503", 500100100)
webErr.SetLocation()
```