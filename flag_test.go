package cli_test

import (
	"strconv"
	"testing"

	cli "github.com/eneskzlcn/incli"
	"github.com/stretchr/testify/assert"
)

func TestFlags(t *testing.T) {
	boolFlagValue := true
	stringFlagValue := "testvalue"
	intFlagValue := 3
	flagName := "testname"
	flagShorthand := "t"
	flagUsage := "Test usage."
	flagNoOptionDefaultValue := "defaultTest"
	flagIsNoOption := false

	t.Run("given string formatted bool value then it should set bool flags value when Set method called", func(t *testing.T) {
		var flag cli.BoolFlag
		testValue := "true"
		flag.Set(testValue)
		assert.Equal(t, flag.String(), testValue)
	})

	t.Run("given bool flag arguments then it should create bool flag with given arguments when NewBoolFlag called ", func(t *testing.T) {
		flag := cli.NewBoolFlag(boolFlagValue, flagName, flagShorthand, flagUsage, flagNoOptionDefaultValue, flagIsNoOption)

		assert.Equal(t, flag.Value.String(), strconv.FormatBool(boolFlagValue))
		assert.Equal(t, flag.Name, flagName)
		assert.Equal(t, flag.Shorthand, flagShorthand)
		assert.Equal(t, flag.NoOptionDefaultValue, flagNoOptionDefaultValue)
		assert.Equal(t, flag.IsNoOption, flagIsNoOption)
	})

	t.Run("given string flag arguments then it should create string flag with given arguments when NewStringFlag called ", func(t *testing.T) {

		flag := cli.NewStringFlag(stringFlagValue, flagName, flagShorthand, flagUsage, flagNoOptionDefaultValue, flagIsNoOption)

		assert.Equal(t, flag.Value.String(), stringFlagValue)
		assert.Equal(t, flag.Name, flagName)
		assert.Equal(t, flag.Shorthand, flagShorthand)
		assert.Equal(t, flag.NoOptionDefaultValue, flagNoOptionDefaultValue)
		assert.Equal(t, flag.IsNoOption, flagIsNoOption)
	})
	t.Run("given string formatted string value then it should set string flags value when Set method called", func(t *testing.T) {
		var flag cli.StringFlag
		testValue := "test"
		flag.Set(testValue)
		assert.Equal(t, flag.String(), testValue)
	})
	t.Run("given int flag arguments then it should create int flag with given arguments when NewIntFlag called ", func(t *testing.T) {

		flag := cli.NewIntFlag(intFlagValue, flagName, flagShorthand, flagUsage, flagNoOptionDefaultValue, flagIsNoOption)
		assert.Equal(t, flag.Value.String(), strconv.FormatInt(int64(intFlagValue), 10))
		assert.Equal(t, flag.Name, flagName)
		assert.Equal(t, flag.Shorthand, flagShorthand)
		assert.Equal(t, flag.NoOptionDefaultValue, flagNoOptionDefaultValue)
		assert.Equal(t, flag.IsNoOption, flagIsNoOption)
	})
	t.Run("given string formatted int value then it should set int flags value when Set method called", func(t *testing.T) {
		var flag cli.IntFlag
		testValue := "3"
		flag.Set(testValue)
		assert.Equal(t, flag.String(), testValue)
	})
}
