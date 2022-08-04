package cli_test

import (
	cli "github.com/eneskzlcn/incli"
	"github.com/eneskzlcn/incli/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAppArgsThenItShouldReturnAppWhenNewAppCalled(t *testing.T) {
	appName := "wordict"
	app := cli.NewApp(nil, nil, appName)
	assert.NotNil(t, app)
	assert.Equal(t, app.GetName(), appName)
}
func TestGivenRootHandlerThenItShouldBeAddedToAppWhenAddRootHandlerCalled(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRootHandler := mocks.NewMockRootHandler(mockController)
	rootHandlerName := "test"
	mockRootHandler.EXPECT().GetName().Return(rootHandlerName).Times(1)
	app := cli.NewApp(nil, nil, "")
	app.AddRootHandler(mockRootHandler)
	assert.Equal(t, app.GetRootHandler().GetName(), rootHandlerName)
}
func TestGivenHandlerThenItShouldBeAddedToAppWhenAddHandlerCalled(t *testing.T) {
	mockController := gomock.NewController(t)
	mockHandler := mocks.NewMockHandler(mockController)
	handlerName := "test"
	mockRootHandler := mocks.NewMockRootHandler(mockController)
	mockRootHandler.EXPECT().AddSubHandler(mockHandler).Return().Times(1)
	mockHandler.EXPECT().GetName().Return(handlerName).Times(1)
	app := cli.NewApp(nil, mockRootHandler, "")
	app.AddHandler(mockHandler)
	assert.Equal(t, true, app.HasHandler(handlerName))
}
