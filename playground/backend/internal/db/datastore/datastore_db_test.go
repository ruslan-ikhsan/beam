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

package datastore

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/datastore"

	pb "beam.apache.org/playground/backend/internal/api/v1"
	"beam.apache.org/playground/backend/internal/constants"
	"beam.apache.org/playground/backend/internal/db/entity"
	"beam.apache.org/playground/backend/internal/db/mapper"
	"beam.apache.org/playground/backend/internal/utils"
)

const (
	datastoreEmulatorHostKey   = "DATASTORE_EMULATOR_HOST"
	datastoreEmulatorHostValue = "127.0.0.1:8888"
	datastoreEmulatorProjectId = "test"
)

var datastoreDb *Datastore
var ctx context.Context

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	datastoreEmulatorHost := os.Getenv(datastoreEmulatorHostKey)
	if datastoreEmulatorHost == "" {
		if err := os.Setenv(datastoreEmulatorHostKey, datastoreEmulatorHostValue); err != nil {
			panic(err)
		}
	}
	ctx = context.Background()
	var err error
	datastoreDb, err = New(ctx, mapper.NewPrecompiledObjectMapper(), datastoreEmulatorProjectId)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	if err := datastoreDb.Client.Close(); err != nil {
		panic(err)
	}
}

func TestDatastore_PutSnippet(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   string
		snip *entity.Snippet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "PutSnippet() in the usual case",
			args: args{ctx: ctx, id: "MOCK_ID", snip: &entity.Snippet{
				IDMeta: &entity.IDMeta{
					Salt:     "MOCK_SALT",
					IdLength: 11,
				},
				Snippet: &entity.SnippetEntity{
					Sdk:           utils.GetSdkKey(pb.Sdk_SDK_GO.String()),
					PipeOpts:      "MOCK_OPTIONS",
					Origin:        constants.UserSnippetOrigin,
					NumberOfFiles: 1,
				},
				Files: []*entity.FileEntity{{
					Name:    "MOCK_NAME",
					Content: "MOCK_CONTENT",
					IsMain:  false,
				}},
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := datastoreDb.PutSnippet(tt.args.ctx, tt.args.id, tt.args.snip)
			if err != nil {
				t.Error("PutSnippet() method failed")
			}
		})
	}

	cleanData(t, constants.FileKind, "MOCK_ID_0", nil)
	cleanData(t, constants.SnippetKind, "MOCK_ID", nil)
}

func TestDatastore_GetSnippet(t *testing.T) {
	nowDate := time.Now()
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		wantErr bool
	}{
		{
			name:    "GetSnippet() with id that is no in the database",
			prepare: func() {},
			args:    args{ctx: ctx, id: "MOCK_ID"},
			wantErr: true,
		},
		{
			name: "GetSnippet() in the usual case",
			prepare: func() {
				_ = datastoreDb.PutSnippet(ctx, "MOCK_ID", &entity.Snippet{
					IDMeta: &entity.IDMeta{
						Salt:     "MOCK_SALT",
						IdLength: 11,
					},
					Snippet: &entity.SnippetEntity{
						Sdk:           utils.GetSdkKey(pb.Sdk_SDK_GO.String()),
						PipeOpts:      "MOCK_OPTIONS",
						Created:       nowDate,
						Origin:        constants.UserSnippetOrigin,
						NumberOfFiles: 1,
					},
					Files: []*entity.FileEntity{{
						Name:    "MOCK_NAME",
						Content: "MOCK_CONTENT",
						IsMain:  false,
					}},
				})
			},
			args:    args{ctx: ctx, id: "MOCK_ID"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			snip, err := datastoreDb.GetSnippet(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSnippet() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if snip.Sdk.Name != pb.Sdk_SDK_GO.String() ||
					snip.PipeOpts != "MOCK_OPTIONS" ||
					snip.Origin != constants.UserSnippetOrigin ||
					snip.OwnerId != "" {
					t.Error("GetSnippet() unexpected result")
				}
			}
		})
	}

	cleanData(t, constants.FileKind, "MOCK_ID_0", nil)
	cleanData(t, constants.SnippetKind, "MOCK_ID", nil)
}

func TestDatastore_PutSDKs(t *testing.T) {
	type args struct {
		ctx  context.Context
		sdks []*entity.SDKEntity
	}
	sdks := getSDKs()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "PutSDKs() in the usual case",
			args: args{
				ctx:  ctx,
				sdks: sdks,
			},
			wantErr: false,
		},
		{
			name: "PutSDKs() when input data is nil",
			args: args{
				ctx:  ctx,
				sdks: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := datastoreDb.PutSDKs(tt.args.ctx, tt.args.sdks)
			if err != nil {
				t.Error("PutSDKs() method failed")
			}
		})
	}

	for _, sdk := range sdks {
		cleanData(t, constants.SdkKind, sdk.Name, nil)
	}
}

func TestDatastore_PutSchemaVersion(t *testing.T) {
	type args struct {
		ctx    context.Context
		id     string
		schema *entity.SchemaEntity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "PutSchemaVersion() in the usual case",
			args: args{
				ctx:    ctx,
				id:     "MOCK_ID",
				schema: &entity.SchemaEntity{Descr: "MOCK_DESCRIPTION"},
			},
			wantErr: false,
		},
		{
			name: "PutSchemaVersion() when input data is nil",
			args: args{
				ctx:    ctx,
				id:     "MOCK_ID",
				schema: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := datastoreDb.PutSchemaVersion(tt.args.ctx, tt.args.id, tt.args.schema)
			if err != nil {
				t.Error("PutSchemaVersion() method failed")
			}
		})
	}

	cleanData(t, constants.SchemaKind, "MOCK_ID", nil)
}

func TestDatastore_GetFiles(t *testing.T) {
	type args struct {
		ctx           context.Context
		snipId        string
		numberOfFiles int
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		wantErr bool
	}{
		{
			name:    "GetFiles() with snippet id that is no in the database",
			prepare: func() {},
			args:    args{ctx: ctx, snipId: "MOCK_ID", numberOfFiles: 1},
			wantErr: true,
		},
		{
			name:    "GetFiles() in the usual case",
			prepare: func() { saveSnippet("MOCK_ID", pb.Sdk_SDK_GO.String()) },
			args:    args{ctx: ctx, snipId: "MOCK_ID", numberOfFiles: 1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			files, err := datastoreDb.GetFiles(tt.args.ctx, tt.args.snipId, tt.args.numberOfFiles)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
			if files != nil {
				if len(files) != 1 ||
					files[0].Content != "MOCK_CONTENT" ||
					files[0].IsMain != true {
					t.Error("GetFiles() unexpected result")
				}
				cleanData(t, constants.FileKind, "MOCK_ID_0", nil)
				cleanData(t, constants.SnippetKind, "MOCK_ID", nil)
			}
		})
	}
}

func TestDatastore_GetSDKs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	sdks := getSDKs()
	tests := []struct {
		name    string
		prepare func()
		args    args
		wantErr bool
	}{
		{
			name:    "GetSDKs() in the usual case",
			prepare: func() { _ = datastoreDb.PutSDKs(ctx, sdks) },
			args:    args{ctx: ctx},
			wantErr: false,
		},
		{
			name:    "GetSDKs() when sdks are missing",
			prepare: func() {},
			args:    args{ctx: ctx},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			sdkEntities, err := datastoreDb.GetSDKs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSDKs() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if len(sdkEntities) != 4 {
					t.Error("GetSDK unexpected result, should be four entities")
				}
				for _, sdk := range sdks {
					cleanData(t, constants.SdkKind, sdk.Name, nil)
				}
			}
		})
	}
}

func TestDatastore_GetCatalog(t *testing.T) {
	type args struct {
		ctx        context.Context
		sdkCatalog []*entity.SDKEntity
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		wantErr bool
	}{
		{
			name: "Getting catalog in the usual case",
			prepare: func() {
				saveExample("MOCK_EXAMPLE", pb.Sdk_SDK_JAVA.String())
				saveSnippet("SDK_JAVA_MOCK_EXAMPLE", pb.Sdk_SDK_JAVA.String())
				savePCObjs("SDK_JAVA_MOCK_EXAMPLE")
			},
			args: args{
				ctx: ctx,
				sdkCatalog: func() []*entity.SDKEntity {
					var sdks []*entity.SDKEntity
					for sdkName := range pb.Sdk_value {
						sdks = append(sdks, &entity.SDKEntity{
							Name:           sdkName,
							DefaultExample: "MOCK_DEFAULT_EXAMPLE",
						})
					}
					return sdks
				}(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			catalog, err := datastoreDb.GetCatalog(tt.args.ctx, tt.args.sdkCatalog)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCatalog() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if catalog[0].GetSdk() != pb.Sdk_SDK_JAVA {
					t.Error("GetCatalog() unexpected result: wrong sdk")
				}
				actualCatName := catalog[0].GetCategories()[0].GetCategoryName()
				actualPCObj := catalog[0].GetCategories()[0].GetPrecompiledObjects()[0]
				if actualCatName != "MOCK_CATEGORY" {
					t.Error("GetCatalog() unexpected result: wrong category")
				}
				if actualPCObj.DefaultExample != false ||
					actualPCObj.Multifile != false ||
					actualPCObj.Name != "MOCK_EXAMPLE" ||
					actualPCObj.Type.String() != "PRECOMPILED_OBJECT_TYPE_EXAMPLE" ||
					actualPCObj.CloudPath != "SDK_JAVA/PRECOMPILED_OBJECT_TYPE_EXAMPLE/MOCK_EXAMPLE" ||
					actualPCObj.PipelineOptions != "MOCK_OPTIONS" ||
					actualPCObj.Description != "MOCK_DESCR" ||
					actualPCObj.Link != "MOCK_PATH" ||
					actualPCObj.ContextLine != 32 {
					t.Error("GetCatalog() unexpected result: wrong precompiled obj")
				}
				cleanPCObjs(t, "SDK_JAVA_MOCK_EXAMPLE")
				cleanFiles(t, "SDK_JAVA_MOCK_EXAMPLE", 1)
				cleanData(t, constants.SnippetKind, "SDK_JAVA_MOCK_EXAMPLE", nil)
				cleanData(t, constants.ExampleKind, "SDK_JAVA_MOCK_EXAMPLE", nil)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ctx       context.Context
		projectId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Initialize datastore database",
			args:    args{ctx: ctx, projectId: datastoreEmulatorProjectId},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(ctx, mapper.NewPrecompiledObjectMapper(), datastoreEmulatorProjectId)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func saveExample(name, sdk string) {
	_, _ = datastoreDb.Client.Put(ctx, utils.GetExampleKey(sdk, name), &entity.ExampleEntity{
		Name:       name,
		Sdk:        utils.GetSdkKey(sdk),
		Descr:      "MOCK_DESCR",
		Cats:       []string{"MOCK_CATEGORY"},
		Complexity: "MEDIUM",
		Path:       "MOCK_PATH",
		Type:       "PRECOMPILED_OBJECT_TYPE_EXAMPLE",
		Origin:     constants.ExampleOrigin,
		SchVer:     utils.GetSchemaVerKey("MOCK_VERSION"),
	})
}

func saveSnippet(snipId, sdk string) {
	_ = datastoreDb.PutSnippet(ctx, snipId, &entity.Snippet{
		IDMeta: &entity.IDMeta{
			Salt:     "MOCK_SALT",
			IdLength: 11,
		},
		Snippet: &entity.SnippetEntity{
			Sdk:           utils.GetSdkKey(sdk),
			PipeOpts:      "MOCK_OPTIONS",
			Origin:        constants.ExampleOrigin,
			NumberOfFiles: 1,
		},
		Files: []*entity.FileEntity{{
			Name:     "MOCK_NAME",
			Content:  "MOCK_CONTENT",
			CntxLine: 32,
			IsMain:   true,
		}},
	})
}

func savePCObjs(exampleId string) {
	pcTypes := []string{constants.PCOutputType, constants.PCLogType, constants.PCGraphType}
	for _, pcType := range pcTypes {
		_, _ = datastoreDb.Client.Put(
			ctx,
			utils.GetPCObjectKey(fmt.Sprintf("%s_%s", exampleId, pcType)),
			&entity.PrecompiledObjectEntity{Content: "MOCK_CONTENT_" + pcType})
	}
}

func cleanPCObjs(t *testing.T, exampleId string) {
	pcTypes := []string{constants.PCOutputType, constants.PCLogType, constants.PCGraphType}
	for _, pcType := range pcTypes {
		cleanData(t, constants.PCObjectKind, utils.GetIDWithDelimiter(exampleId, pcType), nil)
	}
}

func cleanFiles(t *testing.T, exampleId string, numberOfFiles int) {
	for fileIndx := 0; fileIndx < numberOfFiles; fileIndx++ {
		cleanData(t, constants.FileKind, utils.GetIDWithDelimiter(exampleId, fileIndx), nil)
	}
}

func cleanData(t *testing.T, kind, id string, parentId *datastore.Key) {
	key := datastore.NameKey(kind, id, nil)
	if parentId != nil {
		key.Parent = parentId
	}
	key.Namespace = constants.Namespace
	if err := datastoreDb.Client.Delete(ctx, key); err != nil {
		t.Errorf("Error during data cleaning after the test, err: %s", err.Error())
	}
}

func getSDKs() []*entity.SDKEntity {
	var sdkEntities []*entity.SDKEntity
	for _, sdk := range pb.Sdk_name {
		sdkEntities = append(sdkEntities, &entity.SDKEntity{
			Name:           sdk,
			DefaultExample: "MOCK_DEFAULT_EXAMPLE",
		})
	}
	return sdkEntities
}
