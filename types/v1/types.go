package goisilon

type ErrObjects struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Errors struct {
	Errors []ErrObjects `json:"errors"`
}

type DirChildren struct {
	Name string `json:"name"`
}

type Directory struct {
	Children []DirChildren `json:"children"`
}

type Ownership struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Acl struct {
	Trustee      Ownership `json:"trustee"`
	Accesstype   string    `json:"accesstype"`
	Accessrights string    `json:"accessrights"`
	InheritFlags string    `json:"inherit_flags"`
	Op           string    `json:"op"`
}

type AclRequest struct {
	Authoritative string    `json:"authoritative"` // Mandatory
	Action        string    `json:"action"`
	Owner         Ownership `json:"owner"`
	Group         Ownership `json:"group"`
	Mode          string    `json:"mode"`
	Acl           []Acl     `json:"acl"`
}

type SnapshotCreate struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Alias   string `json:"alias"`
	Expires int32  `json:"expires"`
}

type SnapshotID struct {
	Id string `json:"id"`
}
