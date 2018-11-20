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
	Hosts []Host
}
