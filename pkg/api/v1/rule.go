package v1

type Rule struct {
	ID          string `yaml:"id,omitempty"`
	Title       string `yaml:"title,omitempty"`
	Description string `yaml:"description,omitempty"`
	Rationale   string `yaml:"rationale,omitempty"`
	Severity    string `yaml:"severity,omitempty"`
	Identifiers string `yaml:"identifiers,omitempty"`
	References  string `yaml:"references,omitempty"`
}
