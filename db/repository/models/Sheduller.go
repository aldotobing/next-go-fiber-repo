package models

// SchedullerExpiredPackage ...
type SchedullerExpiredPackage struct {
	TotalCount *int `json:"total_count"`
}

// SchedullerExpiredPackageParameter ...
type SchedullerExpiredPackageParameter struct {
}

var (

	// SchedullerExpiredPackageSelectStatement ...
	SchedullerExpiredPackageSelectStatement = `select teke_attendance_for_expired_package as p_count from teke_attendance_for_expired_package(1)`

	// SchedullerExpiredPackageWhereStatement ...
	SchedullerExpiredPackageWhereStatement = ``
)
