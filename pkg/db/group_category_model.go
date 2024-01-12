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
)

func (d *DataBase) InsertGroupCategory(ctx context.Context, groupCategory *model_struct.LocalGroupCategory) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Create(groupCategory).Error, "InsertGroupCategory failed")
}
func (d *DataBase) DeleteGroupCategory(ctx context.Context, categoryID string) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Where("category_id=?", categoryID).Delete(&model_struct.LocalGroupCategory{}).Error, "DeleteGroupCategory failed")
}

func (d *DataBase) DeleteGroupCategoryByServers(ctx context.Context, serverIDs []string) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Where("server_id in (?)", serverIDs).Delete(&model_struct.LocalGroupCategory{}).Error, "DeleteGroupCategory failed")
}

func (d *DataBase) UpdateGroupCategory(ctx context.Context, groupCategory *model_struct.LocalGroupCategory) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	t := d.conn.WithContext(ctx).Model(groupCategory).Select("*").Updates(*groupCategory)
	if t.RowsAffected == 0 {
		return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
	}
	return utils.Wrap(t.Error, "")
}

func (d *DataBase) GetGroupCategories(ctx context.Context, categoryIDs []string) ([]*model_struct.LocalGroupCategory, error) {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	var groupCategoryList []model_struct.LocalGroupCategory
	err := utils.Wrap(d.conn.WithContext(ctx).Where("category_id in ?", categoryIDs).Order("create_time DESC").Find(&groupCategoryList).Error, "")
	if err != nil {
		return nil, utils.Wrap(err, "")
	}
	var transfer []*model_struct.LocalGroupCategory
	for _, v := range groupCategoryList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, nil
}

func (d *DataBase) GetGroupCategoriesByServer(ctx context.Context, serverIDs []string) ([]*model_struct.LocalGroupCategory, error) {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	var groupCategoryList []model_struct.LocalGroupCategory
	err := utils.Wrap(d.conn.WithContext(ctx).Where("server_id in ?", serverIDs).Order("create_time DESC").Find(&groupCategoryList).Error, "")
	if err != nil {
		return nil, utils.Wrap(err, "")
	}
	var transfer []*model_struct.LocalGroupCategory
	for _, v := range groupCategoryList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, nil
}
