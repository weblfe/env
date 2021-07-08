## Env Utils 
读取 env value 转换成 期望类型

## 示例

```go
import ("github.com/weblfe/env")

func main()  {
    var env = env.GetEnv()
    fmt.Println(env.GetOf("test","weblfe"))
    fmt.Println(env.GetIntOf("num",1))
}
```
