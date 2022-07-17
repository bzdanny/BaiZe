package bCryptPasswordEncoder

import (
	"fmt"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password := HashPassword("admin123")
	fmt.Println(password)
	hash := CheckPasswordHash("admin123", "$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2")
	fmt.Println(hash)
}
