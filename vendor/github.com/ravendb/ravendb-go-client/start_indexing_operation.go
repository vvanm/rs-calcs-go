package ravendb

import (
	"net/http"
)

var _ IVoidMaintenanceOperation = &StartIndexingOperation{}

type StartIndexingOperation struct {
	Command *StartIndexingCommand
}

func NewStartIndexingOperation() *StartIndexingOperation {
	return &StartIndexingOperation{}
}

func (o *StartIndexingOperation) GetCommand(conventions *DocumentConventions) RavenCommand {
	o.Command = NewStartIndexingCommand()
	return o.Command
}

var (
	_ RavenCommand = &StartIndexingCommand{}
)

type StartIndexingCommand struct {
	RavenCommandBase
}

func NewStartIndexingCommand() *StartIndexingCommand {
	cmd := &StartIndexingCommand{
		RavenCommandBase: NewRavenCommandBase(),
	}
	cmd.ResponseType = RavenCommandResponseType_EMPTY
	return cmd
}

func (c *StartIndexingCommand) CreateRequest(node *ServerNode) (*http.Request, error) {
	url := node.GetUrl() + "/databases/" + node.GetDatabase() + "/admin/indexes/start"

	return NewHttpPost(url, nil)
}
