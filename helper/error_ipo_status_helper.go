package helper

// ErrorMessageString ...
func ErrorMessageString(code string, message string) string {
	switch {
	case code == "301":
		// return "unknown_app_id"
		return message
	case code == "302":
		// return "invalid_access"
		return message
	case code == "303":
		// return "unknown_event_id"
		return message
	case code == "304":
		// return "invalid_limit"
		return message
	case code == "305":
		return "order_already_exist"
	case code == "306":
		// return "data_not_found"
		return message
	case code == "307":
		// return "exp_password"
		return message
	case code == "308":
		// return "invalid_password"
		return message
	case code == "309":
		// return "invalid_user_id"
		return message
	default:
		return InternalServer
	}

}
