package jwebb

// IsECDHES checks if the algorithm is an ECDH-ES based algorithm
func IsECDHES(alg string) bool { _ = "STUB: not implemented"; return false }

// IsRSA15 checks if the algorithm is RSA1_5
func IsRSA15(alg string) bool { _ = "STUB: not implemented"; return false }

// IsRSAOAEP checks if the algorithm is an RSA-OAEP based algorithm
func IsRSAOAEP(alg string) bool { _ = "STUB: not implemented"; return false }

// IsAESKW checks if the algorithm is an AES key wrap algorithm
func IsAESKW(alg string) bool { _ = "STUB: not implemented"; return false }

// IsAESGCMKW checks if the algorithm is an AES-GCM key wrap algorithm
func IsAESGCMKW(alg string) bool { _ = "STUB: not implemented"; return false }

// IsPBES2 checks if the algorithm is a PBES2 based algorithm
func IsPBES2(alg string) bool { _ = "STUB: not implemented"; return false }

// IsDirect checks if the algorithm is direct encryption
func IsDirect(alg string) bool { _ = "STUB: not implemented"; return false }

// IsMLKEM checks if the algorithm is an ML-KEM based algorithm.
// ML-KEM algorithms are contributed by external modules (e.g.,
// github.com/jwx-go/mlkem) via RegisterMLKEMAlgorithm.
func IsMLKEM(alg string) bool { _ = "STUB: not implemented"; return false }

// IsMLKEMDirect checks if the algorithm is a direct ML-KEM algorithm
// (no key wrapping). Direct algorithms are registered via
// RegisterMLKEMDirectAlgorithm.
func IsMLKEMDirect(alg string) bool { _ = "STUB: not implemented"; return false }

// IsHPKE checks if the algorithm is an HPKE-based algorithm.
// Built-in algorithms are pre-registered in the same registry
// used by RegisterHPKEAlgorithm, so all lookups follow one path.
func IsHPKE(alg string) bool { _ = "STUB: not implemented"; return false }

// IsDirectCEK checks if the algorithm uses direct key agreement
// where the CEK is derived (not encrypted). This includes DIRECT,
// bare ECDH-ES (without key wrapping), and direct ML-KEM modes.
func IsDirectCEK(alg string) bool { _ = "STUB: not implemented"; return false }

// IsSymmetric checks if the algorithm is a symmetric key encryption algorithm
func IsSymmetric(alg string) bool { _ = "STUB: not implemented"; return false }
