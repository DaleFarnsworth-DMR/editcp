// Copyright 2017 Dale Farnsworth. All rights reserved.

// Dale Farnsworth
// 1007 W Mendoza Ave
// Mesa, AZ  85210
// USA
//
// dale@farnsworth.org

// This file is part of Editcp.
//
// Editcp is free software: you can redistribute it and/or modify
// it under the terms of version 3 of the GNU General Public License
// as published by the Free Software Foundation.
//
// Editcp is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Editcp.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/dalefarnsworth/codeplug/codeplug"
	"github.com/dalefarnsworth/codeplug/ui"
)

func scanLists(edt *editor) {
	edt.recordWindow(codeplug.RtScanLists, slRecord)
}

func slRecord(edt *editor, recordBox *ui.HBox) {
	column := recordBox.AddVbox()
	column.SetContentsMargins(0, 0, 0, 0)
	addFieldMembers(column, &settings.sortAvailableChannels,
		codeplug.FtSlName, codeplug.FtSlChannelMembers, "Channels")

	row := column.AddHbox()
	row.SetContentsMargins(0, 0, 0, 0)

	r := currentRecord(recordBox.Window())

	column = row.AddVbox()
	column.SetContentsMargins(0, 0, 0, 0)
	form := column.AddForm()
	form.AddFieldRows(r,
		codeplug.FtSlPriorityChannel1,
		codeplug.FtSlPriorityChannel2,
		codeplug.FtSlTxDesignatedChannel)

	column = row.AddVbox()
	column.SetContentsMargins(0, 0, 0, 0)
	form = column.AddForm()
	form.AddFieldRows(r,
		codeplug.FtSlSignallingHoldTime,
		codeplug.FtSlPrioritySampleTime)
}
