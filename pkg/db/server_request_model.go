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

func (d *DataBase) InsertServerRequest(ctx context.Context, serverRequest *model_struct.LocalServerRequest) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Create(serverRequest).Error, "InsertServerRequest failed")
}
func (d *DataBase) DeleteServerRequest(ctx context.Context, serverID, userID string) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Where("server_id=? and user_id=?", serverID, userID).Delete(&model_struct.LocalServerRequest{}).Error, "DeleteServerRequest failed")
}
func (d *DataBase) UpdateServerRequest(ctx context.Context, serverRequest *model_struct.LocalServerRequest) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	t := d.conn.WithContext(ctx).Model(serverRequest).Select("*").Updates(*serverRequest)
	if t.RowsAffected == 0 {
		return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
	}
	return utils.Wrap(t.Error, "")
}

func (d *DataBase) GetSendServerApplication(ctx context.Context) ([]*model_struct.LocalServerRequest, error) {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	var serverRequestList []model_struct.LocalServerRequest
	err := utils.Wrap(d.conn.WithContext(ctx).Order("create_time DESC").Find(&serverRequestList).Error, "")
	if err != nil {
		return nil, utils.Wrap(err, "")
	}
	var transfer []*model_struct.LocalServerRequest
	for _, v := range serverRequestList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, nil
}
