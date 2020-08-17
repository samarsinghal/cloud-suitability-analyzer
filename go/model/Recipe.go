/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package model

import "time"

type Recipe struct {
	ID        uint      `gorm:"primary_key" json:"-" yaml:"-"`
	CreatedAt time.Time `json:"-" yaml:"-"`
	UpdatedAt time.Time `json:"-" yaml:"-"`
	RuleID    uint      `sql:"type:bigint REFERENCES rules(id) ON DELETE CASCADE" json:"-"  yaml:"-"`
	Rule      Rule      `gorm:"foreignkey:RuleID" json:"-"  yaml:"-"`
	URI       string    `gorm:"type:text"`
}
