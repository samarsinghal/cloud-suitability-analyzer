/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db_test

import (
	"testing"

	"github.com/vmware-samples/cloud-suitability-analyzer/go/db"
	"github.com/stretchr/testify/assert"
)

func TestLoadingOfRules(t *testing.T) {

	rules, _ := db.GetAllResourcesBasedRules("../../rules/")
	// Atleast 10 rules, there are 62 currently.

	assert.True(t, len(rules) > 10)
}
