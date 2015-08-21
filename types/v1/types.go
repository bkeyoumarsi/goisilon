package goisilon

type Timestamp int32

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

type AclUpdateReq struct {
	Authoritative string    `json:"authoritative"` // Mandatory
	Action        string    `json:"action"`
	Owner         Ownership `json:"owner"`
	Group         Ownership `json:"group"`
	Mode          string    `json:"mode"`
	Acl           []Acl     `json:"acl"`
}

type SnapshotCreateReq struct {
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	Alias   string    `json:"alias"`
	Expires Timestamp `json:"expires"`
}

type Snapshot struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Schedule        string    `json:"schedule"`
	Created         Timestamp `json:"created"`
	Expires         Timestamp `json:"expires"`
	Path            string    `json:"path"`
	Size            int       `json:"size"`
	ShadowBlocks    int       `json:"shadow_blocks"`
	PctFilesystem   float32   `json:"pct_filesystem"`
	PctReserve      float32   `json:"pct_reserve"`
	AliasTarget     int       `json:"alias_target"`
	AliasTargetName string    `json:"alias_target_name"`
	HasLocks        bool      `json:"has_locks"`
	State           string    `json:"state"`
}

type SnapshotList struct {
	Snapshots []Snapshot `json:"snapshots"`
	Total     int        `json:"total"`
	resume    string     // Don't really care
}
