package pool

var keyToErrorMapPool = New[map[string]error](allocKeyToErrorMap, freeKeyToErrorMap)

func allocKeyToErrorMap() map[string]error { _ = "STUB: not implemented"; return nil }

func freeKeyToErrorMap(m map[string]error) map[string]error { _ = "STUB: not implemented"; return nil }

// Clear the map

// KeyToErrorMap returns a pool of map[string]error instances.
func KeyToErrorMap() *Pool[map[string]error] { _ = "STUB: not implemented"; return nil }
