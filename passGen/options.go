package passGen

type Options struct {
	// Length of the new password
	Length int
	// WithNumbers is set to true if the password must contain numbers.
	WithNumbers bool
}

// SetDefaults ensure options are usable to avoid certain bugs.
func (o *Options) SetDefaults() *Options {
	if o.Length == 0 {
		o.Length = 6
	}
	return o
}
