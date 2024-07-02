/*
Mailchimp Marketing API

Testing BatchesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mailchimpmarketingapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func Test_mailchimpmarketingapi_BatchesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BatchesAPIService DeleteBatchesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var batchId string

		httpRes, err := apiClient.BatchesAPI.DeleteBatchesId(context.Background(), batchId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchesAPIService GetBatches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchesAPI.GetBatches(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchesAPIService GetBatchesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var batchId string

		resp, httpRes, err := apiClient.BatchesAPI.GetBatchesId(context.Background(), batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchesAPIService PostBatches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchesAPI.PostBatches(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
