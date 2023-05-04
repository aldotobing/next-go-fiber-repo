package viewmodel

type RegionAreaVM struct {
	ID        *string `json:"id"`
	Code      *string `json:"code"`
	Name      *string `json:"name"`
	GroupID   *string `json:"group_id"`
	GroupName *string `json:"group_name"`
}
