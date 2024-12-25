## `.tt` について

`.tt` ディレクトリは，以下のように構成されます．

```
.
└── .tt/
    ├── config
    └── todo.txt
```

このページでは，`.tt` ディレクトリのファイルについて解説します．

### config

config はサブコマンド[`config`](../subcommands/config.md)によって指定される，tt に関わる設定が格納されます．

config には，通常次のような内容が含まれます．

```config
[core]
        mode = todo.txt
        refer = todo.txt
        tz = Asia/Tokyo
```

上の例でいう `mode` は，管理フォーマットを表します．現在は `todo.txt` 形式のみがサポートされていますが，将来的に `RDBMS` などに対応するかもしれません．

加えて，`refer` は，参照するファイルやデータベースの場所を指定します．デフォルト値は todo.txt ですが，場合によっては変更できます．

さらに，`~/.config/tt` にも，`config` は配置され，これはグローバルに適用される設定が格納されます．

多くの場合，次のような情報が格納されているでしょう．

```config
[user]
        name = Author_Name
        email = postmaster@example.com
        tz = Asia/Tokyo

[external]
        llm_completion = gemini
        llm_api_key = [llm_api_key]
```

通常の場合，この `config` ファイルをテキストエディタで編集することはありません．以下のコマンドを使用することができます．

```sh
tt config --global [<セクション>].[キー] [<値>]
```

例えば，タスクを追加するユーザの名前を次のように設定したい場合は，

```sh
tt config --global user.name Author_Name
```

などと設定します．詳しくはサブコマンド [`config`](../subcommands/config.md) を参照してください．

### todo.txt

- [`todo.txt` について]("./todo-txt.md")
