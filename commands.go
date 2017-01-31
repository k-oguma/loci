//
// commands.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package main

import (
	"fmt"
	"os"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// GlobalFlags defines global flags.
var GlobalFlags = []cli.Flag{
	// TODO: Remote docker host.
	// TODO: no-colour option.
	// TODO: delete name option because lots of containers will be created.
	cli.StringFlag{
		Name: "name, n",
		Usage: "creating a container named `NAME` to run tests. " +
			"If name is given, container will not be deleted.",
	},
	cli.StringFlag{
		Name:  "tag, t",
		Usage: "specify a `TAG` name of the docker image to be build.",
	},
	cli.StringFlag{
		Name:  "base, b",
		Usage: "use image `TAG` as the base image.",
		Value: "ubuntu:latest",
	},
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "verbose mode, which prints Dockerfile and entrypoint.sh.",
	},
	cli.StringFlag{
		Name:   "apt-proxy",
		Usage:  "`URL` for a proxy server of apt repository.",
		EnvVar: "APT_PROXY",
	},
	cli.StringFlag{
		Name:   "pypi-proxy",
		Usage:  "`URL` for a proxy server of pypi repository.",
		EnvVar: "PYPI_PROXY",
	},
	cli.StringFlag{
		Name:   "http-proxy",
		Usage:  "`URL` for a http proxy server.",
		EnvVar: "HTTP_PROXY",
	},
	cli.StringFlag{
		Name:   "https-proxy",
		Usage:  "`URL` for a https proxy server.",
		EnvVar: "HTTPS_PROXY",
	},
	cli.StringFlag{
		Name:   "no-proxy",
		Usage:  "Comma separated URL `LIST` for which proxies won't be used.",
		EnvVar: "NO_PROXY",
	},
	cli.BoolFlag{
		Name:  "no-build-cache",
		Usage: "Do not use cache when building the image",
	},
}

// Commands defines sub-commands.
var Commands = []cli.Command{}

// CommandNotFound prints an error message when a given command is not supported.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(
		os.Stderr, chalk.Red.Color("%s: '%s' is not a %s command. See '%s --help'."),
		c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
