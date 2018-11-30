package ravendb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

var _ ILazyOperation = &LazyStartsWithOperation{}

type LazyStartsWithOperation struct {
	_clazz             reflect.Type
	_idPrefix          string
	_matches           string
	_exclude           string
	_start             int
	_pageSize          int
	_sessionOperations *InMemoryDocumentSessionOperations
	_startAfter        string

	result        Object
	queryResult   *QueryResult
	requiresRetry bool
}

func NewLazyStartsWithOperation(clazz reflect.Type, idPrefix string, matches string, exclude string, start int, pageSize int, sessionOperations *InMemoryDocumentSessionOperations, startAfter string) *LazyStartsWithOperation {
	return &LazyStartsWithOperation{
		_clazz:             clazz,
		_idPrefix:          idPrefix,
		_matches:           matches,
		_exclude:           exclude,
		_start:             start,
		_pageSize:          pageSize,
		_sessionOperations: sessionOperations,
		_startAfter:        startAfter,
	}
}

func (o *LazyStartsWithOperation) createRequest() *GetRequest {
	q := fmt.Sprintf("?startsWith=%s&matches=%s&exclude=%s&start=%d&pageSize=%d&startAfter=%s",
		UrlUtils_escapeDataString(o._idPrefix),
		UrlUtils_escapeDataString(o._matches),
		UrlUtils_escapeDataString(o._exclude),
		o._start,
		o._pageSize,
		o._startAfter)

	request := &GetRequest{
		url:   "/docs",
		query: q,
	}

	return request
}

func (o *LazyStartsWithOperation) getResult() Object {
	return o.result
}

func (o *LazyStartsWithOperation) setResult(result Object) {
	o.result = result
}

func (o *LazyStartsWithOperation) getQueryResult() *QueryResult {
	return o.queryResult
}

func (o *LazyStartsWithOperation) setQueryResult(queryResult *QueryResult) {
	o.queryResult = queryResult
}

func (o *LazyStartsWithOperation) isRequiresRetry() bool {
	return o.requiresRetry
}

func (o *LazyStartsWithOperation) setRequiresRetry(requiresRetry bool) {
	o.requiresRetry = requiresRetry
}

func (o *LazyStartsWithOperation) handleResponse(response *GetResponse) error {
	var getDocumentResult *GetDocumentsResult
	err := json.Unmarshal([]byte(response.result), &getDocumentResult)
	if err != nil {
		return err
	}

	finalResults := map[string]Object{}
	//TreeMap<string, Object> finalResults = new TreeMap<>(string::compareToIgnoreCase);

	for _, document := range getDocumentResult.GetResults() {
		newDocumentInfo := DocumentInfo_getNewDocumentInfo(document)
		o._sessionOperations.documentsById.add(newDocumentInfo)

		if newDocumentInfo.id == "" {
			continue // is this possible?
		}

		id := strings.ToLower(newDocumentInfo.id)
		if o._sessionOperations.IsDeleted(newDocumentInfo.id) {
			finalResults[id] = nil
			continue
		}
		doc := o._sessionOperations.documentsById.getValue(newDocumentInfo.id)
		if doc != nil {
			finalResults[id], err = o._sessionOperations.TrackEntityInDocumentInfoOld(o._clazz, doc)
			if err != nil {
				return err
			}
			continue
		}
		finalResults[id] = nil
	}
	o.result = finalResults
	return nil
}
