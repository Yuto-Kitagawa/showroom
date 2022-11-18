# SHOWROOM API 作成手順
### 実行環境
- Windows 10

# 1.Golangのインストール
> [公式ページ](https://go.dev/dl/)からダウンロード。  
ダウンロードしたファイルを実行してGoをダウンロード・インストールを完了してください。 (今回はWindows版をインストールしました。また、インストールもややこしくなかったです)

<br> 

ダウンロード後、コマンドプロンプトでインストールが成功しているかを確認してください。

```
C:> go version
```
(結果)
<br>

> go version go1.19.3 windows/amd64

<br>

![](./src/PATH.png)
>環境変数にGOPATHが生成される。


# 2.サーバー用ファイルの準備
ファイルをダウンロードしてください。  
[ダウンロード](https://github.com/Chouette2100/t999srapi/archive/refs/tags/v0.1.0.zip)  
(ダウンロードしたファイル一式(t999srapi.goがあるディレクトリ)を、任意のフォルダに移してください。)
<br>

    
# 3.サーバーを起動
## 1. まずはコマンドプロンプトを開きます。
windowsマークから、コマンドプロンプトを開いてください。<br>
ディレクトリを移動するんですが、t999srapi.goがあるフォルダーのパスをコピーをして移動してください。<br>
例)  
```
cd C:\Users\user\Desktop\PROGRAMS\Golang\showroom
dir
```
(結果)
> public  
> template  
> \-\-\-  
> t999srapi.go  

***Goファイルが表示されたか確認してください。***

## 2. 以下のコマンドを実行します
```
    go mod init main
    go mod tidy
    go build t999srapi.go
    go run t999srapi.go
```
- ブラウザでlocalhost:8080/t009topを入力するとルームのリストが表示されます。
- フォルダに.exeファイルが作成されているので、それを起動してもサーバーが立ち上がります。