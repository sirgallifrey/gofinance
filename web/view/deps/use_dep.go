package deps

import (
	"context"
	"gofinance/web/view"
	"io"

	"github.com/a-h/templ"
)

func UseDep(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		view.GetViewCtx(ctx).UseDep(name)
		return nil
	})
}
