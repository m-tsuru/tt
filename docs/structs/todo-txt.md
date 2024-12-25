## `todo.txt` について

[todo.txt](https://github.com/todotxt/todo.txt) は，標準的なファイルフォーマットです．

ヒューマンリーダブルなテキストファイルによって提供されます．

一つのタスクは一つの行で構成され，以下のような形式で提供されることは，公式のドキュメントで示されます．

![](https://raw.githubusercontent.com/todotxt/todo.txt/refs/heads/master/description.svg)

```txt
x (A) 2016-05-20 2016-04-20 measure space for +chapelShelving @chapel due:2016-05-30
```

Description タグでは，tags を自由に登録することができます．tt では，メタデータをこの場所に配置します．

例えば，例のようにです．

```txt
x (A) 2024-12-25 2024-12-21 クリスマスケーキを買う +クリスマスパーティ2024 tt-creation-time:2024-12-21T17:37:00 tt-due-time:2024-12-25T18:00:00 tt-done-time:2024-12-25T17:30:00 tt-id:01JFZ53FV9KJRDQ2ASB2MD84X8
```

## 拡張タグ Key-Value

|key|description|example|
|---|---|---|
|tt-creation-time|タスクを作成した時間|2024-12-21T17:37:00|
|tt-done-time|タスクを完了させた時間|2024-12-21T17:30:00|
|tt-due-time|タスクを終了させる時間|2024-12-21T17:37:00|
|tt-id|タスクの固有ID|01JFZ53FV9KJRDQ2ASB2MD84X8|
|tt-parent|タスクの親ID|01JFZ53FV9KJRDQ2ASB2MD84X8|
