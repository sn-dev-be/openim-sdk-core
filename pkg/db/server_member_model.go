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

func (d *DataBase) InsertServerMember(ctx context.Context, serverMember *model_struct.LocalServerMember) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Create(serverMember).Error, "InsertServerMember failed")
}

func (d *DataBase) DeleteServerMember(ctx context.Context, id uint64) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Where("id = ?", id).Delete(&model_struct.LocalServerMember{}).Error, "DeleteServerMember failed")
}

func (d *DataBase) DeleteServerMemberByServer(ctx context.Context, serverID string) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Where("server_id = ?", serverID).Delete(&model_struct.LocalServerMember{}).Error, "DeleteServerMember failed")
}

func (d *DataBase) UpdateServerMember(ctx context.Context, serverMember *model_struct.LocalServerMember) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	t := d.conn.WithContext(ctx).Model(serverMember).Select("*").Updates(*serverMember)
	if t.RowsAffected == 0 {
		return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
	}
	return utils.Wrap(t.Error, "")
}

func (d *DataBase) GetServerMemberByServerID(ctx context.Context, serverID string) (*model_struct.LocalServerMember, error) {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	var serverMember model_struct.LocalServerMember
	err := utils.Wrap(d.conn.WithContext(ctx).Where("server_id = ?", serverID).Order("join_time DESC").Find(&serverMember).Error, "")
	if err != nil {
		return nil, utils.Wrap(err, "")
	}

	return &serverMember, nil
}

func (d *DataBase) GetServerMembers(ctx context.Context, userID string) ([]*model_struct.LocalServerMember, error) {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	var serverMembers []*model_struct.LocalServerMember
	err := utils.Wrap(d.conn.WithContext(ctx).Where("user_id = ?", userID).Order("join_time DESC").Find(&serverMembers).Error, "")
	if err != nil {
		return nil, utils.Wrap(err, "")
	}

	return serverMembers, nil
}
