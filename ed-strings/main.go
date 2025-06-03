package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IdentifyPrefixPostfix(userID, email string) bool {
	if len(userID) < 4 || len(email) < 4 {
		return false // Ensure both userID and email are long enough
	}

	prefix := userID[:3] // Get the first three characters of userID
	postfix := userID[len(userID)-3:] // Get the last three characters of userID

	// Check if email starts with prefix and ends with postfix
	return strings.HasPrefix(email, prefix) && strings.HasSuffix(email, postfix)
}

func ContainsEducative(email string) bool {
	return strings.Contains(email, "educative.io") || strings.Contains(email, "educative.com")
}

func MaskUserName(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return email // No '@' found, return original email
	}
	userName := email[:atIndex]
	if len(userName) <= 2 {
		return email // No masking needed for short usernames
	}
	maskedUserName := userName[:1] + strings.Repeat("*", len(userName)-2) + userName[len(userName)-1:]
	return maskedUserName + email[atIndex:]
}

func IndexOfAtSymbol(email string) int {
	return strings.Index(email, "@")
}

func TrimAndSplitUserID(userID string) string {
	if len(userID) < 4 {
		return "" // Ensure userID is long enough
	}
	trimmed := strings.TrimPrefix(userID, "UID-")
	if len(trimmed) < 1 {
		return "" // Return empty if no valid ID found
	}
	return trimmed
}

func ConvertStringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0 // Return 0 if conversion fails
	}
	return num
}

func main() {
	// Test your functions here
	fmt.Println(IdentifyPrefixPostfix("UID-0123", "evangeline@educative.io")) // true
	fmt.Println(ContainsEducative("evangeline@educative.io"))                 // true
	fmt.Println(MaskUserName("evangeline@educative.io"))                      // e******e@educative.io
	fmt.Println(IndexOfAtSymbol("evangeline@educative.io"))                   // 10
	fmt.Println(TrimAndSplitUserID("UID-0123"))                               // 0123
	fmt.Println(ConvertStringToInt("123"))                                    // 123
}
