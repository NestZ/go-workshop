package main

import (
	"git.wndv.co/go/pillars"
	"git.wndv.co/go/srv/grace"
	"git.wndv.co/thaneat.s/app2/repos"
	"git.wndv.co/thaneat.s/app2/web"
)

func main() {
	internalSvr := pillars.NewInternalServer("email-campaign")
	go internalSvr.RunDefault()

	repositories := repos.Init()
	r := web.Init(repositories)

	grace.OnSignal(func() {
		r.StopGracefully()
		internalSvr.Stop()
	})

	r.RunDefault()
}
