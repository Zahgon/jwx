package jose

import (
	"context"
	"io"
	"sync"
	"testing"
)

var executablePath string
var muExecutablePath sync.RWMutex

func init() {
	findExecutable()
}

func SetExecutable(path string) { _ = "STUB: not implemented"; return }

func findExecutable() { _ = "STUB: not implemented"; return }

func ExecutablePath() string { _ = "STUB: not implemented"; return "" }

func Available() bool { _ = "STUB: not implemented"; return false }

func RunJoseCommand(ctx context.Context, t *testing.T, args []string, outw, errw io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

type AlgorithmSet struct {
	data map[string]struct{}
}

func NewAlgorithmSet() *AlgorithmSet { _ = "STUB: not implemented"; return nil }

func (set *AlgorithmSet) Add(s string) { _ = "STUB: not implemented"; return }

func (set *AlgorithmSet) Has(s string) bool { _ = "STUB: not implemented"; return false }

func Algorithms(ctx context.Context, t *testing.T) (*AlgorithmSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateJwk creates a new key using the jose tool, and returns its filename and
// a cleanup function.
// The caller is responsible for calling the cleanup
// function and make sure all resources are released
func GenerateJwk(ctx context.Context, t *testing.T, template string) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// EncryptJwe creates an encrypted JWE message and returns its filename and
// a cleanup function.
// The caller is responsible for calling the cleanup
// function and make sure all resources are released
func EncryptJwe(ctx context.Context, t *testing.T, payload []byte, alg string, keyfile string, enc string, compact bool) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func DecryptJwe(ctx context.Context, t *testing.T, cfile, kfile string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FmtJwe(ctx context.Context, t *testing.T, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignJws signs a message and returns its filename and
// a cleanup function.
// The caller is responsible for calling the cleanup
// function and make sure all resources are released
func SignJws(ctx context.Context, t *testing.T, payload []byte, keyfile string, compact bool) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func VerifyJws(ctx context.Context, t *testing.T, cfile, kfile string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
