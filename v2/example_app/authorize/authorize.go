package authorize

import (
	"github.com/dedefer/gb_go_mod_example2/example_app/authorize/token"
)

func Authorize() *token.Token {
	tok := token.New()
	// white token to base
	return tok
}
