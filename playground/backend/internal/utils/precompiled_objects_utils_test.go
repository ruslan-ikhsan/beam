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

package utils

import (
	pb "beam.apache.org/playground/backend/internal/api/v1"
	"beam.apache.org/playground/backend/internal/cloud_bucket"
	"reflect"
	"testing"
)

func TestPutPrecompiledObjectsToCategory(t *testing.T) {
	precompiledObjectToAdd := &cloud_bucket.PrecompiledObjects{
		{"TestName", "SDK_JAVA/TestCategory/TestName.java", "TestDescription", pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE, []string{""}, "", "", false},
	}
	type args struct {
		categoryName       string
		precompiledObjects *cloud_bucket.PrecompiledObjects
		sdkCategory        *pb.Categories
	}
	tests := []struct {
		name string
		args args
		want *pb.Categories
	}{
		{
			name: "Test PutPrecompiledObjectsToCategory",
			args: args{
				categoryName:       "TestCategory",
				precompiledObjects: precompiledObjectToAdd,
				sdkCategory: &pb.Categories{
					Sdk:        pb.Sdk_SDK_JAVA,
					Categories: []*pb.Categories_Category{},
				},
			},
			want: &pb.Categories{
				Sdk: pb.Sdk_SDK_JAVA,
				Categories: []*pb.Categories_Category{
					{
						CategoryName: "TestCategory", PrecompiledObjects: []*pb.PrecompiledObject{
							{
								CloudPath:   "SDK_JAVA/TestCategory/TestName.java",
								Name:        "TestName",
								Description: "TestDescription",
								Type:        pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PutPrecompiledObjectsToCategory(tt.args.categoryName, tt.args.precompiledObjects, tt.args.sdkCategory)
			got := tt.args.sdkCategory
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutPrecompiledObjectsToCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPrecompiledObjects(t *testing.T) {
	catalog := []*pb.Categories{
		{
			Sdk: pb.Sdk_SDK_JAVA,
			Categories: []*pb.Categories_Category{
				{
					CategoryName: "TestCategory", PrecompiledObjects: []*pb.PrecompiledObject{
						{
							CloudPath:   "SDK_JAVA/TestCategory/TestName.java",
							Name:        "TestName",
							Description: "TestDescription",
							Type:        pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
						},
					},
				},
				{
					CategoryName: "AnotherTestCategory", PrecompiledObjects: []*pb.PrecompiledObject{
						{
							CloudPath:   "SDK_JAVA/AnotherTestCategory/TestName.java",
							Name:        "TestName",
							Description: "TestDescription",
							Type:        pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
						},
					},
				},
			},
		},
		{
			Sdk: pb.Sdk_SDK_PYTHON,
			Categories: []*pb.Categories_Category{
				{
					CategoryName: "TestCategory", PrecompiledObjects: []*pb.PrecompiledObject{
						{
							CloudPath:   "SDK_PYTHON/TestCategory/TestName.java",
							Name:        "TestName",
							Description: "TestDescription",
							Type:        pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
						},
					},
				},
			},
		},
	}
	catalogWithSpecificCategory := []*pb.Categories{
		{
			Sdk: pb.Sdk_SDK_JAVA,
			Categories: []*pb.Categories_Category{
				{
					CategoryName: "TestCategory", PrecompiledObjects: []*pb.PrecompiledObject{
						{
							CloudPath:   "SDK_JAVA/TestCategory/TestName.java",
							Name:        "TestName",
							Description: "TestDescription",
							Type:        pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
						},
					},
				},
			},
		},
		{
			Sdk: pb.Sdk_SDK_PYTHON,
			Categories: []*pb.Categories_Category{
				{
					CategoryName: "TestCategory", PrecompiledObjects: []*pb.PrecompiledObject{
						{
							CloudPath:   "SDK_PYTHON/TestCategory/TestName.java",
							Name:        "TestName",
							Description: "TestDescription",
							Type:        pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
						},
					},
				},
			},
		},
	}
	type args struct {
		catalog      []*pb.Categories
		sdk          pb.Sdk
		categoryName string
	}
	tests := []struct {
		name string
		args args
		want []*pb.Categories
	}{
		{
			name: "All catalog",
			args: args{
				catalog:      catalog,
				sdk:          pb.Sdk_SDK_UNSPECIFIED,
				categoryName: "",
			},
			want: catalog,
		},
		{
			name: "Specific SDK",
			args: args{
				catalog:      catalog,
				sdk:          pb.Sdk_SDK_JAVA,
				categoryName: "",
			},
			want: catalog[:1],
		},
		{
			name: "Specific Category",
			args: args{
				catalog:      catalog,
				sdk:          pb.Sdk_SDK_UNSPECIFIED,
				categoryName: "TestCategory",
			},
			want: catalogWithSpecificCategory,
		},
		{
			name: "Specific SDK and Category",
			args: args{
				catalog:      catalog,
				sdk:          pb.Sdk_SDK_JAVA,
				categoryName: "TestCategory",
			},
			want: catalogWithSpecificCategory[:1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterCatalog(tt.args.catalog, tt.args.sdk, tt.args.categoryName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterCatalog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultPrecompiledObjects(t *testing.T) {
	preparedPrecompiledObjectJava := pb.PrecompiledObject{
		CloudPath:      "SDK_JAVA/TestCategory/TestName1.java",
		Name:           "TestName1",
		Description:    "TestDescription",
		Type:           pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
		DefaultExample: true,
	}
	preparedPrecompiledObjectGo := pb.PrecompiledObject{
		CloudPath:      "SDK_GO/TestCategory/TestName.go",
		Name:           "TestName",
		Description:    "TestDescription",
		Type:           pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
		DefaultExample: true,
	}
	sdkCategories := []*pb.Categories{
		{
			Sdk: pb.Sdk_SDK_JAVA,
			Categories: []*pb.Categories_Category{
				{
					CategoryName: "TestCategory1", PrecompiledObjects: []*pb.PrecompiledObject{
					{
						CloudPath:      "SDK_JAVA/TestCategory/TestName.java",
						Name:           "TestName",
						Description:    "TestDescription",
						Type:           pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
						DefaultExample: false,
					},
				},
				},
				{
					CategoryName: "TestCategory2", PrecompiledObjects: []*pb.PrecompiledObject{
					&preparedPrecompiledObjectJava,
					{
						CloudPath:      "SDK_JAVA/TestCategory/TestName2.java",
						Name:           "TestName2",
						Description:    "TestDescription",
						Type:           pb.PrecompiledObjectType_PRECOMPILED_OBJECT_TYPE_EXAMPLE,
						DefaultExample: false,
					},
				},
				},
			},
		},
		{
			Sdk: pb.Sdk_SDK_GO,
			Categories: []*pb.Categories_Category{
				{
					CategoryName: "TestCategory", PrecompiledObjects: []*pb.PrecompiledObject{
					&preparedPrecompiledObjectGo,
				},
				},
			},
		},
	}
	expectedDefaultPrecompiledObjects := make(map[pb.Sdk]*pb.PrecompiledObject)
	expectedDefaultPrecompiledObjects[pb.Sdk_SDK_JAVA] = &preparedPrecompiledObjectJava
	expectedDefaultPrecompiledObjects[pb.Sdk_SDK_GO] = &preparedPrecompiledObjectGo
	type args struct {
		sdkCategories []*pb.Categories
		sdk           pb.Sdk
	}
	tests := []struct {
		name string
		args args
		want *pb.PrecompiledObject
	}{
		{
			// Test case with getting default Precompiled Objects from the precompiled objects catalog
			// As a result, want to receive an expected java default Precompiled Objects.
			name: "get java default precompiled objects",
			args: args{sdkCategories: sdkCategories, sdk: pb.Sdk_SDK_JAVA},
			want: &preparedPrecompiledObjectJava,
		},
		{
			// Test case with getting default Precompiled Objects from the precompiled objects catalog
			// As a result, want to receive an expected go default Precompiled Objects.
			name: "get java default precompiled objects",
			args: args{sdkCategories: sdkCategories, sdk: pb.Sdk_SDK_GO},
			want: &preparedPrecompiledObjectGo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefaultPrecompiledObject(tt.args.sdkCategories, tt.args.sdk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDefaultPrecompiledObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
