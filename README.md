# Ginmill

Ginmill helps you create Gin servers with pre-defined, reusable API routes (called "Features").
It is designed to make your Gin-based web services modular, compatible, and easy to extend by focusing on handler implementation instead of repetitive route definitions.

## Features

- Compose Gin servers from reusable route sets ("Features")
- Standardize API endpoints for compatibility
- Focus on implementing business logic, not boilerplate

## Quick Start

### Installation

Add Ginmill to your Go project:

```sh
go get github.com/ginmills/ginmill
```

### Usage Example

Suppose you want to provide a set of OAuth endpoints as a reusable feature:

```go
package mastodon

import (
	"github.com/gin-gonic/gin"
	"github.com/ginmills/ginmill"
)

// IMastodon defines the required handler methods
type IMastodon interface {
	OAuthAuthorize(c *gin.Context)      // GET /oauth/authorize
	OAuthObtainToken(c *gin.Context)   // POST /oauth/token
	OAuthRevokeToken(c *gin.Context)   // POST /oauth/revoke
}

// Features returns a ginmill.Features for Mastodon-compatible OAuth endpoints
func Features(m IMastodon) *ginmill.Features {
	r := gin.New()
	oauth := r.Group("/oauth")
	oauth.GET("/authorize", m.OAuthAuthorize)
	oauth.POST("/token", m.OAuthObtainToken)
	oauth.POST("/revoke", m.OAuthRevokeToken)
	// Add more routes as needed
	return ginmill.NewFeatures(r.Routes())
}
```

Now, you can compose your Gin server with these features:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ginmills/ginmill"
	"your/module/mastodon"
)

type myMastodon struct{}

func (m *myMastodon) OAuthAuthorize(c *gin.Context)    { /* ... */ }
func (m *myMastodon) OAuthObtainToken(c *gin.Context)  { /* ... */ }
func (m *myMastodon) OAuthRevokeToken(c *gin.Context)  { /* ... */ }

func main() {
	engine := gin.New()
	server := &ginmill.Server{Engine: engine}
	features := mastodon.Features(&myMastodon{})
	server.With(features)
	engine.Run(":8080")
}
```

## Why Ginmill?

- **Reusable:** Define a set of routes once, reuse everywhere.
- **Compatible:** Make your API compatible with well-known standards (e.g., Mastodon, ActivityPub, etc.).
- **Modular:** Cleanly separate route definitions from handler logic.

## More Examples

- See [mastodon features](https://github.com/ginmills/mastodon) for a full implementation.

---

Ginmill is MIT licensed. Contributions welcome!
