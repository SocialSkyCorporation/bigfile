//  Copyright 2019 The bigfile Authors. All rights reserved.
//  Use of this source code is governed by a MIT-style
//  license that can be found in the LICENSE file.

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistory_TableName(t *testing.T) {
	assert.Equal(t, "histories", (&History{}).TableName())
}
