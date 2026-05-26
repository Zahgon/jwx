package types

type StringList []string

func (l StringList) Get() []string { _ = "STUB: not implemented"; return nil }

func (l *StringList) Accept(v any) error { _ = "STUB: not implemented"; return nil }

func (l *StringList) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
