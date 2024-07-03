# Hevelop\MailchimpMarketingApi\EcommerceAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEcommerceStoresId**](EcommerceAPI.md#DeleteEcommerceStoresId) | **Delete** /ecommerce/stores/{store_id} | Delete store
[**DeleteEcommerceStoresIdCartsId**](EcommerceAPI.md#DeleteEcommerceStoresIdCartsId) | **Delete** /ecommerce/stores/{store_id}/carts/{cart_id} | Delete cart
[**DeleteEcommerceStoresIdCartsLinesId**](EcommerceAPI.md#DeleteEcommerceStoresIdCartsLinesId) | **Delete** /ecommerce/stores/{store_id}/carts/{cart_id}/lines/{line_id} | Delete cart line item
[**DeleteEcommerceStoresIdCustomersId**](EcommerceAPI.md#DeleteEcommerceStoresIdCustomersId) | **Delete** /ecommerce/stores/{store_id}/customers/{customer_id} | Delete customer
[**DeleteEcommerceStoresIdOrdersId**](EcommerceAPI.md#DeleteEcommerceStoresIdOrdersId) | **Delete** /ecommerce/stores/{store_id}/orders/{order_id} | Delete order
[**DeleteEcommerceStoresIdOrdersIdLinesId**](EcommerceAPI.md#DeleteEcommerceStoresIdOrdersIdLinesId) | **Delete** /ecommerce/stores/{store_id}/orders/{order_id}/lines/{line_id} | Delete order line item
[**DeleteEcommerceStoresIdProductsId**](EcommerceAPI.md#DeleteEcommerceStoresIdProductsId) | **Delete** /ecommerce/stores/{store_id}/products/{product_id} | Delete product
[**DeleteEcommerceStoresIdProductsIdImagesId**](EcommerceAPI.md#DeleteEcommerceStoresIdProductsIdImagesId) | **Delete** /ecommerce/stores/{store_id}/products/{product_id}/images/{image_id} | Delete product image
[**DeleteEcommerceStoresIdProductsIdVariantsId**](EcommerceAPI.md#DeleteEcommerceStoresIdProductsIdVariantsId) | **Delete** /ecommerce/stores/{store_id}/products/{product_id}/variants/{variant_id} | Delete product variant
[**DeleteEcommerceStoresIdPromocodesId**](EcommerceAPI.md#DeleteEcommerceStoresIdPromocodesId) | **Delete** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes/{promo_code_id} | Delete promo code
[**DeleteEcommerceStoresIdPromorulesId**](EcommerceAPI.md#DeleteEcommerceStoresIdPromorulesId) | **Delete** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id} | Delete promo rule
[**GetEcommerceOrders**](EcommerceAPI.md#GetEcommerceOrders) | **Get** /ecommerce/orders | List account orders
[**GetEcommerceStores**](EcommerceAPI.md#GetEcommerceStores) | **Get** /ecommerce/stores | List stores
[**GetEcommerceStoresId**](EcommerceAPI.md#GetEcommerceStoresId) | **Get** /ecommerce/stores/{store_id} | Get store info
[**GetEcommerceStoresIdCarts**](EcommerceAPI.md#GetEcommerceStoresIdCarts) | **Get** /ecommerce/stores/{store_id}/carts | List carts
[**GetEcommerceStoresIdCartsId**](EcommerceAPI.md#GetEcommerceStoresIdCartsId) | **Get** /ecommerce/stores/{store_id}/carts/{cart_id} | Get cart info
[**GetEcommerceStoresIdCartsIdLines**](EcommerceAPI.md#GetEcommerceStoresIdCartsIdLines) | **Get** /ecommerce/stores/{store_id}/carts/{cart_id}/lines | List cart line items
[**GetEcommerceStoresIdCartsIdLinesId**](EcommerceAPI.md#GetEcommerceStoresIdCartsIdLinesId) | **Get** /ecommerce/stores/{store_id}/carts/{cart_id}/lines/{line_id} | Get cart line item
[**GetEcommerceStoresIdCustomers**](EcommerceAPI.md#GetEcommerceStoresIdCustomers) | **Get** /ecommerce/stores/{store_id}/customers | List customers
[**GetEcommerceStoresIdCustomersId**](EcommerceAPI.md#GetEcommerceStoresIdCustomersId) | **Get** /ecommerce/stores/{store_id}/customers/{customer_id} | Get customer info
[**GetEcommerceStoresIdOrders**](EcommerceAPI.md#GetEcommerceStoresIdOrders) | **Get** /ecommerce/stores/{store_id}/orders | List orders
[**GetEcommerceStoresIdOrdersId**](EcommerceAPI.md#GetEcommerceStoresIdOrdersId) | **Get** /ecommerce/stores/{store_id}/orders/{order_id} | Get order info
[**GetEcommerceStoresIdOrdersIdLines**](EcommerceAPI.md#GetEcommerceStoresIdOrdersIdLines) | **Get** /ecommerce/stores/{store_id}/orders/{order_id}/lines | List order line items
[**GetEcommerceStoresIdOrdersIdLinesId**](EcommerceAPI.md#GetEcommerceStoresIdOrdersIdLinesId) | **Get** /ecommerce/stores/{store_id}/orders/{order_id}/lines/{line_id} | Get order line item
[**GetEcommerceStoresIdProducts**](EcommerceAPI.md#GetEcommerceStoresIdProducts) | **Get** /ecommerce/stores/{store_id}/products | List product
[**GetEcommerceStoresIdProductsId**](EcommerceAPI.md#GetEcommerceStoresIdProductsId) | **Get** /ecommerce/stores/{store_id}/products/{product_id} | Get product info
[**GetEcommerceStoresIdProductsIdImages**](EcommerceAPI.md#GetEcommerceStoresIdProductsIdImages) | **Get** /ecommerce/stores/{store_id}/products/{product_id}/images | List product images
[**GetEcommerceStoresIdProductsIdImagesId**](EcommerceAPI.md#GetEcommerceStoresIdProductsIdImagesId) | **Get** /ecommerce/stores/{store_id}/products/{product_id}/images/{image_id} | Get product image info
[**GetEcommerceStoresIdProductsIdVariants**](EcommerceAPI.md#GetEcommerceStoresIdProductsIdVariants) | **Get** /ecommerce/stores/{store_id}/products/{product_id}/variants | List product variants
[**GetEcommerceStoresIdProductsIdVariantsId**](EcommerceAPI.md#GetEcommerceStoresIdProductsIdVariantsId) | **Get** /ecommerce/stores/{store_id}/products/{product_id}/variants/{variant_id} | Get product variant info
[**GetEcommerceStoresIdPromocodes**](EcommerceAPI.md#GetEcommerceStoresIdPromocodes) | **Get** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes | List promo codes
[**GetEcommerceStoresIdPromocodesId**](EcommerceAPI.md#GetEcommerceStoresIdPromocodesId) | **Get** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes/{promo_code_id} | Get promo code
[**GetEcommerceStoresIdPromorules**](EcommerceAPI.md#GetEcommerceStoresIdPromorules) | **Get** /ecommerce/stores/{store_id}/promo-rules | List promo rules
[**GetEcommerceStoresIdPromorulesId**](EcommerceAPI.md#GetEcommerceStoresIdPromorulesId) | **Get** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id} | Get promo rule
[**PatchEcommerceStoresId**](EcommerceAPI.md#PatchEcommerceStoresId) | **Patch** /ecommerce/stores/{store_id} | Update store
[**PatchEcommerceStoresIdCartsId**](EcommerceAPI.md#PatchEcommerceStoresIdCartsId) | **Patch** /ecommerce/stores/{store_id}/carts/{cart_id} | Update cart
[**PatchEcommerceStoresIdCartsIdLinesId**](EcommerceAPI.md#PatchEcommerceStoresIdCartsIdLinesId) | **Patch** /ecommerce/stores/{store_id}/carts/{cart_id}/lines/{line_id} | Update cart line item
[**PatchEcommerceStoresIdCustomersId**](EcommerceAPI.md#PatchEcommerceStoresIdCustomersId) | **Patch** /ecommerce/stores/{store_id}/customers/{customer_id} | Update customer
[**PatchEcommerceStoresIdOrdersId**](EcommerceAPI.md#PatchEcommerceStoresIdOrdersId) | **Patch** /ecommerce/stores/{store_id}/orders/{order_id} | Update order
[**PatchEcommerceStoresIdOrdersIdLinesId**](EcommerceAPI.md#PatchEcommerceStoresIdOrdersIdLinesId) | **Patch** /ecommerce/stores/{store_id}/orders/{order_id}/lines/{line_id} | Update order line item
[**PatchEcommerceStoresIdProductsId**](EcommerceAPI.md#PatchEcommerceStoresIdProductsId) | **Patch** /ecommerce/stores/{store_id}/products/{product_id} | Update product
[**PatchEcommerceStoresIdProductsIdImagesId**](EcommerceAPI.md#PatchEcommerceStoresIdProductsIdImagesId) | **Patch** /ecommerce/stores/{store_id}/products/{product_id}/images/{image_id} | Update product image
[**PatchEcommerceStoresIdProductsIdVariantsId**](EcommerceAPI.md#PatchEcommerceStoresIdProductsIdVariantsId) | **Patch** /ecommerce/stores/{store_id}/products/{product_id}/variants/{variant_id} | Update product variant
[**PatchEcommerceStoresIdPromocodesId**](EcommerceAPI.md#PatchEcommerceStoresIdPromocodesId) | **Patch** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes/{promo_code_id} | Update promo code
[**PatchEcommerceStoresIdPromorulesId**](EcommerceAPI.md#PatchEcommerceStoresIdPromorulesId) | **Patch** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id} | Update promo rule
[**PostEcommerceStores**](EcommerceAPI.md#PostEcommerceStores) | **Post** /ecommerce/stores | Add store
[**PostEcommerceStoresIdCarts**](EcommerceAPI.md#PostEcommerceStoresIdCarts) | **Post** /ecommerce/stores/{store_id}/carts | Add cart
[**PostEcommerceStoresIdCartsIdLines**](EcommerceAPI.md#PostEcommerceStoresIdCartsIdLines) | **Post** /ecommerce/stores/{store_id}/carts/{cart_id}/lines | Add cart line item
[**PostEcommerceStoresIdCustomers**](EcommerceAPI.md#PostEcommerceStoresIdCustomers) | **Post** /ecommerce/stores/{store_id}/customers | Add customer
[**PostEcommerceStoresIdOrders**](EcommerceAPI.md#PostEcommerceStoresIdOrders) | **Post** /ecommerce/stores/{store_id}/orders | Add order
[**PostEcommerceStoresIdOrdersIdLines**](EcommerceAPI.md#PostEcommerceStoresIdOrdersIdLines) | **Post** /ecommerce/stores/{store_id}/orders/{order_id}/lines | Add order line item
[**PostEcommerceStoresIdProducts**](EcommerceAPI.md#PostEcommerceStoresIdProducts) | **Post** /ecommerce/stores/{store_id}/products | Add product
[**PostEcommerceStoresIdProductsIdImages**](EcommerceAPI.md#PostEcommerceStoresIdProductsIdImages) | **Post** /ecommerce/stores/{store_id}/products/{product_id}/images | Add product image
[**PostEcommerceStoresIdProductsIdVariants**](EcommerceAPI.md#PostEcommerceStoresIdProductsIdVariants) | **Post** /ecommerce/stores/{store_id}/products/{product_id}/variants | Add product variant
[**PostEcommerceStoresIdPromocodes**](EcommerceAPI.md#PostEcommerceStoresIdPromocodes) | **Post** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes | Add promo code
[**PostEcommerceStoresIdPromorules**](EcommerceAPI.md#PostEcommerceStoresIdPromorules) | **Post** /ecommerce/stores/{store_id}/promo-rules | Add promo rule
[**PutEcommerceStoresIdCustomersId**](EcommerceAPI.md#PutEcommerceStoresIdCustomersId) | **Put** /ecommerce/stores/{store_id}/customers/{customer_id} | Add or update customer
[**PutEcommerceStoresIdProductsIdVariantsId**](EcommerceAPI.md#PutEcommerceStoresIdProductsIdVariantsId) | **Put** /ecommerce/stores/{store_id}/products/{product_id}/variants/{variant_id} | Add or update product variant



## DeleteEcommerceStoresId

> map[string]interface{} DeleteEcommerceStoresId(ctx, storeId).Execute()

Delete store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresId(context.Background(), storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEcommerceStoresId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.DeleteEcommerceStoresId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdCartsId

> DeleteEcommerceStoresIdCartsId(ctx, storeId, cartId).Execute()

Delete cart



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	cartId := "cartId_example" // string | The id for the cart.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdCartsId(context.Background(), storeId, cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdCartsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**cartId** | **string** | The id for the cart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdCartsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdCartsLinesId

> DeleteEcommerceStoresIdCartsLinesId(ctx, storeId, cartId, lineId).Execute()

Delete cart line item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	cartId := "cartId_example" // string | The id for the cart.
	lineId := "lineId_example" // string | The id for the line item of a cart.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdCartsLinesId(context.Background(), storeId, cartId, lineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdCartsLinesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**cartId** | **string** | The id for the cart. | 
**lineId** | **string** | The id for the line item of a cart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdCartsLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdCustomersId

> DeleteEcommerceStoresIdCustomersId(ctx, storeId, customerId).Execute()

Delete customer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	customerId := "customerId_example" // string | The id for the customer of a store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdCustomersId(context.Background(), storeId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdCustomersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**customerId** | **string** | The id for the customer of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdCustomersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdOrdersId

> DeleteEcommerceStoresIdOrdersId(ctx, storeId, orderId).Execute()

Delete order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	orderId := "orderId_example" // string | The id for the order in a store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdOrdersId(context.Background(), storeId, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdOrdersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**orderId** | **string** | The id for the order in a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdOrdersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdOrdersIdLinesId

> DeleteEcommerceStoresIdOrdersIdLinesId(ctx, storeId, orderId, lineId).Execute()

Delete order line item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	orderId := "orderId_example" // string | The id for the order in a store.
	lineId := "lineId_example" // string | The id for the line item of an order.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdOrdersIdLinesId(context.Background(), storeId, orderId, lineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdOrdersIdLinesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**orderId** | **string** | The id for the order in a store. | 
**lineId** | **string** | The id for the line item of an order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdOrdersIdLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdProductsId

> DeleteEcommerceStoresIdProductsId(ctx, storeId, productId).Execute()

Delete product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdProductsId(context.Background(), storeId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdProductsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdProductsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdProductsIdImagesId

> DeleteEcommerceStoresIdProductsIdImagesId(ctx, storeId, productId, imageId).Execute()

Delete product image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	imageId := "imageId_example" // string | The id for the product image.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdProductsIdImagesId(context.Background(), storeId, productId, imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdProductsIdImagesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 
**imageId** | **string** | The id for the product image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdProductsIdImagesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdProductsIdVariantsId

> DeleteEcommerceStoresIdProductsIdVariantsId(ctx, storeId, productId, variantId).Execute()

Delete product variant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	variantId := "variantId_example" // string | The id for the product variant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdProductsIdVariantsId(context.Background(), storeId, productId, variantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdProductsIdVariantsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 
**variantId** | **string** | The id for the product variant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdProductsIdVariantsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdPromocodesId

> DeleteEcommerceStoresIdPromocodesId(ctx, storeId, promoRuleId, promoCodeId).Execute()

Delete promo code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	promoRuleId := "promoRuleId_example" // string | The id for the promo rule of a store.
	promoCodeId := "promoCodeId_example" // string | The id for the promo code of a store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdPromocodesId(context.Background(), storeId, promoRuleId, promoCodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdPromocodesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**promoRuleId** | **string** | The id for the promo rule of a store. | 
**promoCodeId** | **string** | The id for the promo code of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdPromocodesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcommerceStoresIdPromorulesId

> DeleteEcommerceStoresIdPromorulesId(ctx, storeId, promoRuleId).Execute()

Delete promo rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	promoRuleId := "promoRuleId_example" // string | The id for the promo rule of a store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EcommerceAPI.DeleteEcommerceStoresIdPromorulesId(context.Background(), storeId, promoRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.DeleteEcommerceStoresIdPromorulesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**promoRuleId** | **string** | The id for the promo rule of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcommerceStoresIdPromorulesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceOrders

> Orders GetEcommerceOrders(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).CampaignId(campaignId).OutreachId(outreachId).CustomerId(customerId).HasOutreach(hasOutreach).Execute()

List account orders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	campaignId := "campaignId_example" // string | Restrict results to orders with a specific `campaign_id` value. (optional)
	outreachId := "outreachId_example" // string | Restrict results to orders with a specific `outreach_id` value. (optional)
	customerId := "customerId_example" // string | Restrict results to orders made by a specific customer. (optional)
	hasOutreach := true // bool | Restrict results to orders that have an outreach attached. For example, an email campaign or Facebook ad. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceOrders(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).CampaignId(campaignId).OutreachId(outreachId).CustomerId(customerId).HasOutreach(hasOutreach).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceOrders`: Orders
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **campaignId** | **string** | Restrict results to orders with a specific &#x60;campaign_id&#x60; value. | 
 **outreachId** | **string** | Restrict results to orders with a specific &#x60;outreach_id&#x60; value. | 
 **customerId** | **string** | Restrict results to orders made by a specific customer. | 
 **hasOutreach** | **bool** | Restrict results to orders that have an outreach attached. For example, an email campaign or Facebook ad. | 

### Return type

[**Orders**](Orders.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStores

> ECommerceStores GetEcommerceStores(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List stores



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStores(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStores`: ECommerceStores
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**ECommerceStores**](ECommerceStores.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresId

> ECommerceStore GetEcommerceStoresId(ctx, storeId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get store info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresId(context.Background(), storeId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresId`: ECommerceStore
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceStore**](ECommerceStore.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdCarts

> Carts GetEcommerceStoresIdCarts(ctx, storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List carts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCarts(context.Background(), storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdCarts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdCarts`: Carts
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdCarts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdCartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**Carts**](Carts.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdCartsId

> ECommerceCart GetEcommerceStoresIdCartsId(ctx, storeId, cartId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get cart info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	cartId := "cartId_example" // string | The id for the cart.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCartsId(context.Background(), storeId, cartId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdCartsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdCartsId`: ECommerceCart
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdCartsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**cartId** | **string** | The id for the cart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdCartsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceCart**](ECommerceCart.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdCartsIdLines

> CartLines GetEcommerceStoresIdCartsIdLines(ctx, storeId, cartId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List cart line items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	cartId := "cartId_example" // string | The id for the cart.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCartsIdLines(context.Background(), storeId, cartId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdCartsIdLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdCartsIdLines`: CartLines
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdCartsIdLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**cartId** | **string** | The id for the cart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdCartsIdLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**CartLines**](CartLines.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdCartsIdLinesId

> ECommerceCartLineItem GetEcommerceStoresIdCartsIdLinesId(ctx, storeId, cartId, lineId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get cart line item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	cartId := "cartId_example" // string | The id for the cart.
	lineId := "lineId_example" // string | The id for the line item of a cart.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCartsIdLinesId(context.Background(), storeId, cartId, lineId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdCartsIdLinesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdCartsIdLinesId`: ECommerceCartLineItem
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdCartsIdLinesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**cartId** | **string** | The id for the cart. | 
**lineId** | **string** | The id for the line item of a cart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdCartsIdLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceCartLineItem**](ECommerceCartLineItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdCustomers

> Customers GetEcommerceStoresIdCustomers(ctx, storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).EmailAddress(emailAddress).Execute()

List customers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	emailAddress := "emailAddress_example" // string | Restrict the response to customers with the email address. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCustomers(context.Background(), storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).EmailAddress(emailAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdCustomers`: Customers
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdCustomers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **emailAddress** | **string** | Restrict the response to customers with the email address. | 

### Return type

[**Customers**](Customers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdCustomersId

> ECommerceCustomer GetEcommerceStoresIdCustomersId(ctx, storeId, customerId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get customer info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	customerId := "customerId_example" // string | The id for the customer of a store.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdCustomersId(context.Background(), storeId, customerId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdCustomersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdCustomersId`: ECommerceCustomer
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdCustomersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**customerId** | **string** | The id for the customer of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdCustomersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceCustomer**](ECommerceCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdOrders

> Orders1 GetEcommerceStoresIdOrders(ctx, storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).CustomerId(customerId).HasOutreach(hasOutreach).CampaignId(campaignId).OutreachId(outreachId).Execute()

List orders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	customerId := "customerId_example" // string | Restrict results to orders made by a specific customer. (optional)
	hasOutreach := true // bool | Restrict results to orders that have an outreach attached. For example, an email campaign or Facebook ad. (optional)
	campaignId := "campaignId_example" // string | Restrict results to orders with a specific `campaign_id` value. (optional)
	outreachId := "outreachId_example" // string | Restrict results to orders with a specific `outreach_id` value. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdOrders(context.Background(), storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).CustomerId(customerId).HasOutreach(hasOutreach).CampaignId(campaignId).OutreachId(outreachId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdOrders`: Orders1
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **customerId** | **string** | Restrict results to orders made by a specific customer. | 
 **hasOutreach** | **bool** | Restrict results to orders that have an outreach attached. For example, an email campaign or Facebook ad. | 
 **campaignId** | **string** | Restrict results to orders with a specific &#x60;campaign_id&#x60; value. | 
 **outreachId** | **string** | Restrict results to orders with a specific &#x60;outreach_id&#x60; value. | 

### Return type

[**Orders1**](Orders1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdOrdersId

> ECommerceOrder GetEcommerceStoresIdOrdersId(ctx, storeId, orderId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get order info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	orderId := "orderId_example" // string | The id for the order in a store.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdOrdersId(context.Background(), storeId, orderId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdOrdersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdOrdersId`: ECommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdOrdersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**orderId** | **string** | The id for the order in a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdOrdersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceOrder**](ECommerceOrder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdOrdersIdLines

> OrderLines GetEcommerceStoresIdOrdersIdLines(ctx, storeId, orderId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List order line items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	orderId := "orderId_example" // string | The id for the order in a store.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdOrdersIdLines(context.Background(), storeId, orderId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdOrdersIdLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdOrdersIdLines`: OrderLines
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdOrdersIdLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**orderId** | **string** | The id for the order in a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdOrdersIdLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**OrderLines**](OrderLines.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdOrdersIdLinesId

> ECommerceOrderLineItem GetEcommerceStoresIdOrdersIdLinesId(ctx, storeId, orderId, lineId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get order line item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	orderId := "orderId_example" // string | The id for the order in a store.
	lineId := "lineId_example" // string | The id for the line item of an order.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdOrdersIdLinesId(context.Background(), storeId, orderId, lineId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdOrdersIdLinesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdOrdersIdLinesId`: ECommerceOrderLineItem
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdOrdersIdLinesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**orderId** | **string** | The id for the order in a store. | 
**lineId** | **string** | The id for the line item of an order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdOrdersIdLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceOrderLineItem**](ECommerceOrderLineItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdProducts

> Products GetEcommerceStoresIdProducts(ctx, storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProducts(context.Background(), storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdProducts`: Products
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**Products**](Products.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdProductsId

> ECommerceProduct GetEcommerceStoresIdProductsId(ctx, storeId, productId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get product info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsId(context.Background(), storeId, productId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdProductsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdProductsId`: ECommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdProductsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdProductsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceProduct**](ECommerceProduct.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdProductsIdImages

> EcommerceProductImages GetEcommerceStoresIdProductsIdImages(ctx, storeId, productId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List product images



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsIdImages(context.Background(), storeId, productId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdProductsIdImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdProductsIdImages`: EcommerceProductImages
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdProductsIdImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdProductsIdImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**EcommerceProductImages**](EcommerceProductImages.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdProductsIdImagesId

> ECommerceProductImage GetEcommerceStoresIdProductsIdImagesId(ctx, storeId, productId, imageId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get product image info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	imageId := "imageId_example" // string | The id for the product image.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsIdImagesId(context.Background(), storeId, productId, imageId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdProductsIdImagesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdProductsIdImagesId`: ECommerceProductImage
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdProductsIdImagesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 
**imageId** | **string** | The id for the product image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdProductsIdImagesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceProductImage**](ECommerceProductImage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdProductsIdVariants

> EcommerceProductVariants GetEcommerceStoresIdProductsIdVariants(ctx, storeId, productId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List product variants



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsIdVariants(context.Background(), storeId, productId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdProductsIdVariants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdProductsIdVariants`: EcommerceProductVariants
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdProductsIdVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdProductsIdVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**EcommerceProductVariants**](EcommerceProductVariants.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdProductsIdVariantsId

> ECommerceProductVariant GetEcommerceStoresIdProductsIdVariantsId(ctx, storeId, productId, variantId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get product variant info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	variantId := "variantId_example" // string | The id for the product variant.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdProductsIdVariantsId(context.Background(), storeId, productId, variantId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdProductsIdVariantsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdProductsIdVariantsId`: ECommerceProductVariant
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdProductsIdVariantsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 
**variantId** | **string** | The id for the product variant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdProductsIdVariantsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommerceProductVariant**](ECommerceProductVariant.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdPromocodes

> PromoCodes GetEcommerceStoresIdPromocodes(ctx, promoRuleId, storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List promo codes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	promoRuleId := "promoRuleId_example" // string | The id for the promo rule of a store.
	storeId := "storeId_example" // string | The store id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdPromocodes(context.Background(), promoRuleId, storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdPromocodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdPromocodes`: PromoCodes
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdPromocodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promoRuleId** | **string** | The id for the promo rule of a store. | 
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdPromocodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**PromoCodes**](PromoCodes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdPromocodesId

> ECommercePromoCode GetEcommerceStoresIdPromocodesId(ctx, storeId, promoRuleId, promoCodeId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get promo code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	promoRuleId := "promoRuleId_example" // string | The id for the promo rule of a store.
	promoCodeId := "promoCodeId_example" // string | The id for the promo code of a store.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdPromocodesId(context.Background(), storeId, promoRuleId, promoCodeId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdPromocodesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdPromocodesId`: ECommercePromoCode
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdPromocodesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**promoRuleId** | **string** | The id for the promo rule of a store. | 
**promoCodeId** | **string** | The id for the promo code of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdPromocodesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommercePromoCode**](ECommercePromoCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdPromorules

> PromoRules GetEcommerceStoresIdPromorules(ctx, storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List promo rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdPromorules(context.Background(), storeId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdPromorules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdPromorules`: PromoRules
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdPromorules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdPromorulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**PromoRules**](PromoRules.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcommerceStoresIdPromorulesId

> ECommercePromoRule GetEcommerceStoresIdPromorulesId(ctx, storeId, promoRuleId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get promo rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	promoRuleId := "promoRuleId_example" // string | The id for the promo rule of a store.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.GetEcommerceStoresIdPromorulesId(context.Background(), storeId, promoRuleId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.GetEcommerceStoresIdPromorulesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcommerceStoresIdPromorulesId`: ECommercePromoRule
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.GetEcommerceStoresIdPromorulesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**promoRuleId** | **string** | The id for the promo rule of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcommerceStoresIdPromorulesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ECommercePromoRule**](ECommercePromoRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresId

> ECommerceStore PatchEcommerceStoresId(ctx, storeId).Body(body).Execute()

Update store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	body := *openapiclient.NewECommerceStore2() // ECommerceStore2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresId(context.Background(), storeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresId`: ECommerceStore
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ECommerceStore2**](ECommerceStore2.md) |  | 

### Return type

[**ECommerceStore**](ECommerceStore.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdCartsId

> ECommerceCart PatchEcommerceStoresIdCartsId(ctx, storeId, cartId).Body(body).Execute()

Update cart



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	cartId := "cartId_example" // string | The id for the cart.
	body := *openapiclient.NewECommerceCart2() // ECommerceCart2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdCartsId(context.Background(), storeId, cartId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdCartsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdCartsId`: ECommerceCart
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdCartsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**cartId** | **string** | The id for the cart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdCartsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceCart2**](ECommerceCart2.md) |  | 

### Return type

[**ECommerceCart**](ECommerceCart.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdCartsIdLinesId

> ECommerceCartLineItem PatchEcommerceStoresIdCartsIdLinesId(ctx, storeId, cartId, lineId).Body(body).Execute()

Update cart line item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	cartId := "cartId_example" // string | The id for the cart.
	lineId := "lineId_example" // string | The id for the line item of a cart.
	body := *openapiclient.NewECommerceCartLineItem2() // ECommerceCartLineItem2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdCartsIdLinesId(context.Background(), storeId, cartId, lineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdCartsIdLinesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdCartsIdLinesId`: ECommerceCartLineItem
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdCartsIdLinesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**cartId** | **string** | The id for the cart. | 
**lineId** | **string** | The id for the line item of a cart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdCartsIdLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ECommerceCartLineItem2**](ECommerceCartLineItem2.md) |  | 

### Return type

[**ECommerceCartLineItem**](ECommerceCartLineItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdCustomersId

> ECommerceCustomer PatchEcommerceStoresIdCustomersId(ctx, storeId, customerId).Body(body).Execute()

Update customer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	customerId := "customerId_example" // string | The id for the customer of a store.
	body := *openapiclient.NewECommerceCustomer4() // ECommerceCustomer4 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdCustomersId(context.Background(), storeId, customerId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdCustomersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdCustomersId`: ECommerceCustomer
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdCustomersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**customerId** | **string** | The id for the customer of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdCustomersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceCustomer4**](ECommerceCustomer4.md) |  | 

### Return type

[**ECommerceCustomer**](ECommerceCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdOrdersId

> ECommerceOrder PatchEcommerceStoresIdOrdersId(ctx, storeId, orderId).Body(body).Execute()

Update order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	orderId := "orderId_example" // string | The id for the order in a store.
	body := *openapiclient.NewECommerceOrder2() // ECommerceOrder2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdOrdersId(context.Background(), storeId, orderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdOrdersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdOrdersId`: ECommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdOrdersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**orderId** | **string** | The id for the order in a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdOrdersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceOrder2**](ECommerceOrder2.md) |  | 

### Return type

[**ECommerceOrder**](ECommerceOrder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdOrdersIdLinesId

> ECommerceOrderLineItem PatchEcommerceStoresIdOrdersIdLinesId(ctx, storeId, orderId, lineId).Body(body).Execute()

Update order line item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	orderId := "orderId_example" // string | The id for the order in a store.
	lineId := "lineId_example" // string | The id for the line item of an order.
	body := *openapiclient.NewECommerceOrderLineItem2() // ECommerceOrderLineItem2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdOrdersIdLinesId(context.Background(), storeId, orderId, lineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdOrdersIdLinesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdOrdersIdLinesId`: ECommerceOrderLineItem
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdOrdersIdLinesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**orderId** | **string** | The id for the order in a store. | 
**lineId** | **string** | The id for the line item of an order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdOrdersIdLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ECommerceOrderLineItem2**](ECommerceOrderLineItem2.md) |  | 

### Return type

[**ECommerceOrderLineItem**](ECommerceOrderLineItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdProductsId

> ECommerceProduct PatchEcommerceStoresIdProductsId(ctx, storeId, productId).Body(body).Execute()

Update product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	body := *openapiclient.NewECommerceProduct2() // ECommerceProduct2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdProductsId(context.Background(), storeId, productId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdProductsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdProductsId`: ECommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdProductsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdProductsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceProduct2**](ECommerceProduct2.md) |  | 

### Return type

[**ECommerceProduct**](ECommerceProduct.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdProductsIdImagesId

> ECommerceProductImage PatchEcommerceStoresIdProductsIdImagesId(ctx, storeId, productId, imageId).Body(body).Execute()

Update product image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	imageId := "imageId_example" // string | The id for the product image.
	body := *openapiclient.NewECommerceProductImage2() // ECommerceProductImage2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdProductsIdImagesId(context.Background(), storeId, productId, imageId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdProductsIdImagesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdProductsIdImagesId`: ECommerceProductImage
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdProductsIdImagesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 
**imageId** | **string** | The id for the product image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdProductsIdImagesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ECommerceProductImage2**](ECommerceProductImage2.md) |  | 

### Return type

[**ECommerceProductImage**](ECommerceProductImage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdProductsIdVariantsId

> ECommerceProductVariant PatchEcommerceStoresIdProductsIdVariantsId(ctx, storeId, productId, variantId).Body(body).Execute()

Update product variant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	variantId := "variantId_example" // string | The id for the product variant.
	body := *openapiclient.NewECommerceProductVariant2() // ECommerceProductVariant2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdProductsIdVariantsId(context.Background(), storeId, productId, variantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdProductsIdVariantsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdProductsIdVariantsId`: ECommerceProductVariant
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdProductsIdVariantsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 
**variantId** | **string** | The id for the product variant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdProductsIdVariantsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ECommerceProductVariant2**](ECommerceProductVariant2.md) |  | 

### Return type

[**ECommerceProductVariant**](ECommerceProductVariant.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdPromocodesId

> ECommercePromoCode PatchEcommerceStoresIdPromocodesId(ctx, storeId, promoRuleId, promoCodeId).Body(body).Execute()

Update promo code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	promoRuleId := "promoRuleId_example" // string | The id for the promo rule of a store.
	promoCodeId := "promoCodeId_example" // string | The id for the promo code of a store.
	body := *openapiclient.NewECommercePromoCode2() // ECommercePromoCode2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdPromocodesId(context.Background(), storeId, promoRuleId, promoCodeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdPromocodesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdPromocodesId`: ECommercePromoCode
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdPromocodesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**promoRuleId** | **string** | The id for the promo rule of a store. | 
**promoCodeId** | **string** | The id for the promo code of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdPromocodesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ECommercePromoCode2**](ECommercePromoCode2.md) |  | 

### Return type

[**ECommercePromoCode**](ECommercePromoCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEcommerceStoresIdPromorulesId

> ECommercePromoRule PatchEcommerceStoresIdPromorulesId(ctx, storeId, promoRuleId).Body(body).Execute()

Update promo rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	promoRuleId := "promoRuleId_example" // string | The id for the promo rule of a store.
	body := *openapiclient.NewECommercePromoRule2() // ECommercePromoRule2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PatchEcommerceStoresIdPromorulesId(context.Background(), storeId, promoRuleId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PatchEcommerceStoresIdPromorulesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEcommerceStoresIdPromorulesId`: ECommercePromoRule
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PatchEcommerceStoresIdPromorulesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**promoRuleId** | **string** | The id for the promo rule of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEcommerceStoresIdPromorulesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommercePromoRule2**](ECommercePromoRule2.md) |  | 

### Return type

[**ECommercePromoRule**](ECommercePromoRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStores

> ECommerceStore PostEcommerceStores(ctx).Body(body).Execute()

Add store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	body := *openapiclient.NewECommerceStore1("example_store", "1a2df69511", "Freddie's Cat Hat Emporium", "USD") // ECommerceStore1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStores(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStores`: ECommerceStore
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ECommerceStore1**](ECommerceStore1.md) |  | 

### Return type

[**ECommerceStore**](ECommerceStore.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdCarts

> ECommerceCart PostEcommerceStoresIdCarts(ctx, storeId).Body(body).Execute()

Add cart



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	body := *openapiclient.NewECommerceCart1("Id_example", *openapiclient.NewECommerceCustomer1("Id_example", false), "CurrencyCode_example", float32(123), []openapiclient.ECommerceCartLineItem1{*openapiclient.NewECommerceCartLineItem1("Id_example", "ProductId_example", "ProductVariantId_example", int32(123), float32(123))}) // ECommerceCart1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdCarts(context.Background(), storeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdCarts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdCarts`: ECommerceCart
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdCarts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdCartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ECommerceCart1**](ECommerceCart1.md) |  | 

### Return type

[**ECommerceCart**](ECommerceCart.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdCartsIdLines

> ECommerceCartLineItem PostEcommerceStoresIdCartsIdLines(ctx, storeId, cartId).Body(body).Execute()

Add cart line item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	cartId := "cartId_example" // string | The id for the cart.
	body := *openapiclient.NewECommerceCartLineItem1("Id_example", "ProductId_example", "ProductVariantId_example", int32(123), float32(123)) // ECommerceCartLineItem1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdCartsIdLines(context.Background(), storeId, cartId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdCartsIdLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdCartsIdLines`: ECommerceCartLineItem
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdCartsIdLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**cartId** | **string** | The id for the cart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdCartsIdLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceCartLineItem1**](ECommerceCartLineItem1.md) |  | 

### Return type

[**ECommerceCartLineItem**](ECommerceCartLineItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdCustomers

> ECommerceCustomer PostEcommerceStoresIdCustomers(ctx, storeId).Body(body).Execute()

Add customer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	body := *openapiclient.NewECommerceCustomer3("Id_example", "EmailAddress_example", false) // ECommerceCustomer3 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdCustomers(context.Background(), storeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdCustomers`: ECommerceCustomer
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdCustomers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ECommerceCustomer3**](ECommerceCustomer3.md) |  | 

### Return type

[**ECommerceCustomer**](ECommerceCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdOrders

> ECommerceOrder PostEcommerceStoresIdOrders(ctx, storeId).Body(body).Execute()

Add order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	body := *openapiclient.NewECommerceOrder1("Id_example", *openapiclient.NewECommerceCustomer5("Id_example"), "CurrencyCode_example", float32(123), []openapiclient.ECommerceOrderLineItem1{*openapiclient.NewECommerceOrderLineItem1("Id_example", "ProductId_example", "ProductVariantId_example", int32(123), float32(123))}) // ECommerceOrder1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdOrders(context.Background(), storeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdOrders`: ECommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ECommerceOrder1**](ECommerceOrder1.md) |  | 

### Return type

[**ECommerceOrder**](ECommerceOrder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdOrdersIdLines

> ECommerceOrderLineItem PostEcommerceStoresIdOrdersIdLines(ctx, storeId, orderId).Body(body).Execute()

Add order line item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	orderId := "orderId_example" // string | The id for the order in a store.
	body := *openapiclient.NewECommerceOrderLineItem1("Id_example", "ProductId_example", "ProductVariantId_example", int32(123), float32(123)) // ECommerceOrderLineItem1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdOrdersIdLines(context.Background(), storeId, orderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdOrdersIdLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdOrdersIdLines`: ECommerceOrderLineItem
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdOrdersIdLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**orderId** | **string** | The id for the order in a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdOrdersIdLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceOrderLineItem1**](ECommerceOrderLineItem1.md) |  | 

### Return type

[**ECommerceOrderLineItem**](ECommerceOrderLineItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdProducts

> ECommerceProduct PostEcommerceStoresIdProducts(ctx, storeId).Body(body).Execute()

Add product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	body := *openapiclient.NewECommerceProduct1("Id_example", "Cat Hat", []openapiclient.ECommerceProductVariant1{*openapiclient.NewECommerceProductVariant1("Id_example", "Cat Hat")}) // ECommerceProduct1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdProducts(context.Background(), storeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdProducts`: ECommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ECommerceProduct1**](ECommerceProduct1.md) |  | 

### Return type

[**ECommerceProduct**](ECommerceProduct.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdProductsIdImages

> ECommerceProductImage PostEcommerceStoresIdProductsIdImages(ctx, storeId, productId).Body(body).Execute()

Add product image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	body := *openapiclient.NewECommerceProductImage1("Id_example", "Url_example") // ECommerceProductImage1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdProductsIdImages(context.Background(), storeId, productId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdProductsIdImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdProductsIdImages`: ECommerceProductImage
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdProductsIdImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdProductsIdImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceProductImage1**](ECommerceProductImage1.md) |  | 

### Return type

[**ECommerceProductImage**](ECommerceProductImage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdProductsIdVariants

> ECommerceProductVariant PostEcommerceStoresIdProductsIdVariants(ctx, storeId, productId).Body(body).Execute()

Add product variant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	body := *openapiclient.NewECommerceProductVariant1("Id_example", "Cat Hat") // ECommerceProductVariant1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdProductsIdVariants(context.Background(), storeId, productId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdProductsIdVariants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdProductsIdVariants`: ECommerceProductVariant
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdProductsIdVariants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdProductsIdVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceProductVariant1**](ECommerceProductVariant1.md) |  | 

### Return type

[**ECommerceProductVariant**](ECommerceProductVariant.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdPromocodes

> ECommercePromoCode PostEcommerceStoresIdPromocodes(ctx, storeId, promoRuleId).Body(body).Execute()

Add promo code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	promoRuleId := "promoRuleId_example" // string | The id for the promo rule of a store.
	body := *openapiclient.NewECommercePromoCode1("Id_example", "summersale", "A url that applies promo code directly at checkout or a url that points to sale page or store url") // ECommercePromoCode1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdPromocodes(context.Background(), storeId, promoRuleId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdPromocodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdPromocodes`: ECommercePromoCode
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdPromocodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**promoRuleId** | **string** | The id for the promo rule of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdPromocodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommercePromoCode1**](ECommercePromoCode1.md) |  | 

### Return type

[**ECommercePromoCode**](ECommercePromoCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcommerceStoresIdPromorules

> ECommercePromoRule PostEcommerceStoresIdPromorules(ctx, storeId).Body(body).Execute()

Add promo rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	body := *openapiclient.NewECommercePromoRule1("Id_example", "Save BIG during our summer sale!", float32(0.5), "Type_example", "Target_example") // ECommercePromoRule1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PostEcommerceStoresIdPromorules(context.Background(), storeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PostEcommerceStoresIdPromorules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEcommerceStoresIdPromorules`: ECommercePromoRule
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PostEcommerceStoresIdPromorules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEcommerceStoresIdPromorulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ECommercePromoRule1**](ECommercePromoRule1.md) |  | 

### Return type

[**ECommercePromoRule**](ECommercePromoRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEcommerceStoresIdCustomersId

> ECommerceCustomer PutEcommerceStoresIdCustomersId(ctx, storeId, customerId).Body(body).Execute()

Add or update customer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	customerId := "customerId_example" // string | The id for the customer of a store.
	body := *openapiclient.NewECommerceCustomer2("Id_example", "EmailAddress_example", false) // ECommerceCustomer2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PutEcommerceStoresIdCustomersId(context.Background(), storeId, customerId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PutEcommerceStoresIdCustomersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutEcommerceStoresIdCustomersId`: ECommerceCustomer
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PutEcommerceStoresIdCustomersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**customerId** | **string** | The id for the customer of a store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEcommerceStoresIdCustomersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ECommerceCustomer2**](ECommerceCustomer2.md) |  | 

### Return type

[**ECommerceCustomer**](ECommerceCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEcommerceStoresIdProductsIdVariantsId

> ECommerceProductVariant PutEcommerceStoresIdProductsIdVariantsId(ctx, storeId, productId, variantId).Body(body).Execute()

Add or update product variant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	storeId := "storeId_example" // string | The store id.
	productId := "productId_example" // string | The id for the product of a store.
	variantId := "variantId_example" // string | The id for the product variant.
	body := *openapiclient.NewECommerceProductVariant1("Id_example", "Cat Hat") // ECommerceProductVariant1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcommerceAPI.PutEcommerceStoresIdProductsIdVariantsId(context.Background(), storeId, productId, variantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcommerceAPI.PutEcommerceStoresIdProductsIdVariantsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutEcommerceStoresIdProductsIdVariantsId`: ECommerceProductVariant
	fmt.Fprintf(os.Stdout, "Response from `EcommerceAPI.PutEcommerceStoresIdProductsIdVariantsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The store id. | 
**productId** | **string** | The id for the product of a store. | 
**variantId** | **string** | The id for the product variant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEcommerceStoresIdProductsIdVariantsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ECommerceProductVariant1**](ECommerceProductVariant1.md) |  | 

### Return type

[**ECommerceProductVariant**](ECommerceProductVariant.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

