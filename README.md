# GITLAB API library

This library is aimed for cooperative access of many users to a single gitlab instance.

## Usage

```go
package main

import (
	"context"
	"log"
	
	"github.com/sirkon/gitlab"
)

func main() {
	access := gitlab.NewAPIAccess(nil, "gitlab.com/api/v4")
	client := access.Client("user-token")
	
	tags, err := client.Tags(context.Background(), "user/project", "")
	if err != nil {
		log.Fatal(err)
	}
	
	log.Printf("%#v", tags)
}
```