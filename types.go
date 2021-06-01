package bitwardenwrapper

type BwFolderId string
type BwCollectionId string

type BwFolder struct {
	Object string     `json:"object"`
	Id     BwFolderId `json:"id"`
	Name   string     `json:"name"`
}

type BwUri struct {
	Match bool   `json:"match"`
	URI   string `json:"uri"`
}

type BwLogin struct {
	Uris                 []BwUri `json:"uris"`
	Username             string  `json:"username"`
	Password             string  `json:"password"`
	TOTP                 bool    `json:"totp"`
	PasswordRevisionDate *string `json:"passwordRevisionDate"`
}

type BwItem struct {
	Object         string           `json:"object"`
	Id             string           `json:"id"`
	OrganizationId *string          `json:"organizationId"`
	FolderId       BwFolderId       `json:"folderId"`
	Type           int              `json:"type"`
	Name           string           `json:"name"`
	Notes          string           `json:"notes"`
	Favourite      bool             `json:"favourite"`
	Login          BwLogin          `json:"login"`
	CollectionIds  []BwCollectionId `json:"colletionIds"`
	RevisionDate   string           `json:"revisionDate"`
}
