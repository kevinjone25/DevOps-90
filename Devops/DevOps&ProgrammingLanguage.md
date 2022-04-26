# DevOps跟程式語言的關西

## 這篇大概在講要用甚麼樣的語言 本來想跳過 但後來發現蠻有趣的 簡單來說是Go的推坑?


## 搞清楚一些東西

- 作者推薦是用Python或Go，因為大多數工具是用這兩種語言編寫
- 他在的Veeam by Kasten是一家雲原生公司，數據管理都是用Go
- 配置文件的YAML也很重要

## Why Go?

- 除卻他很熱門
- K8S, Docker(甚至要取代他的Podman), Grafana, Prometheus都是用Go編寫的

## Go的優勢

- python雖然不用compile
- 但Go的編譯速度夠快(轉成machine code)
- Go 是 statically linked
- 所有內容都在一個二進制可執行檔當中 所以不需要再裝dependency在remote端(相較python而言要裝一堆)
- 因為上一點所以可以跨平台，Python要為特定OS建binary file不容易
- CPU和記憶體使用率低，而且包含許多優化
- 與Python比，Go的Standard Library包含了大多DevOps的服務(HTTP Web,JSON EXECUTE......)
- 最後他並沒有把Python拋棄，他說他只是給出選Go的理由(還有因為他公司是用Go)
