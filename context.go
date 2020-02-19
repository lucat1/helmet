package helmet

import "github.com/lucat1/randr"

// HelmetContext is the key used to retrive the
// helmet library context (which is a string)
const HelmetContext randr.ContextKey = "helmet"

// ExtractHead extracts the given head elements
// and returns them as a string
func ExtractHead(ctx randr.Context) string {
	// Handle cases where the context has never
	// been initialized, which means no
	// Head components have been rendered
	if ctx.Data[HelmetContext] == nil {
		return ""
	}

	return ctx.Data[HelmetContext].(string)
}