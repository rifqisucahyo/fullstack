package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Nickname Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	if strings.Contains(err, "record not found") {
		return errors.New("Data tidak ditemukan")
	}
	return errors.New(err)
}

func FormatErrorMsg(err string) (map[string]interface{}, error) {
	param := make(map[string]interface{})

	if strings.Contains(err, "email") {
		param["error"] = "Email not found."

		return param, nil
	}

	return param, nil
}
