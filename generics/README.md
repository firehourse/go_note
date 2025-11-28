## generics

泛型是什麼? 

泛型實際上是一個偷懶的方法，當你在編譯的時候，泛型會幫你寫型別的檢查
主要還是因為go 是強型別的語言
所以當你在編譯期間的時候他就可以找到所有的呼叫點，幫你補上相應的型別
避免明明只是輸入型別不一樣，卻要寫好幾次代碼的狀況發生

```go
func compareValue[T constraints.Ordered](a, b T) T{
    if a > b {
        return a
    }
    return b
}
```
那除了any 你當然也可以做限制
```go
func compareValue[T int | float64](a, b T) T{
    if a > b {
        return a
    }
    return b
}
```
