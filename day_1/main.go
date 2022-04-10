package main

import (
	"os"
	"os/exec"

	"git.wndv.co/go/srv/errorx"
	"git.wndv.co/go/srv/gin"
	"git.wndv.co/go/srv/grace"
)

type GoRunReq struct {
	Source string `json:"source"`
}

type GoRunRes struct {
	Output string
}

func main() {
	r := gin.Default()
	r.Configurer().EnableHealthz()

	r.GETw("/ping", func(ctx *gin.Context) (interface{}, error) {
		return gin.H{
			"message": "pong",
		}, nil
	})

	r.GETw("/hello/:name", func(ctx *gin.Context) (interface{}, error) {
		name := ctx.Param("name")
		if len(name) <= 3 {
			return nil, errorx.ErrBadRequest.Msg("Name is too short")
		}

		return gin.H{
			"message": "hello" + name,
		}, nil
	})

	r.POSTw("/go/run", func(ctx *gin.Context) (interface{}, error) {
		var req GoRunReq
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			return nil, errorx.ErrBadRequest.Msg("Error binding json.")
		}

		err2 := os.MkdirAll("./tmp", os.ModePerm)
		if err2 != nil {
			println("error " + err2.Error())
		}

		f, err3 := os.Create("./tmp/src.go")
		if err3 != nil {
			println("error " + err3.Error())
		}

		defer f.Close()

		source := req.Source
		_, err4 := f.WriteString(source)

		if err4 != nil {
			println("error " + err4.Error())
		}

		stdout, err5 := exec.Command("go", "run", "./tmp/src.go").CombinedOutput()
		var output string
		if err5 != nil {
			println(err5.Error())
			output = string(stdout) + " " + err5.Error()
		} else {
			output = string(stdout)
		}

		return GoRunRes{
			Output: output,
		}, nil
	})

	grace.OnSignal(func() {
		r.StopGracefully()
	})

	r.RunDefault()
}
