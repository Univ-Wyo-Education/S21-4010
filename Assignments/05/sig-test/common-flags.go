package main

// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// List of all of the flags that are commonly used in this program

import (
	"github.com/urfave/cli"
)

var passphraseFlag = cli.StringFlag{
	Name:  "passwordfile",
	Usage: "the file that contains the passphrase for the keyfile",
}

var jsonFlag = cli.BoolFlag{
	Name:  "json",
	Usage: "output JSON instead of human-text format",
}

var newPassphraseFlag = cli.StringFlag{
	Name:  "newpasswordfile",
	Usage: "the file that contains the new passphrase for the keyfile",
}

var debugFlag = cli.StringFlag{
	Name:  "debug",
	Usage: "turn on debug flags",
}

var msgfileFlag = cli.StringFlag{
	Name:  "msgfile",
	Usage: "file containing the message to sign/verify",
}

var defaultNameFlag = cli.BoolFlag{
	Name:  "default-name",
	Usage: "Use a default name for the output file, `DateTime--AccountNo`.",
}

var genMsgFlag = cli.BoolFlag{
	Name:  "gen-msg",
	Usage: "randomly generate a message to sign",
}

var randomPassFlag = cli.BoolFlag{
	Name:  "random-pass",
	Usage: "randomly generate passwords",
}

var logFileFlag = cli.StringFlag{
	Name:  "log-file",
	Usage: "file where info will be logged",
}

var newNameFlag = cli.StringFlag{
	Name:  "newname",
	Usage: "new name of keyfile, set to '-' to overwrite the existing file.",
}
