/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

*/

package main

import (
	//	"io" //　ログ出力設定用。必要に応じて。
	"log"
	"net/http"
	"net/http/cgi"
	"os"
	"time"

	"github.com/Chouette2100/srhandler"
)

/*

	SHOWROOMのAPIを利用したプログラム例です。
	このプログラムはSCRIPT_NAME が設定されていなければWebサーバーとして、設定されていればCGIとして起動されます。

	ソースのダウンロード、ビルドについて以下簡単に説明します。詳細は以下の記事を参考にしてください。
	WindowsもLinuxも特記した部分以外は同じです。

		【Windows】かんたんなWebサーバーの作り方
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/c5cab5

		---------------------

		【Windows】Githubにあるサンプルプログラムの実行方法
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/e27fc9

		【Unix/Linux】Githubにあるサンプルプログラムの実行方法
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/220e38

			ロードモジュールさえできればいいということでしたらコマンド一つでできます。

以下にビルドし、Webサーバーとして起動する例を示します。CGIとして使う場合はできあがったロードモジュールを
ディレクトリごと（Webサーバーの設定に応じて）所定のディレクトリーに配置してください。

【Unix/Linux】

	$ cd ~/go/src
	$ curl -OL https://github.com/Chouette2100/t999srapi/archive/refs/tags/v0.1.0.tar.gz
	$ tar xvf v0.1.0.tar.gz
	$ mv t999srapi-0.1.0 t999srapi
	$ cd t999srapi
	$ go mod init
	$ go mod tidy
	$ go build t999srapi.go
	$ ./t999srapi

	ここでブラウザを起動し

	http://localhost:8080/.......

	のようにURLを入力する。

	上の ..... の部分については以下のプログラムのハンドラー定義のところに説明があります。

	例えば

	http://localhost:8080/apieventroomlist

	のようにURLを入力すると開催中イベントの参加ルーム一覧（のイベント選択ページ）が表示されます。


	【Windows】

	Microsoft Windows [Version 10.0.22000.856]
	(c) Microsoft Corporation. All rights reserved.

	C:\Users\chouette>cd go

	C:\Users\chouette\go>cd src

	作業はかならず %HOMEPATH%\go\src の下で行います。

	以下、要するに https://github.com/Chouette2100/t999srapi/releases にあるv0.1.0のZIPファイルSource code (zip) からソースをとりだしてくださいということなので、ブラウザでダウンロードしてエクスプローラで解凍というこでもけっこうです。なんならこの記事の最後にあるgithubのソースをエディターにコピペで作るということでもかまいません（この場合文字コードはかならずUTF-8にしてください 改行はLFになっています。というようなことを考えるとやっぱりダウンロードして解凍が安全かも）

	C:\Users\chouette\go\src>mkdir t999srapi

	C:\Users\chouette\go\src>cd t999srapi

	C:\Users\chouette\go\src\t999srapi>curl -OL https://github.com/Chouette2100/t999srapi/archive/refs/tags/v0.1.0.zip
	  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
	                                 Dload  Upload   Total   Spent    Left  Speed
	  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
	100  6265    0  6265    0     0   6777      0 --:--:-- --:--:-- --:--:-- 16400

	C:\Users\chouette\go\src\t999srapi>call powershell -command "Expand-Archive v0.1.0.zip"

	C:\Users\chouette\go\src\t999srapi>tree
	フォルダー パスの一覧
	ボリューム シリアル番号は E2CD-BDF1 です
	C:.
	└─v0.1.0
	    └─t999srapi-0.1.0
	        ├─public
	        └─templates

	C:\Users\chouette\go\src\t999srapi>xcopy /e v0.1.0\t999srapi-0.1.0\*.*
	v0.1.0\t999srapi-0.1.0\LICENSE
	v0.1.0\t999srapi-0.1.0\README.md
	v0.1.0\t999srapi-0.1.0\t999srapi.go
	v0.1.0\t999srapi-0.1.0\public\index.html
	v0.1.0\t999srapi-0.1.0\templates\apieventroomlist.gtpl
	以下省略（リリースによって内容が異なる場合があります。）

	C:\Users\chouette\go\src\t999srapi>rmdir /s /q v0.1.0

	C:\Users\chouette\go\src\t999srapi>del v0.1.0.zip

	ここで次のような構成になっていればOKです。top.gtpl と index.html が所定の場所にあることをかならず確かめてください。

	C:%HOMEPATH%\go\src\t999srapi --+-- t999srapi.go 他設定ファイル、シェルスクリプト等
	                                |
	                                +-- \templates --- apieventroomlist.gtpl 他.gtplファイル
	                                |
	                                +-- \public    --- index.html

	ここからはコマンド三つでビルドが完了します。

	C:\Users\chouette\go\src\t999srapi>go mod init
	go: creating new go.mod: module t999srapi
	go: to add module requirements and sums:
	        go mod tidy

	C:\Users\chouette\go\src\t999srapi>go mod tidy
	go: finding module for package github.com/dustin/go-humanize
	go: downloading github.com/dustin/go-humanize v1.0.0
	go: found github.com/dustin/go-humanize in github.com/dustin/go-humanize v1.0.0

	C:\Users\chouette\go\src\t999srapi>go build t999srapi.go

	あとは

	C:\Users\chouette\go\src\t999srapi>t999srapi

	でWebサーバが起動します。ここでセキュリティー上の警告が出ると思いますが、説明をよく読んで問題ないと思ったらアクセスを許可してください（もちろん許可しなければWebサーバは使えなくなります）

	ここでブラウザを開き

	http://localhost:8080/.......

	のようにURLを入力する。

	上の ..... の部分については以下のプログラムのハンドラー定義のところに説明があります。

	例えば

	http://localhost:8080/apieventroomlist

	のようにURLを入力すると開催中イベントの参加ルーム一覧（のイベント選択ページ）が表示されます。


	Ver. 0.1.0

*/



//Webサーバーを起動する。
func main() {

	logfilename := time.Now().Format("20060102") + ".txt"
	logfile, err := os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open logfile: " + logfilename + err.Error())
	}
	defer logfile.Close()

	/*
		//	ログをコンソールにも出力する。原則これは行わないこと。とくにCGIの場合はぜったいダメ。
		//	"io"をimportすること。
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	*/

	//	ログ出力を設定する。
	log.SetOutput(logfile)

	//	環境変数 SCRIPT_NAME を取得しrootPathとする。この環境変数が設定されていればCGIとして起動されている。
	//	SCRIPT_NAME が設定されていなければWebサーバーとして起動されている。
	rootPath := os.Getenv("SCRIPT_NAME")
	log.Printf("rootPath: \"%s\"\n", rootPath)

	//	URLに対するハンドラーを定義する。

	//	開催中イベント一覧表を表示する。
	//	http://localhost:8080/t008top で呼び出される。
	http.HandleFunc(rootPath+"/t008top", srhandler.HandlerT008topForm)

	//	配信中ルーム一覧を表示する。
	//	http://localhost:8080/t009top で呼び出される。
	http.HandleFunc(rootPath+"/t009top", srhandler.HandlerT009topForm)

	//	ルーム状況を表示する。
	//	http://localhost:8080/apiroomstatus?room_url_key=... で呼び出される。
	http.HandleFunc(rootPath+"/apiroomstatus", srhandler.HandlerApiRoomStatus)

	//	イベント参加ルーム一覧を表示する。
	//	http://localhost:8080/apieventroomlist?evenid=... で呼び出される。
	//	eventid がわからないときは「イベント一覧表」にある「参加者一覧へのリンク」を利用します。
	http.HandleFunc(rootPath+"/apieventroomlist", srhandler.HandlerApiEventRoomList)

	//	ポートは8080などを使います。
	//	Webサーバーはroot権限のない（特権昇格ができない）ユーザーで起動した方が安全だと思います。
	//	その場合80や443のポートはlistenできないので、ルータやOSの設定でポートフォワーディングするか
	//	ケーパビリティを設定してください。
	//	# setcap cap_net_bind_service=+ep ShowroomCGI
	//　（設置したあとこの操作を行うこと）
	httpport := "8080"

	//		CGIとして起動されたときはWebサーバーやCGIの設置場所にあわせて変更すること。
	//		さくらのレンタルサーバーでwwwの直下にCGIを設置したときはこのままでいいです。

	if rootPath == "" {
		//	Webサーバーとして起動された
		//		URLがホスト名だけのときは public/index.htmlが表示されるようにしておきます。
		//		public ディレクトリーはグラフ画像の作成場所としても使っています。
		http.Handle("/", http.FileServer(http.Dir("public"))) // http://localhost:8080/ で呼び出される。
		err = http.ListenAndServe(":"+httpport, nil)
	} else {
		//	cgiとして起動された
		err = cgi.Serve(nil)
	}
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}
