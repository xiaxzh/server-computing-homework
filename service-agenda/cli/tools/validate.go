package tools

import (
	// "fmt"
	"strconv"
	"strings"
)

func ValidatePass(p string) (bool, string) {
	if len(p) > 0 {
		return true, ""
	} else {
		return false, "Invalid password length"
	}
}

func ValidateUsername(u string) (bool, string) {
	if len(u) > 0 {
		return true, ""
	} else {
		return false, "Invalid username length"
	}
}

func ValidateEmail(e string) (bool, string) {
	aite := strings.Index(e, "@")
	point := strings.Index(e, ".")
	if len(e) > 0 && aite != -1 && point != -1 && point > aite {
		return true, ""
	} else {
		return false, "Invalid email format"
	}
}

func ValidateId(i string) (bool, string) {
	for _, i1 := range i {
		if !(i1 <= '9' && i1 >= '0') {
			return false, "Invalid ID"
		}
	}
	return true, ""
}

func ValidateOffset(i string) (bool, string) {
	for _, i1 := range i {
		if !(i1 <= '9' && i1 >= '0') {
			return false, "Invalid Offset"
		}
	}
	return true, ""
}

func ValidatePhone(p string) (bool, string) {
	if len(p) != 11 {
		return false, "Invalid phone length"
	}
	for _, i1 := range p {
		if !(i1 <= '9' && i1 >= '0') {
			return false, "Invalid phone number"
		}
	}
	return true, ""
}

func ValidateLimit(n string) (bool, string) {
	if len(n) < 0 {
		return false, "Empty limit"
	}
	if i, err := strconv.Atoi(n); err != nil || i <= 0 {
		return false, "Invalid limit"
	}
	return true, ""
}
