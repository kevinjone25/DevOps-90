# GO的Hello world解釋

---

## Compile

- 在撰寫高階語言像是C++,Python,Go之類的時候，我們是撰寫人能夠看的懂的語言，但是機器在編譯的時候總不可能直接去看我們寫的，它要轉換機器能夠看的懂的形式，而這個過程我們稱之為編譯

![](https://i.imgur.com/fB8Ct3H.png)

## package

- 有寫過C或C++的話 這跟他的library是一樣的概念
    - ex: <stdio.h>
- 其實就是一堆的.go檔案
- 可以直接引用已經寫好的.go 檔案 不用自己刻

```go=
package main

import 'fmt' //另一個可以把package引入的地方

func main(){
    fmt.Println("DevOps90")   
}
```

- 第一行是指會用到main 的package 因為他是整隻程式的entry point

- import 'fmt' 是因為下面的Println是fmt package裡的東西
