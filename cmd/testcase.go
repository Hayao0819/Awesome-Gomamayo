package cmd

import (
	"github.com/Hayao0819/awesome-gomamayo/cmd/testcase"
	"github.com/Hayao0819/nahi/cobrautils"
)

func init() {
	cobrautils.AddSubCmds(testcase.Cmd())
}
