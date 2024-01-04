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

func (d *DataBase) InsertServer(ctx context.Context, server *model_struct.LocalServer) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Create(server).Error, "InsertServer failed")
}
func (d *DataBase) DeleteServer(ctx context.Context, serverID string) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Where("server_id=?", serverID).Delete(&model_struct.LocalServer{}).Error, "DeleteServer failed")
}
func (d *DataBase) UpdateServer(ctx context.Context, server *model_struct.LocalServer) error {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	t := d.conn.WithContext(ctx).Model(server).Select("*").Updates(*server)
	if t.RowsAffected == 0 {
		return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
	}
	return utils.Wrap(t.Error, "")
}

func (d *DataBase) GetServers(ctx context.Context, serverIDs []string) ([]*model_struct.LocalServer, error) {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	var serverList []model_struct.LocalServer
	err := utils.Wrap(d.conn.WithContext(ctx).Where("server_id in (?)", serverIDs).Order("create_time DESC").Find(&serverList).Error, "")
	if err != nil {
		return nil, utils.Wrap(err, "")
	}
	var transfer []*model_struct.LocalServer
	for _, v := range serverList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, nil
}

func (d *DataBase) GetAllServers(ctx context.Context) ([]*model_struct.LocalServer, error) {
	d.serverMtx.Lock()
	defer d.serverMtx.Unlock()
	var serverList []model_struct.LocalServer
	err := utils.Wrap(d.conn.WithContext(ctx).Order("create_time DESC").Find(&serverList).Error, "")
	if err != nil {
		return nil, utils.Wrap(err, "")
	}
	var transfer []*model_struct.LocalServer
	for _, v := range serverList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, nil
}
