package info

import (
	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	g.RunFn(appDetails(opts))

	// if err := g.Box(packr.New("github.com/gobuffalo/buffalo/genny/info/templates", "../info/templates")); err != nil {
	//   return g, errors.WithStack(err)
	// }
	// ctx := plush.NewContext()
	// ctx.Set("opts", opts)
	// g.Transformer(plushgen.Transformer(ctx))
	return g, nil
}
