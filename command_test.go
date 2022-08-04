package cli_test

import (
	"testing"

	cli "github.com/eneskzlcn/incli"
	"github.com/stretchr/testify/assert"
)

func TestGivenCommandArgsThenItShouldReturnCommandWithArgsWhenNewCommandCalled(t *testing.T) {
	name := "testname"
	shortDesc := "test short description"
	longDesc := "test long description"
	subCommandName := "subcommand"
	flagName := "testflag"
	subCommands := []*cli.Command{
		cli.NewCommand(subCommandName, "", "", nil, func(c *cli.Command, s []string) {}, nil),
	}
	handler := func(command *cli.Command, args []string) {}
	flags := []*cli.Flag{
		cli.NewBoolFlag(true, flagName, "", "", "", false),
	}
	command := cli.NewCommand(name, shortDesc, longDesc, subCommands, handler, flags)
	assert.NotNil(t, command)
	assert.Equal(t, command.GetName(), name)
	assert.Equal(t, command.GetShort(), shortDesc)
	assert.Equal(t, command.GetLong(), longDesc)
	assert.Equal(t, command.HasSubCommand(subCommandName), true)
	assert.Equal(t, command.HasFlag(flagName), true)
}
