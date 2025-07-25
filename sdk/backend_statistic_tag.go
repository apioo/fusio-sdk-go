
// BackendStatisticTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    "encoding/json"
    "errors"
    "fmt"
    "github.com/apioo/sdkgen-go/v2"
    "io"
    "net/http"
    "net/url"
    
)

type BackendStatisticTag struct {
    internal *sdkgen.TagAbstract
}



// GetActivitiesPerUser Returns a statistic containing the activities per user
func (client *BackendStatisticTag) GetActivitiesPerUser(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/activities_per_user", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetCountRequests Returns a statistic containing the request count
func (client *BackendStatisticTag) GetCountRequests(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticCount, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/count_requests", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticCount
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetErrorsPerOperation Returns a statistic containing the errors per operation
func (client *BackendStatisticTag) GetErrorsPerOperation(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/errors_per_operation", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetIncomingRequests Returns a statistic containing the incoming requests
func (client *BackendStatisticTag) GetIncomingRequests(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/incoming_requests", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetIncomingTransactions Returns a statistic containing the incoming transactions
func (client *BackendStatisticTag) GetIncomingTransactions(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/incoming_transactions", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetIssuedTokens Returns a statistic containing the issues tokens
func (client *BackendStatisticTag) GetIssuedTokens(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/issued_tokens", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetMostUsedActivities Returns a statistic containing the most used activities
func (client *BackendStatisticTag) GetMostUsedActivities(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/most_used_activities", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetMostUsedApps Returns a statistic containing the most used apps
func (client *BackendStatisticTag) GetMostUsedApps(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/most_used_apps", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetMostUsedOperations Returns a statistic containing the most used operations
func (client *BackendStatisticTag) GetMostUsedOperations(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/most_used_operations", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetTestCoverage Returns a statistic containing the test coverage
func (client *BackendStatisticTag) GetTestCoverage() (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/test_coverage", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetTimeAverage Returns a statistic containing the time average
func (client *BackendStatisticTag) GetTimeAverage(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/time_average", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetTimePerOperation Returns a statistic containing the time per operation
func (client *BackendStatisticTag) GetTimePerOperation(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/time_per_operation", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetUsedPoints Returns a statistic containing the used points
func (client *BackendStatisticTag) GetUsedPoints(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/used_points", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetUserRegistrations Returns a statistic containing the user registrations
func (client *BackendStatisticTag) GetUserRegistrations(startIndex int, count int, search string, from string, to string, operationId int, appId int, userId int, ip string, userAgent string, method string, path string, header string, body string) (*BackendStatisticChart, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["operationId"] = operationId
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["ip"] = ip
    queryParams["userAgent"] = userAgent
    queryParams["method"] = method
    queryParams["path"] = path
    queryParams["header"] = header
    queryParams["body"] = body

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/statistic/user_registrations", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendStatisticChart
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewBackendStatisticTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendStatisticTag {
	return &BackendStatisticTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
