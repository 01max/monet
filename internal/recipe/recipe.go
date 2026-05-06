package recipe

// Step is a single ImageMagick operation in a recipe pipeline.
type Step struct {
	Op    string `yaml:"op"`
	Value string `yaml:"value"`
}

// Output describes the output configuration for a recipe.
type Output struct {
	Format string `yaml:"format,omitempty"`
}

// Recipe is a named ImageMagick pipeline stored as a YAML file.
type Recipe struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Steps       []Step `yaml:"steps"`
	Output      Output `yaml:"output"`
}
