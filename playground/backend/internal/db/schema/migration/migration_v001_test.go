// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package migration

import (
	"context"
	"os"
	"testing"

	"beam.apache.org/playground/backend/internal/app_constants"
	"beam.apache.org/playground/backend/internal/db/datastore"
	"beam.apache.org/playground/backend/internal/db/mapper"
	"beam.apache.org/playground/backend/internal/db/schema"
	"beam.apache.org/playground/backend/internal/environment"
)

var datastoreDb *datastore.Datastore
var ctx context.Context

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	datastoreEmulatorHost := os.Getenv(app_constants.EmulatorHostKey)
	if datastoreEmulatorHost == "" {
		if err := os.Setenv(app_constants.EmulatorHostKey, app_constants.EmulatorHostValue); err != nil {
			panic(err)
		}
	}
	ctx = context.Background()
	ctx = context.WithValue(ctx, app_constants.DatastoreNamespaceKey, "migration")
	var err error
	datastoreDb, err = datastore.New(ctx, mapper.NewPrecompiledObjectMapper(), app_constants.EmulatorProjectId)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	if err := datastoreDb.Client.Close(); err != nil {
		panic(err)
	}
}

func TestInitialStructure_InitiateData(t *testing.T) {
	appEnvs := environment.NewApplicationEnvs("/app", "", "", "", "../../../../../sdks-emulator.yaml", "../../../../.", nil, 0)
	props, err := environment.NewProperties(appEnvs.PropertyPath())
	if err != nil {
		t.Errorf("InitiateData(): error during properties initialization, err: %s", err.Error())
	}
	tests := []struct {
		name    string
		dbArgs  *schema.DBArgs
		wantErr bool
	}{
		{
			name: "Test migration with version 0.0.1 in the usual case",
			dbArgs: &schema.DBArgs{
				Ctx:    ctx,
				Db:     datastoreDb,
				AppEnv: appEnvs,
				Props:  props,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := new(InitialStructure)
			err = is.InitiateData(tt.dbArgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitiateData(): error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
