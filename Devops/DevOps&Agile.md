# Devops 和 Agile 的差別

## Agile Development

- 高速release version(一個禮拜、月)，週期短
- 以提供終端使用者最好的體驗為目標

## DevOps

- Operation和Developer之間的合作
- 大幅簡化開發過程中溝通不順的問題

### 差別的部分

- 整體來看他們是互補的
- Agile需要快速的更替，而DevOps剛好可以補足
- Agile希望客戶能嘗試新版本並馬上給予意見，這只有在DevOps讓環境創建變得容易才有可能

## 框架

- Agile有很多框架來進行實現Scrum > Kanban > Lean > Extreme > Crystal > Dynamic > Feature-Driven
- DevOps則是不提供特定方法，專注於協作，只要能夠實踐就好


###  focus在哪?

- Agile focus 在 使用者和開發人員間的溝通 => 客戶導向
- DevOps則是開發人員和維運團隊之間的溝通 => 內部實踐     

### 誰的意見重要

- Agile => 使用者
- DevOps => stakeholders(權益人)和團隊

### 關注的點

- Agile 注重 software 開發
- DevOps也注重software 開發，但部署和監控也算在內

### Docs

- Agile 較為注重 靈活性和task
- DevOps 則是將Docs 視為必要的項目之一

### 風險

- Agile 的風險就是因為過於靈活 很難去預設和評估
- DevOps風險則是對專業名詞的誤解和缺乏適用的工具，有些人就認為DevOps是Deployment和CI的software

### 使用的工具

- Agile
    - JIRA、Trello、Slack、Zoom、SurveyMonkey 
- DevOps
    - Jenkins、GitHub Actions、Circle CI 之類的咚咚

## 整合Agile 和 DevOps

- Agile 可以幫助DevOps團隊有效地傳遞他們的訊息
- 高強度的管理和強力的工具
- 採用Agile 可以增進合作增加團隊動力和減少員工離職的願望
- product quality 提升

#### Agile 允許產品rollback回去開發階段來修正錯誤 防止技術債的累積

## 採用Agile和DOevOps的7個Step

1.  統整開發(Dev)和維運(Ops)團隊
2.  創立build和run的team，所有開發和維運的問題由DevOps團隊討論
3.  將等價值的開發和DevOps task分配優先度
4.  把QA納入Dev的環節
5.  選擇對的工具
6.  自動化所有你能力所及的事物
7.  用數值化來進行評估和控制