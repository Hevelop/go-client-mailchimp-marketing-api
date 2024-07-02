/*
Mailchimp Marketing API

Testing ActivityFeedAPIService

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

func Test_mailchimpmarketingapi_ActivityFeedAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ActivityFeedAPIService GetActivityFeedChimpChatter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ActivityFeedAPI.GetActivityFeedChimpChatter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
