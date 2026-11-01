# Purin

> **開発終了（2026年11月）**：このリポジトリは今後開発しません。

Pudding in the fridge.

## やりたかったこと

ソースコードを取ってきて、ビルドして、自分のホームフォルダにインストールするパッケージマネージャー（Homebrew と同じ種類のツール）。名前は、Homebrew がビール醸造にちなんだ用語を使っているのに対して「冷蔵庫の中のプリン」。

構想していた使い方：

    # tar.gz を取ってきて、コンパイル時のフラグを指定してビルド
    purin get http://foo.tar.gz --flag COMPILE_FLAG1=bar --flag COMPILE_FLAG2=egg
    # GitHub のリポジトリ（どちらも同じ）
    purin get tetor/some
    purin get github.com/tetor/some.git
    # 削除
    purin delete github.com/tetor/some.git

インストール先は `$HOME/purin/bin`。

- 言語：Go（コマンド部分は cobra、git 操作は go-git）
- 期間：2017年5月〜8月
- 実装は `get` の作りかけまで。固定のリポジトリを clone して `./configure` → `make` → `make install` を実行するところで止まっており、引数はまだ受け付けない

## 終了理由

同じ役割のツールが十分に揃っており、自分で作る必要がなくなったため。

- macOS・Linux のパッケージ管理は Homebrew が定番になっている
- ビルドの再現性まで求めるなら Nix がある
- 言語やツールのバージョン管理は asdf や mise で足りる

今から作る価値は薄いと判断し、開発を終了する。
