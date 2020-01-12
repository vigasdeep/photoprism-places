package opencage

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ProviderKey = "f9585987890741ceac625709e8efea3b"

	os.Exit(m.Run())
}
