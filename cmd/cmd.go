package cmd

import (
	"github.com/leaanthony/clir"
)

var cmd *clir.Cli

func customBanner(cli *clir.Cli) string {
	return ``
}

func init() {
	cmd = clir.NewCli("minimok", "mini mock server", "v1")
	// cmd.SetBannerFunction(customBanner)

	// default action
	cmd.Action(func() (err error) {
		cmd.PrintHelp()
		return
	})

	cmdStart := cmd.NewSubCommand("start", "start minimok")
	{
		cfgpath := ""
		cmdStart.StringFlag("conf", "yaml config file path", &cfgpath)
		cmdStart.Action(func() (err error) {
			return runStart(configStart{
				ConfigPath: cfgpath,
			})
		})
	}
}

func Execute() error {
	return cmd.Run()
}
