package response

type ApplicationModel struct {
	Id         int64    `json:"id"`
	Uid        int64    `json:"uid"`
	Name       string   `json:"name"`
	OwnerLogin string   `json:"owner_login"`
	Permission string   `json:"permission"`
	Features   []string `json:"features"`
	CreateDate string   `json:"create_date"`
	TimeZone   string   `json:"time_zone_name"`
}
