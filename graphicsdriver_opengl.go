// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build android darwin,ios,386 darwin,ios,amd64 freebsd js linux windows

// As the Go playground tries to compile this with CGO_ENABLED=0 and GOOS=linux, check Cgo on build tags.

package ebiten

import (
	"github.com/hajimehoshi/ebiten/internal/driver"
	"github.com/hajimehoshi/ebiten/internal/graphicsdriver/opengl"
)

func graphicsDriver() driver.Graphics {
	return opengl.Get()
}
