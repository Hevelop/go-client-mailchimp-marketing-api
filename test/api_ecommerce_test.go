/*
Mailchimp Marketing API

Testing EcommerceAPIService

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

func Test_mailchimpmarketingapi_EcommerceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresId(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdCartsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var cartId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdCartsId(context.Background(), storeId, cartId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdCartsLinesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var cartId string
		var lineId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdCartsLinesId(context.Background(), storeId, cartId, lineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdCustomersId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var customerId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdCustomersId(context.Background(), storeId, customerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdOrdersId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var orderId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdOrdersId(context.Background(), storeId, orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdOrdersIdLinesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var orderId string
		var lineId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdOrdersIdLinesId(context.Background(), storeId, orderId, lineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdProductsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdProductsId(context.Background(), storeId, productId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdProductsIdImagesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string
		var imageId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdProductsIdImagesId(context.Background(), storeId, productId, imageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdProductsIdVariantsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string
		var variantId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdProductsIdVariantsId(context.Background(), storeId, productId, variantId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdPromocodesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var promoRuleId string
		var promoCodeId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdPromocodesId(context.Background(), storeId, promoRuleId, promoCodeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService DeleteEcommerceStoresIdPromorulesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var promoRuleId string

		httpRes, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdPromorulesId(context.Background(), storeId, promoRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceOrders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStores", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStores(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresId(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdCarts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCarts(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdCartsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var cartId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCartsId(context.Background(), storeId, cartId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdCartsIdLines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var cartId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCartsIdLines(context.Background(), storeId, cartId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdCartsIdLinesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var cartId string
		var lineId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCartsIdLinesId(context.Background(), storeId, cartId, lineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdCustomers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCustomers(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdCustomersId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var customerId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCustomersId(context.Background(), storeId, customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdOrders(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdOrdersId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var orderId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdOrdersId(context.Background(), storeId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdOrdersIdLines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var orderId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdOrdersIdLines(context.Background(), storeId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdOrdersIdLinesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var orderId string
		var lineId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdOrdersIdLinesId(context.Background(), storeId, orderId, lineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdProducts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProducts(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdProductsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsId(context.Background(), storeId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdProductsIdImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsIdImages(context.Background(), storeId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdProductsIdImagesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string
		var imageId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsIdImagesId(context.Background(), storeId, productId, imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdProductsIdVariants", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsIdVariants(context.Background(), storeId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdProductsIdVariantsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string
		var variantId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsIdVariantsId(context.Background(), storeId, productId, variantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdPromocodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var promoRuleId string
		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdPromocodes(context.Background(), promoRuleId, storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdPromocodesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var promoRuleId string
		var promoCodeId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdPromocodesId(context.Background(), storeId, promoRuleId, promoCodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdPromorules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdPromorules(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService GetEcommerceStoresIdPromorulesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var promoRuleId string

		resp, httpRes, err := apiClient.EcommerceAPI.GetEcommerceStoresIdPromorulesId(context.Background(), storeId, promoRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresId(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdCartsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var cartId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdCartsId(context.Background(), storeId, cartId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdCartsIdLinesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var cartId string
		var lineId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdCartsIdLinesId(context.Background(), storeId, cartId, lineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdCustomersId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var customerId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdCustomersId(context.Background(), storeId, customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdOrdersId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var orderId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdOrdersId(context.Background(), storeId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdOrdersIdLinesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var orderId string
		var lineId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdOrdersIdLinesId(context.Background(), storeId, orderId, lineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdProductsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdProductsId(context.Background(), storeId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdProductsIdImagesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string
		var imageId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdProductsIdImagesId(context.Background(), storeId, productId, imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdProductsIdVariantsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string
		var variantId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdProductsIdVariantsId(context.Background(), storeId, productId, variantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdPromocodesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var promoRuleId string
		var promoCodeId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdPromocodesId(context.Background(), storeId, promoRuleId, promoCodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PatchEcommerceStoresIdPromorulesId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var promoRuleId string

		resp, httpRes, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdPromorulesId(context.Background(), storeId, promoRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStores", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStores(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdCarts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdCarts(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdCartsIdLines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var cartId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdCartsIdLines(context.Background(), storeId, cartId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdCustomers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdCustomers(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdOrders(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdOrdersIdLines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var orderId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdOrdersIdLines(context.Background(), storeId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdProducts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdProducts(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdProductsIdImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdProductsIdImages(context.Background(), storeId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdProductsIdVariants", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdProductsIdVariants(context.Background(), storeId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdPromocodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var promoRuleId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdPromocodes(context.Background(), storeId, promoRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PostEcommerceStoresIdPromorules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string

		resp, httpRes, err := apiClient.EcommerceAPI.PostEcommerceStoresIdPromorules(context.Background(), storeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PutEcommerceStoresIdCustomersId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var customerId string

		resp, httpRes, err := apiClient.EcommerceAPI.PutEcommerceStoresIdCustomersId(context.Background(), storeId, customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EcommerceAPIService PutEcommerceStoresIdProductsIdVariantsId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var storeId string
		var productId string
		var variantId string

		resp, httpRes, err := apiClient.EcommerceAPI.PutEcommerceStoresIdProductsIdVariantsId(context.Background(), storeId, productId, variantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}