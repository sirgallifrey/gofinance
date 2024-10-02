package hx

import "github.com/a-h/templ"

const (
	SwapInnerHTML   = "innerHTML"
	SwapOuterHTML   = "outerHTML"
	SwapBeforebegin = "beforebegin"
	SwapAfterbegin  = "afterbegin"
	SwapBeforeend   = "beforeend"
	SwapAfterend    = "afterend"
	SwapDelete      = "delete"
	SwapNone        = "none"
)

type Htmx struct {
	ok       bool
	swap     string
	method   string
	endpoint string
}

func (h Htmx) ToAttr() templ.Attributes {
	attrs := make(templ.Attributes)
	if h.swap != "" {
		attrs["hx-swap"] = h.swap
	}
	if h.method != "" && h.endpoint != "" {
		attrs[h.method] = h.endpoint
	}
	return attrs
}

func Hx(options ...func(*Htmx)) Htmx {
	h := &Htmx{}
	if len(options) > 0 {
		h.ok = true
	}
	for _, o := range options {
		o(h)
	}
	return *h
}

func Swap(swap string) func(*Htmx) {
	return func(h *Htmx) {
		h.swap = swap
	}
}

func request(method string, endpoint string) func(*Htmx) {
	return func(h *Htmx) {
		h.method = method
		h.endpoint = endpoint
	}
}

func Get(endpoint string) func(*Htmx) {
	return request("hx-get", endpoint)
}

func Post(endpoint string) func(*Htmx) {
	return request("hx-post", endpoint)
}

func Put(endpoint string) func(*Htmx) {
	return request("hx-put", endpoint)
}

func Patch(endpoint string) func(*Htmx) {
	return request("hx-patch", endpoint)
}

func Delete(endpoint string) func(*Htmx) {
	return request("hx-delete", endpoint)
}
