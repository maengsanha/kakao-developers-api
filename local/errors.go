// Copyright 2022 Sanha Maeng, Soyang Baek, Jinmyeong Kim
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package local

import (
	"errors"
	"internal/common"
)

var (
	Done                            = common.ErrEndPage
	ErrUnsupportedCategoryGroupCode = errors.New(
		`category group code must be one of the following options:
		MT1, CS2, PS3, SC4, AC5, PK6, OL7, SW8, CT1, AG2, PO3, AT4, FD6, CE7, HP8, PM9, BK9, AD5`)
	ErrRadiusOutOfBound = errors.New("radius must be between 0 and 20000")
)
