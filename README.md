## Kakao Developers Client

### Kakao Developers client library

[![Go Reference](https://pkg.go.dev/badge/github.com/maengsanha/kakao-developers-client.svg)](https://pkg.go.dev/github.com/maengsanha/kakao-developers-client)

#### Installation

```sh
go get -u github.com/maengsanha/kakao-developers-client
```

#### Features

* [x] Local
  - Address Search
  - Place search by keyword
  - Place search by category
  - Convert coordinates to administrative information
  - Convert coordinates to address
  - Convert coordinate system

* [x] Daum Search
  - Web document search
  - Video search
  - Image search
  - Blog search
  - Book search
  - Cafe search

* [x] Translation
  - Text translation
  - Language detection

* [x] Pose
  - Analyze image
  - Analyze video
  - check video analysis results

* [x] Vision
  - Face detection
  - Product detection
  - Adult image detection
  - Thumbnail creation
  - Multi-tag creation
  - OCR

#### Quick start

```go
package main

import "github.com/maengsanha/kakao-developers-client/local"

func main() {
  it := local.AddressSearch("을지로").
              AuthorizeWith("deadbeef").
		          Analyze("similar").
		          FormatAs("json").
		          Display(30).
		          Result(1)

	for {
		item, err := it.Next()
		if err == local.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		t.Log(item)
	}
}
```

#### License

This library is licensed under the Apache License, Version 2.0;
<br/>
see [LICENSE](https://github.com/maengsanha/kakao-developers-client/blob/master/LICENSE) for the full license text.

#### References

  - [Kakao developers documentations](https://developers.kakao.com/)
