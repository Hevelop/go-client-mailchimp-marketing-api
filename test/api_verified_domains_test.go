/*
Mailchimp Marketing API

Testing VerifiedDomainsAPIService

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

func Test_mailchimpmarketingapi_VerifiedDomainsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VerifiedDomainsAPIService CreateVerifiedDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VerifiedDomainsAPI.CreateVerifiedDomain(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VerifiedDomainsAPIService DeleteVerifiedDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		httpRes, err := apiClient.VerifiedDomainsAPI.DeleteVerifiedDomain(context.Background(), domainName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VerifiedDomainsAPIService GetVerifiedDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		resp, httpRes, err := apiClient.VerifiedDomainsAPI.GetVerifiedDomain(context.Background(), domainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VerifiedDomainsAPIService GetVerifiedDomains", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VerifiedDomainsAPI.GetVerifiedDomains(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VerifiedDomainsAPIService VerifyDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		resp, httpRes, err := apiClient.VerifiedDomainsAPI.VerifyDomain(context.Background(), domainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}