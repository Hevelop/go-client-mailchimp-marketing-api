/*
Mailchimp Marketing API

Testing ConversationsAPIService

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

func Test_mailchimpmarketingapi_ConversationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConversationsAPIService GetConversations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConversationsAPI.GetConversations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConversationsAPIService GetConversationsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var conversationId string

		resp, httpRes, err := apiClient.ConversationsAPI.GetConversationsId(context.Background(), conversationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConversationsAPIService GetConversationsIdMessages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var conversationId string

		resp, httpRes, err := apiClient.ConversationsAPI.GetConversationsIdMessages(context.Background(), conversationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConversationsAPIService GetConversationsIdMessagesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var conversationId string
		var messageId string

		resp, httpRes, err := apiClient.ConversationsAPI.GetConversationsIdMessagesId(context.Background(), conversationId, messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
