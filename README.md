## err

```go
import "github.com/qq1060656096/err"

err := err.NewErr("test", 500100100)
err.SetLocation()

webErr := err.NewWebErr(503, "http 503", 500100101)
webErr.SetLocation()

webErr = err.CustomUnprocessableEntity("参数错误", 422100100, nil)
webErr.SetLocation()

webErr = err.CustomInternalServerError("数据库连接失败", 500100102, nil)
webErr.SetLocation()

```