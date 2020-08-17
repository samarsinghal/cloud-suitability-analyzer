/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package db

import (
	"github.com/jinzhu/gorm"
	"github.com/vmware-samples/cloud-suitability-analyzer/go/model"
)

type ReportDataRepository interface {
	SaveReportData(reportData *model.ReportData) error
}

func NewReportDataRepository(db *gorm.DB) ReportDataRepository {
	return &OrmRepository{
		dbconn: db,
	}
}

func NewReportDataRepositoryForRun(run *model.Run) ReportDataRepository {
	return &OrmRepository{
		dbconn: run.DB,
	}
}

func (reportDataRepository *OrmRepository) SaveReportData(reportData *model.ReportData) error {
	res := reportDataRepository.dbconn.Create(reportData)
	return res.Error
}
