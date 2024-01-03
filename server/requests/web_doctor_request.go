package requests

// CustomerRequest ...
type WebDoctorRequest struct {
	ID            string `json:"doctor_id"`
	Code          string `json:"doctor_code"`
	DoctorName    string `json:"doctor_name"`
	DoctorAddress string `json:"doctor_address"`
	DoctorPhone   string `json:"doctor_phone"`
	DoctorUserID  string `json:"doctor_user_id"`
}
