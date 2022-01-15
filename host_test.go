package envaddr

import (
	"os"
	"testing"
)

func TestGetNone(t *testing.T) {
	os.Unsetenv("PORT")
	os.Unsetenv("HOST")
	os.Unsetenv("IP")
	addr := Get(":8443")
	if addr != ":8443" {
		t.Errorf("Expected 8443, got %s", addr)
	}
}

func TestGetOnlyPort(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Unsetenv("HOST")
	os.Unsetenv("IP")
	addr := Get(":8443")
	if addr != ":8080" {
		t.Errorf("Expected 8080, got %s", addr)
	}
}

func TestGetIPv4Port(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Unsetenv("HOST")
	os.Setenv("IP", "127.255.255.255")
	addr := Get(":8443")
	if addr != "127.255.255.255:8080" {
		t.Errorf("Expected 127.255.255.255:8080, got %s", addr)
	}
}

func TestGetIPv6Port(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Unsetenv("HOST")
	os.Setenv("IP", "::1")
	addr := Get(":8443")
	if addr != "[::1]:8080" {
		t.Errorf("Expected [::1]:8080, got %s", addr)
	}
}

func TestGetHostPort(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("HOST", "localhost")
	os.Unsetenv("IP")
	addr := Get(":8443")
	if addr != "localhost:8080" {
		t.Errorf("Expected localhost:8080, got %s", addr)
	}
}
