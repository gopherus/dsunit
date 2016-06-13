/*
 *
 *
 * Copyright 2012-2016 Viant.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 *  use this file except in compliance with the License. You may obtain a copy of
 *  the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 *  License for the specific language governing permissions and limitations under
 *  the License.
 *
 */
package dsunit_test


import (
	"testing"
	"github.com/viant/dsunit"
	"github.com/stretchr/testify/assert"
	"github.com/viant/dsc"
)



func TestDataset(test *testing.T) {
	var datasetFactory dsunit.DatasetFactory = dsunit.NewDatasetTestManager().DatasetFactory()
	descriptor := &dsc.TableDescriptor{Table:"users",Autoincrement: true, PkColumns:[]string{"id"}}

	dataset:= datasetFactory.Create(descriptor,
		map[string]interface{}{
			"id":1,
			"username":"Dudi",
			"active":true,
			"comments":"abc",
		},
		map[string]interface{}{
			"id":2,
			"username":"Bogi",
			"active":false,
		},

	)

	assert.NotNil(test, dataset, "Should a dataset")
	assert.Equal(test, "users",dataset.Table, "Should a dataset for users table")
	assert.Equal(test, "id", dataset.PkColumns[0], "Should have a dataset with id pkcolumn")
	assert.Equal(test, true, dataset.Autoincrement,"Should have a dataset with autoincrement")
	assert.Equal(test, 2, len(dataset.Rows),"Should have a dataset with 2 rows")

	{
		row := dataset.Rows[0];
		assert.Equal(test, 4, len(row.Columns()), "The first row should have 4 columns")
		assert.Equal(test, 1, row.Value("id"), "The first row should have id")
		assert.Equal(test, true, row.Value("active"), "The first row should be active")
		assert.True(test, row.HasColumn("id"), "The first row should have column id")
		assert.True(test, row.HasColumn("comments"), "The first row should have column comments")

	}

	{
		row := dataset.Rows[1];
		assert.Equal(test, 3, len(row.Columns()), "The second row should have 3 columns")
		assert.Equal(test, 2, row.Value("id"), "The second row should have id")
		assert.Equal(test, false, row.Value("active"), "The second row should be inactive")
		assert.True(test, row.HasColumn("id"), "The second row should have column id")
		assert.False(test, row.HasColumn("comments"), "The second row should not have column comments")

	}

}
