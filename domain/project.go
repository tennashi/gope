package domain

// Project represent a project
type Project struct {
	Name     string
	BasePath string `toml:"base_path"`
}

// Projects represent projects
type Projects struct {
	Projects []*Project
}
