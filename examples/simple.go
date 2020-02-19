package main

import (
	"fmt"

	"github.com/lucat1/helmet"
	"github.com/lucat1/randr"
)

func example(ctx randr.Context) string {
	return randr.HTML(`
		<h1>Hello world!</h1>

		<helmet.Head>
			<title>Hello world website</title>
		</helmet.Head>
	`)
}

func main() {
	res, ctx := randr.Render(example, nil)

	fmt.Printf("Result:\n%s\nHead:\n%s\n", res, helmet.ExtractHead(ctx))
}