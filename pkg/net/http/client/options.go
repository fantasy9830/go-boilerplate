package client

type OptionFunc func(*Options)
type CheckFunc func(body []byte) error

type Options struct {
	Headers   map[string]string
	CheckFunc CheckFunc
}

func WithHeader(key string, value string) OptionFunc {
	return func(o *Options) {
		o.Headers[key] = value
	}
}

func WithCheckFunc(f CheckFunc) OptionFunc {
	return func(o *Options) {
		o.CheckFunc = f
	}
}
