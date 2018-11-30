package ravendb

type DeleteAttachmentCommandData struct {
	*CommandData
}

// NewDeleteAttachmentCommandData creates CommandData for Delete Attachment command
func NewDeleteAttachmentCommandData(documentId string, name string, changeVector *string) *DeleteAttachmentCommandData {
	res := &DeleteAttachmentCommandData{
		&CommandData{
			Type:         CommandType_DELETE,
			ID:           documentId,
			Name:         name,
			ChangeVector: changeVector,
		},
	}
	return res
}

func (d *DeleteAttachmentCommandData) serialize(conventions *DocumentConventions) (interface{}, error) {
	res := d.baseJSON()
	res["Type"] = "AttachmentDELETE"
	res["Name"] = d.Name
	return res, nil
}
