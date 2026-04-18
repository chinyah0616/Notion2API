package app

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveBrowserExecutablePathPrefersCHROMEBIN(t *testing.T) {
	tmpDir := t.TempDir()
	chromeBin := filepath.Join(tmpDir, "chrome-bin")
	chromiumBin := filepath.Join(tmpDir, "chromium-bin")
	if err := os.WriteFile(chromeBin, []byte(""), 0o644); err != nil {
		t.Fatalf("write chrome bin failed: %v", err)
	}
	if err := os.WriteFile(chromiumBin, []byte(""), 0o644); err != nil {
		t.Fatalf("write chromium bin failed: %v", err)
	}

	t.Setenv("CHROME_BIN", chromeBin)
	t.Setenv("CHROMIUM_BIN", chromiumBin)

	if got := resolveBrowserExecutablePath(); got != chromeBin {
		t.Fatalf("resolveBrowserExecutablePath() = %q, want %q", got, chromeBin)
	}
}

func TestResolveBrowserExecutablePathSkipsMissingCHROMEBIN(t *testing.T) {
	tmpDir := t.TempDir()
	chromiumBin := filepath.Join(tmpDir, "chromium-bin")
	if err := os.WriteFile(chromiumBin, []byte(""), 0o644); err != nil {
		t.Fatalf("write chromium bin failed: %v", err)
	}

	t.Setenv("CHROME_BIN", filepath.Join(tmpDir, "missing-chrome"))
	t.Setenv("CHROMIUM_BIN", chromiumBin)

	if got := resolveBrowserExecutablePath(); got != chromiumBin {
		t.Fatalf("resolveBrowserExecutablePath() = %q, want %q", got, chromiumBin)
	}
}
