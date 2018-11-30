package domain

// Host represents a host
type Host struct {
	Name    string
	Address string
	Act     string
	User    string
	Key     string
}

// Hosts represents hosts
type Hosts struct {
	Name  string
	Hosts []Host
}

// HostFile represents a host file
type HostFile struct {
	Name  string
	Path  string
	Bytes []byte `json:"-"`
}
