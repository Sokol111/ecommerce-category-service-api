# DefaultApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createCategory**](#createcategory) | **POST** /category/create | Create a new category|
|[**getAll**](#getall) | **GET** /category/list | Get a list of all categories|
|[**getCategoryById**](#getcategorybyid) | **GET** /category/get/{id} | Get a category by ID|
|[**updateCategory**](#updatecategory) | **PUT** /category/update | Update an existing category|

# **createCategory**
> CategoryResponse createCategory(createCategoryRequest)


### Example

```typescript
import {
    DefaultApi,
    Configuration,
    CreateCategoryRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let createCategoryRequest: CreateCategoryRequest; //

const { status, data } = await apiInstance.createCategory(
    createCategoryRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createCategoryRequest** | **CreateCategoryRequest**|  | |


### Return type

**CategoryResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Category created successfully |  -  |
|**400** | Validation error |  -  |
|**500** | Standard internal error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAll**
> Array<CategoryResponse> getAll()


### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

const { status, data } = await apiInstance.getAll();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Array<CategoryResponse>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of categories retrieved successfully |  -  |
|**500** | Standard internal error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCategoryById**
> CategoryResponse getCategoryById()


### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let id: string; // (default to undefined)

const { status, data } = await apiInstance.getCategoryById(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] |  | defaults to undefined|


### Return type

**CategoryResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Category retrieved successfully |  -  |
|**404** | Category not found |  -  |
|**500** | Standard internal error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateCategory**
> CategoryResponse updateCategory(updateCategoryRequest)


### Example

```typescript
import {
    DefaultApi,
    Configuration,
    UpdateCategoryRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let updateCategoryRequest: UpdateCategoryRequest; //

const { status, data } = await apiInstance.updateCategory(
    updateCategoryRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateCategoryRequest** | **UpdateCategoryRequest**|  | |


### Return type

**CategoryResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Category updated successfully |  -  |
|**412** | Versions mismatch |  -  |
|**400** | Validation error |  -  |
|**500** | Standard internal error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

