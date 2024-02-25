package view

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type ViewCtx struct {
	Env  string
	Path string
	deps map[string]bool
}

func (vCtx *ViewCtx) UseDep(name string) {
	if vCtx.deps == nil {
		vCtx.deps = make(map[string]bool)
	}
	log.Info("Using Dependency :", name)
	vCtx.deps[name] = true
}

func (vCtx *ViewCtx) UsingDep(name string) bool {
	if vCtx.deps == nil {
		return false
	}
	var v = vCtx.deps[name]
	return v
}

func GetViewCtx(ctx context.Context) *ViewCtx {
	if ctx, ok := ctx.Value("viewCtx").(*ViewCtx); ok {
		return ctx
	}
	return &ViewCtx{}
}

func GetPath(ctx context.Context) string {
	return GetViewCtx(ctx).Path
}

func IsLocal(ctx context.Context) bool {
	return GetViewCtx(ctx).Env == "local"
}
