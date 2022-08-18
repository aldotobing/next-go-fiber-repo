package helper

import "testing"

func TestCheckPassword(t *testing.T) {
	successResult := CheckPassword("Testing123")
	if successResult != nil {
		t.Errorf("Failed, expected nil, got %v", successResult.Error())
	}

	errResult := CheckPassword("testing")
	if errResult == nil {
		t.Errorf("Failed, expected error, got nil")
	} else if errResult.Error() != PasswordLength {
		t.Errorf("Failed, expected %v, got %v", PasswordLength, errResult.Error())
	}
}
