package cmd

import (
	"fmt"

	"github.com/leaanthony/clir"
)

var cmd *clir.Cli

func customBanner(cli *clir.Cli) string {
	return ``
}

func Initialize() {
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
			subcmd, err := NewCommandStart(ConfigCommandStart{
				ConfigPath: cfgpath,
			})
			if err != nil {
				return
			}

			return subcmd.Run()
		})
	}
}

func Execute() {
	if err := cmd.Run(); err != nil {
		err = fmt.Errorf("unexpected error: %w", err)
		fmt.Println(err)
		return
	}
}
