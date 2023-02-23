package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Room struct {
	Password string `json:"password"`
	PaticNum int    `json:"particNum"`
}

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/myPage", hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する
	//curl -d "password = name" -d "particNum=10" http://localhost:1323/createRoom
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// ハンドラーを定義
func hello(c echo.Context) error {

	var fukami string = "Kensuke"
	return c.String(http.StatusOK, fukami)
}

//$body = @{
//     password = "mypassword"
//     particNum = 5
// } | ConvertTo-Json

// Invoke-RestMethod -Method POST -Uri http://localhost:1323/createRoom -Body $body -ContentType "application/json"
