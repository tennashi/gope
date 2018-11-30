package domain

// Procedure represents a procedure.
type Procedure struct {
	Name     string
	Path     string
	Sections []*Section
}

// Section represents a section.
type Section struct {
	Type  string
	Value []byte
}

// ProcedureFile represents a procedure file.
type ProcedureFile struct {
	Name  string
	Path  string
	Bytes []byte `json:"-"`
}
