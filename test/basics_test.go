/*
 * Copyright (C) 2020 ObjectBox Ltd. All rights reserved.
 * https://objectbox.io
 *
 * This file is part of ObjectBox Generator.
 *
 * ObjectBox Generator is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * ObjectBox Generator is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with ObjectBox Generator.  If not, see <http://www.gnu.org/licenses/>.
 */

package test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/objectbox/objectbox-generator/test/assert"
)

// Because of Go generator comparison tests, the go tool may update go.mod file to import `github.com/objectbox/objectbox-go`
// That would cause a circular dependency as objectbox-go needs to import the generator to provide the command-line
// generation that was there since the beginning. This test ensures such updated go.mod file isn't allowed to stay.
func TestCircularDependencies(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoErr(t, err)

	goModPath := filepath.Join(cwd, "..", "go.mod")
	goModBytes, err := ioutil.ReadFile(goModPath)
	assert.NoErr(t, err)

	goMod := string(goModBytes)

	// check the content looks as expected first
	assert.True(t, strings.Contains(goMod, "module github.com/objectbox/objectbox-generator"))

	// and check it doesn't contain the objectbox-go reference
	assert.True(t, !strings.Contains(goMod, "github.com/objectbox/objectbox-go"))
}
