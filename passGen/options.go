package passGen

type Options struct {
	Length      int
	WithNumbers bool
}

func (o *Options) SetDefaults() *Options {
	if o.Length == 0 {
		o.Length = 6
	}
	return o
}
