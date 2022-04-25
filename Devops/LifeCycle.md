# DevOps 8字環

## 大概講解一下各個步驟 (應該說成DevOps在每個階段扮演的腳色)

## Plan

- 在開發團隊決定下個版本要發布哪些feature和bug fixes的過程中，DevOps Engineer可以去協助他們並影響他們的選擇
- 選擇是指往好的方向發展，指出會往壞方向發展的可能

## Code

- 因為基本上不會參與到開發 所以可以協助開發團隊了解Infra(Why?)
- A:這樣開發團隊可以了解到說如何更好地與這些API或者Service做溝通

## Build

- 第一個自動化的地方
- 根據code來build，一律透過CI/CD

## Testing

- build完後就要測試
- 透過輸入來減少生產中的錯誤

## Release

- 測試通過後透過git放到github

## Deploy

- 已經是最終環結
- 將code真正投入生產

## Operate

- 處理可能發生的事情
- 伺服器在巔峰時期需要做load baleance，在非高峰時期減少伺服器數目
- 反饋給開發團隊

## monitor

- 監控CPU 使用率、磁碟空間、api、反應時間、和log
- log要讓開發人員無須訪問生產系統就可以看的到!!!

## Repeat

- 重複這個cycle

## CD

CD = Plan + Code + Build + Test

## CI

CI = Plan + Code + Build + Test + Release
CI = CD + Release(不管有沒有成功)

## CI cont.

- 假如CI成功 會轉到下個階段
=> CD = Deploy > Operate > Monitor

- 上述是DevOps生命週期的簡單集合