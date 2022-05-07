# Go DataType & Variable


## 宣告變數

- 在Go裡可以用var去做宣告
- 但也可以用更clear的放是作宣告

```go=
var challenge = "90DaysOfDevOps" # origin method

challenge := "90DaysOfDevOps" # more clear method
```

## 資料型態

- 如果沒有給型態 GO會幫你自動辯識型態

- 有四種型態
	1. Basic type: Numbers, strings booleans
	2. Aggregate type: Array, Structs
	3. Reference: Pointers, slices, maps, functions, channels
	4. Interface type
- 在定義型態如下列這樣
```go=
var TwitterComment string
var DaysCompleted uint
```

- 如果要印出變數 可以像這樣
```go=
fmt.Printf("number is %T, daystotal is %T\n", number, daysComplete)
```

- 在int和float型態中,有更多的detail
	- int : 數字
	- unint : 正整數
	- floating point types: 有浮點數



