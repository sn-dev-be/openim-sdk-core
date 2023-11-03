// Copyright © 2023 OpenIM SDK. All rights reserved.
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

//go:build !js
// +build !js

package db

import (
	"context"
	"errors"

	"github.com/openimsdk/openim-sdk-core/v3/pkg/db/model_struct"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"
	"gorm.io/gorm"
)

func (d *DataBase) InsertGroupSaved(ctx context.Context, groupSaved *model_struct.LocalGroupSaved) error {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Create(groupSaved).Error, "InsertGroupSaved failed")
}
func (d *DataBase) DeleteGroupSaved(ctx context.Context, groupID string) error {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Where("group_id=?", groupID).Delete(model_struct.LocalGroupSaved{}).Error, "DeleteSaved failed")
}

func (d *DataBase) UpdateGroupSaved(ctx context.Context, groupSaved *model_struct.LocalGroupSaved) error {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	t := d.conn.WithContext(ctx).Model(groupSaved).Select("*").Updates(*groupSaved)
	if t.RowsAffected == 0 {
		return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
	}
	return utils.Wrap(t.Error, "")
}

func (d *DataBase) GetGroupSavedListDB(ctx context.Context) ([]*model_struct.LocalGroupSaved, error) {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	var groupSavedList []model_struct.LocalGroupSaved

	err := d.conn.WithContext(ctx).Find(&groupSavedList).Error
	if err != nil {
		return nil, err
	}
	var transfer []*model_struct.LocalGroupSaved
	for _, v := range groupSavedList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, utils.Wrap(err, "GetGroupSavedListDB failed ")
}

func (d *DataBase) GetGroupSavedListSplit(ctx context.Context, offset, count int) ([]*model_struct.LocalGroupSaved, error) {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	var groupSavedList []model_struct.LocalGroupSaved

	err := d.conn.WithContext(ctx).Order("create_time DESC").Offset(offset).Limit(count).Find(&groupSavedList).Error
	if err != nil {
		return nil, err
	}
	var transfer []*model_struct.LocalGroupSaved
	for _, v := range groupSavedList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, utils.Wrap(err, "GetJoinedGroupList failed ")
}

func (d *DataBase) IsSaved(ctx context.Context, groupID string) (bool, error) {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	local_group := &model_struct.LocalGroupSaved{}
	err := d.conn.WithContext(ctx).Where("user_id = ? and group_id = ?", d.loginUserID, groupID).Take(&local_group).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, err
}
