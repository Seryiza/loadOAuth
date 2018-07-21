# Load OAuth
Simple Go package to get oauth2 tokens/configs and to set it (if changed). I use it for unit-tests.

# Dependencies
* [github.com/golang/oauth2](https://github.com/golang/oauth2)

# Examples
## Token get and set from file
```go
import (
	"golang.org/x/oauth2"
	"github.com/seryiza/loadOAuth/token"
)

// Get token from file (by env-variable)
tok, _ := token.FromFile("SHIKI")	// filepath is $SHIKI_TOKEN_FILE

// conf := &oauth2.Config{...}
client := conf.Client(ctx, tok)

// some operations...

// Set token to file (by env-variable)
token.ToFile("SHIKI", client)
```

## Config get from file
```go
import (
	"golang.org/x/oauth2"
	"github.com/seryiza/loadOAuth/conf"
)

conf, err := conf.FromFile("SHIKI")	// filepath is $SHIKI_CONF_FILE
```