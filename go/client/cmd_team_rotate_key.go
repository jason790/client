// Copyright 2017 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"context"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

type CmdTeamRotateKey struct {
	libkb.Contextified
	TeamID keybase1.TeamID
}

func newCmdTeamRotateKey(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name:         "team-rotate-key",
		ArgumentHelp: "<team id>",
		Usage:        "Rotate a team's key",
		Action: func(c *cli.Context) {
			cmd := NewCmdTeamRotateKeyRunner(g)
			cl.ChooseCommand(cmd, "team-rotate-key", c)
		},
		Flags: []cli.Flag{},
	}
}

func NewCmdTeamRotateKeyRunner(g *libkb.GlobalContext) *CmdTeamRotateKey {
	return &CmdTeamRotateKey{Contextified: libkb.NewContextified(g)}
}

func (c *CmdTeamRotateKey) ParseArgv(ctx *cli.Context) error {
	var err error
	c.TeamID, err = ParseOneTeamID(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *CmdTeamRotateKey) Run() error {
	cli, err := GetTeamsClient(c.G())
	if err != nil {
		return err
	}

	return cli.TeamRotateKey(context.Background(), c.TeamID)
}

func (c *CmdTeamRotateKey) GetUsage() libkb.Usage {
	return libkb.Usage{
		Config:    true,
		API:       true,
		KbKeyring: true,
	}
}
