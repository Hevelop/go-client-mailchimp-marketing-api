/*
Mailchimp Marketing API

Testing CustomerJourneysAPIService

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

func Test_mailchimpmarketingapi_CustomerJourneysAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomerJourneysAPIService PostCustomerJourneysJourneysIdStepsIdActionsTrigger", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var journeyId int32
		var stepId int32

		resp, httpRes, err := apiClient.CustomerJourneysAPI.PostCustomerJourneysJourneysIdStepsIdActionsTrigger(context.Background(), journeyId, stepId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
