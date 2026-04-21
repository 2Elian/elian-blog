package types

type SiteConfigVO struct {
	ID    uint   `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SetSiteConfigReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
