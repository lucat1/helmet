package helmet

import "github.com/lucat1/randr"

// HeadProps are the props that the head
// component can receive
type HeadProps randr.BasicProps

// Head is a custom element used to manage
// all your header tags, similarly to `react-helmet`
func Head(ctx randr.Context) string {
	props := ctx.Props.(*HeadProps)
	// Check that the context has been set
	// if not, we define it and we're done
	if ctx.Data[HelmetContext] == nil {
		ctx.Data[HelmetContext] = props.Children
		return ""
	}

	if val, ok := ctx.Data[HelmetContext].(string); ok {
		ctx.Data[HelmetContext] = val + props.Children
	}

	// Never render anything
	return ""
}