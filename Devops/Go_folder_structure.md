# GO folder structure

---

## src

- 這裡是放你GO程式和Project的地方
- 就是放你Hello資料夾的地方

## pkg

- 放壓縮package的地方
- package會幫助你加速編譯 

## bin

- 放build出來的binary檔案(就是.exe檔)
- 注意這個資料夾不能放進git
	- 原因是exe屬於結果 只要你原本source code有改 exe也要跟著改
	- git tree 會需要一直去計算hash值 然後變得超大

## 比較複雜的架構

- 原文章裡有提供比較複雜的架構




## 編譯和執行

- 在Go裡有三種編譯和執行的方法

```bash

go run
# 會編譯和執行你加在go run後面的.go file


go build
# 編譯package和他的dependencies
# 如果是main package 它會把.exe放在目前目錄底下
# 如果不是 會放在pkg folder
# go build也支援編譯exe在GO 支援的作業系統上

go install
# 跟go build 一樣 不過它生成完會放在/bin的目錄底下


```



