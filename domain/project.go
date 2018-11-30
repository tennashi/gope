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

/*
func NewProject(name, basePath string) (*Project, error) {
	return &Project{
		Name:     name,
		BasePath: basePath,
	}, nil
}

// TODO: to Util
func getName(filePath string) string {
	_, fileName := filepath.Split(filePath)
	name := strings.TrimSuffix(fileName, ".toml")
	return name
}
*/
