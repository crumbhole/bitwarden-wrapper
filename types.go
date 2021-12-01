package bitwardenwrapper

// BwFolderID is an opaque ID for a bitwarden folder
type BwFolderID string

// BwCollectionID is an opaque ID for a bitwarden collection
type BwCollectionID string

// BwFolder is a go structure mapping to bitwarden's JSON API for a folder in bitwarden
type BwFolder struct {
	Object string     `json:"object"`
	ID     BwFolderID `json:"id"`
	Name   string     `json:"name"`
}

// BwURI is a go structure mapping to bitwarden's JSON API for a URI in bitwarden
type BwURI struct {
	Match bool   `json:"match"`
	URI   string `json:"uri"`
}

// BwLogin is a go structure mapping to bitwarden's JSON API for a login in bitwarden
type BwLogin struct {
	URIs                 []BwURI `json:"uris"`
	Username             string  `json:"username"`
	Password             string  `json:"password"`
	TOTP                 bool    `json:"totp"`
	PasswordRevisionDate *string `json:"passwordRevisionDate"`
}

// BwItem is a go structure mapping to bitwarden's JSON API for an item in bitwarden
type BwItem struct {
	Object         string           `json:"object"`
	ID             string           `json:"id"`
	OrganizationID *string          `json:"organizationId"`
	FolderID       BwFolderID       `json:"folderId"`
	Type           int              `json:"type"`
	Name           string           `json:"name"`
	Notes          string           `json:"notes"`
	Favourite      bool             `json:"favourite"`
	Login          BwLogin          `json:"login"`
	CollectionIDs  []BwCollectionID `json:"colletionIds"`
	RevisionDate   string           `json:"revisionDate"`
}
