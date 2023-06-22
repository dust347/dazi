package uuid

import "github.com/google/uuid"

// New 创建 uuid
func New() string {
	return uuid.New().String()
}
