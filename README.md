# mercari-gopherdojo-03

## 概要

おみくじAPI

## build

```
go mod tidy
go build
```

## usage

```
./omikuji [port] &
curl localhost:[port]
```

終了時は
```
killall omikuji
```
又はpsコマンドでPIDを探してkillコマンドで終了

## 注釈

* https://omikuji-guide.com/number/ からデータをいただきました。

