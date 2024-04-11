package options

type Options struct {
	Name          string
	DefaultValue  any
	DefaultObject any
	Description   string
	Sensitive     bool
	Order         int
	TypeExpress   any
	// IgnoreWidget used for skipping add widget for any type in sub schema.
	IgnoreWidget bool
}
