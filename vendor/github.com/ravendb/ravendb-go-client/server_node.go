package ravendb

const (
	// ServerNode.Role
	ServerNode_Role_NONE       = "None"
	ServerNode_Role_PROMOTABLE = "Promotable"
	ServerNode_Role_MEMBER     = "Member"
	ServerNode_Role_REHAB      = "Rehab"
)

// ServerNode describes a single server node
type ServerNode struct {
	URL        string `json:"Url"`
	Database   string `json:"Database"`
	ClusterTag string `json:"ClusterTag"`
	ServerRole string `json:"ServerRole"`
}

// NewServerNode creates a new ServerNode
func NewServerNode() *ServerNode {
	return &ServerNode{
		ServerRole: ServerNode_Role_NONE,
	}
}

func (n *ServerNode) GetUrl() string {
	return n.URL
}

func (n *ServerNode) GetDatabase() string {
	return n.Database
}

func (n *ServerNode) GetServerRole() string {
	return n.ServerRole
}

func (n *ServerNode) GetClusterTag() string {
	return n.ClusterTag
}

func (n *ServerNode) SetDatabase(database string) {
	n.Database = database
}

func (n *ServerNode) SetUrl(url string) {
	n.URL = url
}

func (n *ServerNode) SetClusterTag(clusterTag string) {
	n.ClusterTag = clusterTag
}
