package ravendb

import (
	"net/http"
)

var _ IVoidMaintenanceOperation = &DeleteIndexOperation{}

type DeleteIndexOperation struct {
	_indexName string

	Command *DeleteIndexCommand
}

func NewDeleteIndexOperation(indexName string) *DeleteIndexOperation {
	panicIf(indexName == "", "indexName cannot be empty")

	return &DeleteIndexOperation{
		_indexName: indexName,
	}
}

func (o *DeleteIndexOperation) GetCommand(conventions *DocumentConventions) RavenCommand {
	o.Command = NewDeleteIndexCommand(o._indexName)
	return o.Command
}

var (
	_ RavenCommand = &DeleteIndexCommand{}
)

type DeleteIndexCommand struct {
	RavenCommandBase

	_indexName string
}

func NewDeleteIndexCommand(indexName string) *DeleteIndexCommand {
	panicIf(indexName == "", "indexName cannot be empty")
	cmd := &DeleteIndexCommand{
		RavenCommandBase: NewRavenCommandBase(),

		_indexName: indexName,
	}
	cmd.ResponseType = RavenCommandResponseType_EMPTY
	return cmd
}

func (c *DeleteIndexCommand) CreateRequest(node *ServerNode) (*http.Request, error) {
	url := node.GetUrl() + "/databases/" + node.GetDatabase() + "/indexes?name=" + UrlUtils_escapeDataString(c._indexName)

	return NewHttpDelete(url, nil)
}
