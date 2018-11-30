package ravendb

func IndexCreation_createIndexes(indexes []*AbstractIndexCreationTask, store *IDocumentStore, conventions *DocumentConventions) error {

	if conventions == nil {
		conventions = store.GetConventions()
	}

	indexesToAdd := IndexCreation_createIndexesToAdd(indexes, conventions)
	op := NewPutIndexesOperation(indexesToAdd...)
	err := store.Maintenance().Send(op)
	if err == nil {
		return nil
	}

	// For old servers that don't have the new endpoint for executing multiple indexes
	for _, index := range indexes {
		err = index.Execute2(store, conventions, "")
		if err != nil {
			return err
		}
	}
	return nil
}

func IndexCreation_createIndexesToAdd(indexCreationTasks []*AbstractIndexCreationTask, conventions *DocumentConventions) []*IndexDefinition {
	var res []*IndexDefinition
	for _, x := range indexCreationTasks {
		x.SetConventions(conventions)
		definition := x.CreateIndexDefinition()
		definition.Name = x.GetIndexName()
		pri := x.GetPriority()
		if pri == "" {
			pri = IndexPriority_NORMAL
		}
		definition.Priority = pri
		res = append(res, definition)
	}
	return res
}
