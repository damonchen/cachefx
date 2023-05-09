# cachefx

example：

```go

import (
    "github.com/damonchen/cachefx"
    "go.uber.org/fx"
)

// using memory cache
fx.Module("cache",
    fx.Provide(
        fx.Annotate(
            cachefx.NewMemoryCache,
            fx.As(new(Cache)),
            fx.ResultTags(`name:"cache"`)),
    ))

```
