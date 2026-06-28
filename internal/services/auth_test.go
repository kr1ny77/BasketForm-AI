package services

import (
	"testing"
)

func TestAuthService_Register(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	user, err := auth.Register("test@example.com", "testuser", "password123")
	if err != nil {
		t.Fatalf("Register error: %v", err)
	}
	if user.Email != "test@example.com" {
		t.Fatalf("expected email test@example.com, got %s", user.Email)
	}
	if user.Nickname != "testuser" {
		t.Fatalf("expected nickname testuser, got %s", user.Nickname)
	}
	if user.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	if user.PasswordHash == "" {
		t.Fatal("expected non-empty password hash")
	}
}

func TestAuthService_Register_DuplicateEmail(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	_, err := auth.Register("test@example.com", "user1", "password123")
	if err != nil {
		t.Fatalf("first register error: %v", err)
	}

	_, err = auth.Register("test@example.com", "user2", "password456")
	if err != ErrEmailTaken {
		t.Fatalf("expected ErrEmailTaken, got %v", err)
	}
}

func TestAuthService_Register_DuplicateNickname(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	_, err := auth.Register("test1@example.com", "testuser", "password123")
	if err != nil {
		t.Fatalf("first register error: %v", err)
	}

	_, err = auth.Register("test2@example.com", "testuser", "password456")
	if err != ErrNicknameTaken {
		t.Fatalf("expected ErrNicknameTaken, got %v", err)
	}
}

func TestAuthService_Login_Success(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	_, err := auth.Register("test@example.com", "testuser", "password123")
	if err != nil {
		t.Fatalf("register error: %v", err)
	}

	token, user, err := auth.Login("test@example.com", "password123")
	if err != nil {
		t.Fatalf("Login error: %v", err)
	}
	if token == "" {
		t.Fatal("expected non-empty token")
	}
	if user.Email != "test@example.com" {
		t.Fatalf("expected email test@example.com, got %s", user.Email)
	}
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	_, err := auth.Register("test@example.com", "testuser", "password123")
	if err != nil {
		t.Fatalf("register error: %v", err)
	}

	_, _, err = auth.Login("test@example.com", "wrongpassword")
	if err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestAuthService_Login_NonExistentEmail(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	_, _, err := auth.Login("nonexistent@example.com", "password123")
	if err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestAuthService_ValidateToken(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	_, err := auth.Register("test@example.com", "testuser", "password123")
	if err != nil {
		t.Fatalf("register error: %v", err)
	}

	token, _, err := auth.Login("test@example.com", "password123")
	if err != nil {
		t.Fatalf("login error: %v", err)
	}

	userID, err := auth.ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken error: %v", err)
	}
	if userID == "" {
		t.Fatal("expected non-empty user ID")
	}
}

func TestAuthService_ValidateToken_Invalid(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	_, err := auth.ValidateToken("invalid.token.here")
	if err != ErrInvalidToken {
		t.Fatalf("expected ErrInvalidToken, got %v", err)
	}
}

func TestAuthService_ChangePassword(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	user, _ := auth.Register("test@example.com", "testuser", "password123")

	err := auth.ChangePassword(user.ID, "password123", "newpassword456")
	if err != nil {
		t.Fatalf("ChangePassword error: %v", err)
	}

	_, _, err = auth.Login("test@example.com", "password123")
	if err != ErrInvalidCredentials {
		t.Fatalf("old password should not work after change")
	}

	token, _, err := auth.Login("test@example.com", "newpassword456")
	if err != nil {
		t.Fatalf("new password should work: %v", err)
	}
	if token == "" {
		t.Fatal("expected non-empty token after password change")
	}
}

func TestAuthService_ChangePassword_WrongOldPassword(t *testing.T) {
	upload, results := tempDirs(t)
	s := NewStorage(upload, results)
	auth := NewAuthService(s)

	user, _ := auth.Register("test@example.com", "testuser", "password123")

	err := auth.ChangePassword(user.ID, "wrongoldpassword", "newpassword456")
	if err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}
