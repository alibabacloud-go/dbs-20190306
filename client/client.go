// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":            dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-chengdu":            dara.String("dbs-api.cn-chengdu.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-huhehaote":          dara.String("dbs-api.cn-huhehaote.aliyuncs.com"),
		"cn-hangzhou":           dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen":           dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"ap-southeast-1":        dara.String("dbs-api.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        dara.String("dbs-api.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        dara.String("dbs-api.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        dara.String("dbs-api.ap-southeast-5.aliyuncs.com"),
		"ap-northeast-1":        dara.String("dbs-api.ap-northeast-1.aliyuncs.com"),
		"us-west-1":             dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"us-east-1":             dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"eu-central-1":          dara.String("dbs-api.eu-central-1.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"ap-south-1":            dara.String("dbs-api.ap-south-1.aliyuncs.com"),
		"eu-west-1":             dara.String("dbs-api.eu-west-1.aliyuncs.com"),
		"me-east-1":             dara.String("dbs-api.me-east-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dbs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a DBS backup schedule.
//
// @param request - ConfigureBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureBackupPlanResponse
func (client *Client) ConfigureBackupPlanWithOptions(request *ConfigureBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *ConfigureBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoStartBackup) {
		query["AutoStartBackup"] = request.AutoStartBackup
	}

	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupLogIntervalSeconds) {
		query["BackupLogIntervalSeconds"] = request.BackupLogIntervalSeconds
	}

	if !dara.IsNil(request.BackupObjects) {
		query["BackupObjects"] = request.BackupObjects
	}

	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupPlanName) {
		query["BackupPlanName"] = request.BackupPlanName
	}

	if !dara.IsNil(request.BackupRateLimit) {
		query["BackupRateLimit"] = request.BackupRateLimit
	}

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.BackupSpeedLimit) {
		query["BackupSpeedLimit"] = request.BackupSpeedLimit
	}

	if !dara.IsNil(request.BackupStartTime) {
		query["BackupStartTime"] = request.BackupStartTime
	}

	if !dara.IsNil(request.BackupStorageType) {
		query["BackupStorageType"] = request.BackupStorageType
	}

	if !dara.IsNil(request.BackupStrategyType) {
		query["BackupStrategyType"] = request.BackupStrategyType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossAliyunId) {
		query["CrossAliyunId"] = request.CrossAliyunId
	}

	if !dara.IsNil(request.CrossRoleName) {
		query["CrossRoleName"] = request.CrossRoleName
	}

	if !dara.IsNil(request.DuplicationArchivePeriod) {
		query["DuplicationArchivePeriod"] = request.DuplicationArchivePeriod
	}

	if !dara.IsNil(request.DuplicationInfrequentAccessPeriod) {
		query["DuplicationInfrequentAccessPeriod"] = request.DuplicationInfrequentAccessPeriod
	}

	if !dara.IsNil(request.EnableBackupLog) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !dara.IsNil(request.OSSBucketName) {
		query["OSSBucketName"] = request.OSSBucketName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a DBS backup schedule.
//
// @param request - ConfigureBackupPlanRequest
//
// @return ConfigureBackupPlanResponse
func (client *Client) ConfigureBackupPlan(request *ConfigureBackupPlanRequest) (_result *ConfigureBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigureBackupPlanResponse{}
	_body, _err := client.ConfigureBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates, configures, and starts a backup schedule.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/70005.html) of Database Backup (DBS).
//
// @param request - CreateAndStartBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndStartBackupPlanResponse
func (client *Client) CreateAndStartBackupPlanWithOptions(request *CreateAndStartBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateAndStartBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupLogIntervalSeconds) {
		query["BackupLogIntervalSeconds"] = request.BackupLogIntervalSeconds
	}

	if !dara.IsNil(request.BackupMethod) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !dara.IsNil(request.BackupObjects) {
		query["BackupObjects"] = request.BackupObjects
	}

	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupPlanName) {
		query["BackupPlanName"] = request.BackupPlanName
	}

	if !dara.IsNil(request.BackupRateLimit) {
		query["BackupRateLimit"] = request.BackupRateLimit
	}

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.BackupSpeedLimit) {
		query["BackupSpeedLimit"] = request.BackupSpeedLimit
	}

	if !dara.IsNil(request.BackupStartTime) {
		query["BackupStartTime"] = request.BackupStartTime
	}

	if !dara.IsNil(request.BackupStorageType) {
		query["BackupStorageType"] = request.BackupStorageType
	}

	if !dara.IsNil(request.BackupStrategyType) {
		query["BackupStrategyType"] = request.BackupStrategyType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossAliyunId) {
		query["CrossAliyunId"] = request.CrossAliyunId
	}

	if !dara.IsNil(request.CrossRoleName) {
		query["CrossRoleName"] = request.CrossRoleName
	}

	if !dara.IsNil(request.DatabaseRegion) {
		query["DatabaseRegion"] = request.DatabaseRegion
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.DuplicationArchivePeriod) {
		query["DuplicationArchivePeriod"] = request.DuplicationArchivePeriod
	}

	if !dara.IsNil(request.DuplicationInfrequentAccessPeriod) {
		query["DuplicationInfrequentAccessPeriod"] = request.DuplicationInfrequentAccessPeriod
	}

	if !dara.IsNil(request.EnableBackupLog) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !dara.IsNil(request.FromApp) {
		query["FromApp"] = request.FromApp
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OSSBucketName) {
		query["OSSBucketName"] = request.OSSBucketName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	if !dara.IsNil(request.StorageRegion) {
		query["StorageRegion"] = request.StorageRegion
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAndStartBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndStartBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates, configures, and starts a backup schedule.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/70005.html) of Database Backup (DBS).
//
// @param request - CreateAndStartBackupPlanRequest
//
// @return CreateAndStartBackupPlanResponse
func (client *Client) CreateAndStartBackupPlan(request *CreateAndStartBackupPlanRequest) (_result *CreateAndStartBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAndStartBackupPlanResponse{}
	_body, _err := client.CreateAndStartBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backup schedule.
//
// Description:
//
// For more information about how to create a backup schedule in the Database Backup (DBS) console, see [Purchase a backup schedule](https://help.aliyun.com/document_detail/65909.html).
//
// @param request - CreateBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupPlanResponse
func (client *Client) CreateBackupPlanWithOptions(request *CreateBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupMethod) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DatabaseRegion) {
		query["DatabaseRegion"] = request.DatabaseRegion
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.FromApp) {
		query["FromApp"] = request.FromApp
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StorageRegion) {
		query["StorageRegion"] = request.StorageRegion
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup schedule.
//
// Description:
//
// For more information about how to create a backup schedule in the Database Backup (DBS) console, see [Purchase a backup schedule](https://help.aliyun.com/document_detail/65909.html).
//
// @param request - CreateBackupPlanRequest
//
// @return CreateBackupPlanResponse
func (client *Client) CreateBackupPlan(request *CreateBackupPlanRequest) (_result *CreateBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CreateBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and starts a full backup set download task.
//
// @param request - CreateFullBackupSetDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFullBackupSetDownloadResponse
func (client *Client) CreateFullBackupSetDownloadWithOptions(request *CreateFullBackupSetDownloadRequest, runtime *dara.RuntimeOptions) (_result *CreateFullBackupSetDownloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetDataFormat) {
		query["BackupSetDataFormat"] = request.BackupSetDataFormat
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFullBackupSetDownload"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFullBackupSetDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and starts a full backup set download task.
//
// @param request - CreateFullBackupSetDownloadRequest
//
// @return CreateFullBackupSetDownloadResponse
func (client *Client) CreateFullBackupSetDownload(request *CreateFullBackupSetDownloadRequest) (_result *CreateFullBackupSetDownloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFullBackupSetDownloadResponse{}
	_body, _err := client.CreateFullBackupSetDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task to obtain a database list by using a backup gateway.
//
// Description:
//
// This API operation returns a task ID. You can call the [GetDBListFromAgent](https://help.aliyun.com/document_detail/2869852.html) operation to query the task result based on the task ID.
//
// @param request - CreateGetDBListFromAgentTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGetDBListFromAgentTaskResponse
func (client *Client) CreateGetDBListFromAgentTaskWithOptions(request *CreateGetDBListFromAgentTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateGetDBListFromAgentTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGetDBListFromAgentTask"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGetDBListFromAgentTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to obtain a database list by using a backup gateway.
//
// Description:
//
// This API operation returns a task ID. You can call the [GetDBListFromAgent](https://help.aliyun.com/document_detail/2869852.html) operation to query the task result based on the task ID.
//
// @param request - CreateGetDBListFromAgentTaskRequest
//
// @return CreateGetDBListFromAgentTaskResponse
func (client *Client) CreateGetDBListFromAgentTask(request *CreateGetDBListFromAgentTaskRequest) (_result *CreateGetDBListFromAgentTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGetDBListFromAgentTaskResponse{}
	_body, _err := client.CreateGetDBListFromAgentTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and starts an incremental backup set download task.
//
// @param request - CreateIncrementBackupSetDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIncrementBackupSetDownloadResponse
func (client *Client) CreateIncrementBackupSetDownloadWithOptions(request *CreateIncrementBackupSetDownloadRequest, runtime *dara.RuntimeOptions) (_result *CreateIncrementBackupSetDownloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetDataFormat) {
		query["BackupSetDataFormat"] = request.BackupSetDataFormat
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.BackupSetName) {
		query["BackupSetName"] = request.BackupSetName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIncrementBackupSetDownload"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIncrementBackupSetDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and starts an incremental backup set download task.
//
// @param request - CreateIncrementBackupSetDownloadRequest
//
// @return CreateIncrementBackupSetDownloadResponse
func (client *Client) CreateIncrementBackupSetDownload(request *CreateIncrementBackupSetDownloadRequest) (_result *CreateIncrementBackupSetDownloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIncrementBackupSetDownloadResponse{}
	_body, _err := client.CreateIncrementBackupSetDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a restoration task.
//
// @param request - CreateRestoreTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRestoreTaskResponse
func (client *Client) CreateRestoreTaskWithOptions(request *CreateRestoreTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRestoreTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossAliyunId) {
		query["CrossAliyunId"] = request.CrossAliyunId
	}

	if !dara.IsNil(request.CrossRoleName) {
		query["CrossRoleName"] = request.CrossRoleName
	}

	if !dara.IsNil(request.DestinationEndpointDatabaseName) {
		query["DestinationEndpointDatabaseName"] = request.DestinationEndpointDatabaseName
	}

	if !dara.IsNil(request.DestinationEndpointIP) {
		query["DestinationEndpointIP"] = request.DestinationEndpointIP
	}

	if !dara.IsNil(request.DestinationEndpointInstanceID) {
		query["DestinationEndpointInstanceID"] = request.DestinationEndpointInstanceID
	}

	if !dara.IsNil(request.DestinationEndpointInstanceType) {
		query["DestinationEndpointInstanceType"] = request.DestinationEndpointInstanceType
	}

	if !dara.IsNil(request.DestinationEndpointOracleSID) {
		query["DestinationEndpointOracleSID"] = request.DestinationEndpointOracleSID
	}

	if !dara.IsNil(request.DestinationEndpointPassword) {
		query["DestinationEndpointPassword"] = request.DestinationEndpointPassword
	}

	if !dara.IsNil(request.DestinationEndpointPort) {
		query["DestinationEndpointPort"] = request.DestinationEndpointPort
	}

	if !dara.IsNil(request.DestinationEndpointRegion) {
		query["DestinationEndpointRegion"] = request.DestinationEndpointRegion
	}

	if !dara.IsNil(request.DestinationEndpointUserName) {
		query["DestinationEndpointUserName"] = request.DestinationEndpointUserName
	}

	if !dara.IsNil(request.DuplicateConflict) {
		query["DuplicateConflict"] = request.DuplicateConflict
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RestoreDir) {
		query["RestoreDir"] = request.RestoreDir
	}

	if !dara.IsNil(request.RestoreHome) {
		query["RestoreHome"] = request.RestoreHome
	}

	if !dara.IsNil(request.RestoreObjects) {
		query["RestoreObjects"] = request.RestoreObjects
	}

	if !dara.IsNil(request.RestoreTaskName) {
		query["RestoreTaskName"] = request.RestoreTaskName
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRestoreTask"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRestoreTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a restoration task.
//
// @param request - CreateRestoreTaskRequest
//
// @return CreateRestoreTaskResponse
func (client *Client) CreateRestoreTask(request *CreateRestoreTaskRequest) (_result *CreateRestoreTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRestoreTaskResponse{}
	_body, _err := client.CreateRestoreTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries backup gateways.
//
// @param request - DescribeBackupGatewayListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupGatewayListResponse
func (client *Client) DescribeBackupGatewayListWithOptions(request *DescribeBackupGatewayListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupGatewayListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupGatewayList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupGatewayListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backup gateways.
//
// @param request - DescribeBackupGatewayListRequest
//
// @return DescribeBackupGatewayListResponse
func (client *Client) DescribeBackupGatewayList(request *DescribeBackupGatewayListRequest) (_result *DescribeBackupGatewayListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupGatewayListResponse{}
	_body, _err := client.DescribeBackupGatewayListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the billing information of a backup schedule.
//
// @param request - DescribeBackupPlanBillingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPlanBillingResponse
func (client *Client) DescribeBackupPlanBillingWithOptions(request *DescribeBackupPlanBillingRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPlanBillingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ShowStorageType) {
		query["ShowStorageType"] = request.ShowStorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPlanBilling"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPlanBillingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billing information of a backup schedule.
//
// @param request - DescribeBackupPlanBillingRequest
//
// @return DescribeBackupPlanBillingResponse
func (client *Client) DescribeBackupPlanBilling(request *DescribeBackupPlanBillingRequest) (_result *DescribeBackupPlanBillingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupPlanBillingResponse{}
	_body, _err := client.DescribeBackupPlanBillingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of backup plans
//
// Description:
//
// Before using this interface, please activate the OSS service in advance. For more information, see [Object Storage Service (OSS)](https://help.aliyun.com/document_detail/31817.html).
//
// @param request - DescribeBackupPlanListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPlanListResponse
func (client *Client) DescribeBackupPlanListWithOptions(request *DescribeBackupPlanListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPlanListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupPlanName) {
		query["BackupPlanName"] = request.BackupPlanName
	}

	if !dara.IsNil(request.BackupPlanStatus) {
		query["BackupPlanStatus"] = request.BackupPlanStatus
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPlanList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPlanListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of backup plans
//
// Description:
//
// Before using this interface, please activate the OSS service in advance. For more information, see [Object Storage Service (OSS)](https://help.aliyun.com/document_detail/31817.html).
//
// @param request - DescribeBackupPlanListRequest
//
// @return DescribeBackupPlanListResponse
func (client *Client) DescribeBackupPlanList(request *DescribeBackupPlanListRequest) (_result *DescribeBackupPlanListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupPlanListResponse{}
	_body, _err := client.DescribeBackupPlanListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries backup set download tasks.
//
// @param request - DescribeBackupSetDownloadTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetDownloadTaskListResponse
func (client *Client) DescribeBackupSetDownloadTaskListWithOptions(request *DescribeBackupSetDownloadTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupSetDownloadTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupSetDownloadTaskId) {
		query["BackupSetDownloadTaskId"] = request.BackupSetDownloadTaskId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupSetDownloadTaskList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupSetDownloadTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backup set download tasks.
//
// @param request - DescribeBackupSetDownloadTaskListRequest
//
// @return DescribeBackupSetDownloadTaskListResponse
func (client *Client) DescribeBackupSetDownloadTaskList(request *DescribeBackupSetDownloadTaskListRequest) (_result *DescribeBackupSetDownloadTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupSetDownloadTaskListResponse{}
	_body, _err := client.DescribeBackupSetDownloadTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Data Lake Analytics (DLA) service for a backup schedule.
//
// @param request - DescribeDLAServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDLAServiceResponse
func (client *Client) DescribeDLAServiceWithOptions(request *DescribeDLAServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDLAServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDLAService"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDLAServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Data Lake Analytics (DLA) service for a backup schedule.
//
// @param request - DescribeDLAServiceRequest
//
// @return DescribeDLAServiceResponse
func (client *Client) DescribeDLAService(request *DescribeDLAServiceRequest) (_result *DescribeDLAServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDLAServiceResponse{}
	_body, _err := client.DescribeDLAServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// cn-hangzhou
//
// @param request - DescribeFullBackupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFullBackupListResponse
func (client *Client) DescribeFullBackupListWithOptions(request *DescribeFullBackupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeFullBackupListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShowStorageType) {
		query["ShowStorageType"] = request.ShowStorageType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFullBackupList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFullBackupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// cn-hangzhou
//
// @param request - DescribeFullBackupListRequest
//
// @return DescribeFullBackupListResponse
func (client *Client) DescribeFullBackupList(request *DescribeFullBackupListRequest) (_result *DescribeFullBackupListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFullBackupListResponse{}
	_body, _err := client.DescribeFullBackupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries incremental backup tasks.
//
// @param request - DescribeIncrementBackupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIncrementBackupListResponse
func (client *Client) DescribeIncrementBackupListWithOptions(request *DescribeIncrementBackupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeIncrementBackupListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShowStorageType) {
		query["ShowStorageType"] = request.ShowStorageType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIncrementBackupList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIncrementBackupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries incremental backup tasks.
//
// @param request - DescribeIncrementBackupListRequest
//
// @return DescribeIncrementBackupListResponse
func (client *Client) DescribeIncrementBackupList(request *DescribeIncrementBackupListRequest) (_result *DescribeIncrementBackupListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIncrementBackupListResponse{}
	_body, _err := client.DescribeIncrementBackupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the error information of a Database Backup (DBS) task.
//
// @param request - DescribeJobErrorCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobErrorCodeResponse
func (client *Client) DescribeJobErrorCodeWithOptions(request *DescribeJobErrorCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobErrorCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobErrorCode"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobErrorCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the error information of a Database Backup (DBS) task.
//
// @param request - DescribeJobErrorCodeRequest
//
// @return DescribeJobErrorCodeResponse
func (client *Client) DescribeJobErrorCode(request *DescribeJobErrorCodeRequest) (_result *DescribeJobErrorCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeJobErrorCodeResponse{}
	_body, _err := client.DescribeJobErrorCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the CIDR blocks of nodes on which Database Backup (DBS) is running.
//
// @param request - DescribeNodeCidrListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeCidrListResponse
func (client *Client) DescribeNodeCidrListWithOptions(request *DescribeNodeCidrListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNodeCidrListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNodeCidrList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNodeCidrListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the CIDR blocks of nodes on which Database Backup (DBS) is running.
//
// @param request - DescribeNodeCidrListRequest
//
// @return DescribeNodeCidrListResponse
func (client *Client) DescribeNodeCidrList(request *DescribeNodeCidrListRequest) (_result *DescribeNodeCidrListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNodeCidrListResponse{}
	_body, _err := client.DescribeNodeCidrListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the precheck progress of a backup schedule or a restore task.
//
// @param request - DescribePreCheckProgressListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreCheckProgressListResponse
func (client *Client) DescribePreCheckProgressListWithOptions(request *DescribePreCheckProgressListRequest, runtime *dara.RuntimeOptions) (_result *DescribePreCheckProgressListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RestoreTaskId) {
		query["RestoreTaskId"] = request.RestoreTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePreCheckProgressList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePreCheckProgressListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the precheck progress of a backup schedule or a restore task.
//
// @param request - DescribePreCheckProgressListRequest
//
// @return DescribePreCheckProgressListResponse
func (client *Client) DescribePreCheckProgressList(request *DescribePreCheckProgressListRequest) (_result *DescribePreCheckProgressListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePreCheckProgressListResponse{}
	_body, _err := client.DescribePreCheckProgressListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions that Database Backup (DBS) supports.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions that Database Backup (DBS) supports.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the range of time to which you can restore data in a backup schedule.
//
// @param request - DescribeRestoreRangeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreRangeInfoResponse
func (client *Client) DescribeRestoreRangeInfoWithOptions(request *DescribeRestoreRangeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreRangeInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BeginTimestampForRestore) {
		query["BeginTimestampForRestore"] = request.BeginTimestampForRestore
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestampForRestore) {
		query["EndTimestampForRestore"] = request.EndTimestampForRestore
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecentlyRestore) {
		query["RecentlyRestore"] = request.RecentlyRestore
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreRangeInfo"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreRangeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the range of time to which you can restore data in a backup schedule.
//
// @param request - DescribeRestoreRangeInfoRequest
//
// @return DescribeRestoreRangeInfoResponse
func (client *Client) DescribeRestoreRangeInfo(request *DescribeRestoreRangeInfoRequest) (_result *DescribeRestoreRangeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRestoreRangeInfoResponse{}
	_body, _err := client.DescribeRestoreRangeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries restore tasks.
//
// @param request - DescribeRestoreTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreTaskListResponse
func (client *Client) DescribeRestoreTaskListWithOptions(request *DescribeRestoreTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RestoreTaskId) {
		query["RestoreTaskId"] = request.RestoreTaskId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreTaskList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries restore tasks.
//
// @param request - DescribeRestoreTaskListRequest
//
// @return DescribeRestoreTaskListResponse
func (client *Client) DescribeRestoreTaskList(request *DescribeRestoreTaskListRequest) (_result *DescribeRestoreTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRestoreTaskListResponse{}
	_body, _err := client.DescribeRestoreTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables incremental backup for a backup schedule.
//
// Description:
//
// ### Impact
//
// After you disable the incremental log backup feature, your backup schedule no longer performs incremental log backups.
//
// @param request - DisableBackupLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableBackupLogResponse
func (client *Client) DisableBackupLogWithOptions(request *DisableBackupLogRequest, runtime *dara.RuntimeOptions) (_result *DisableBackupLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableBackupLog"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableBackupLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables incremental backup for a backup schedule.
//
// Description:
//
// ### Impact
//
// After you disable the incremental log backup feature, your backup schedule no longer performs incremental log backups.
//
// @param request - DisableBackupLogRequest
//
// @return DisableBackupLogResponse
func (client *Client) DisableBackupLog(request *DisableBackupLogRequest) (_result *DisableBackupLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableBackupLogResponse{}
	_body, _err := client.DisableBackupLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables incremental backup for a backup schedule.
//
// Description:
//
// ## Impact
//
// It is free to enable the incremental log backup feature. However, the backup traffic and storage capacity generated by the feature are billed in the same way as the full backup feature, and can be offset by the free quota of backup schedules or storage plans.
//
// @param request - EnableBackupLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableBackupLogResponse
func (client *Client) EnableBackupLogWithOptions(request *EnableBackupLogRequest, runtime *dara.RuntimeOptions) (_result *EnableBackupLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableBackupLog"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableBackupLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables incremental backup for a backup schedule.
//
// Description:
//
// ## Impact
//
// It is free to enable the incremental log backup feature. However, the backup traffic and storage capacity generated by the feature are billed in the same way as the full backup feature, and can be offset by the free quota of backup schedules or storage plans.
//
// @param request - EnableBackupLogRequest
//
// @return EnableBackupLogResponse
func (client *Client) EnableBackupLog(request *EnableBackupLogRequest) (_result *EnableBackupLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableBackupLogResponse{}
	_body, _err := client.EnableBackupLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of a task that is used to query a database list by using a backup gateway based on the ID of the task.
//
// @param request - GetDBListFromAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDBListFromAgentResponse
func (client *Client) GetDBListFromAgentWithOptions(request *GetDBListFromAgentRequest, runtime *dara.RuntimeOptions) (_result *GetDBListFromAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDBListFromAgent"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDBListFromAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of a task that is used to query a database list by using a backup gateway based on the ID of the task.
//
// @param request - GetDBListFromAgentRequest
//
// @return GetDBListFromAgentResponse
func (client *Client) GetDBListFromAgent(request *GetDBListFromAgentRequest) (_result *GetDBListFromAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDBListFromAgentResponse{}
	_body, _err := client.GetDBListFromAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants the AliyunServiceRoleForDBS role to Database Backup (DBS).
//
// @param request - InitializeDbsServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeDbsServiceLinkedRoleResponse
func (client *Client) InitializeDbsServiceLinkedRoleWithOptions(runtime *dara.RuntimeOptions) (_result *InitializeDbsServiceLinkedRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("InitializeDbsServiceLinkedRole"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeDbsServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants the AliyunServiceRoleForDBS role to Database Backup (DBS).
//
// @return InitializeDbsServiceLinkedRoleResponse
func (client *Client) InitializeDbsServiceLinkedRole() (_result *InitializeDbsServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitializeDbsServiceLinkedRoleResponse{}
	_body, _err := client.InitializeDbsServiceLinkedRoleWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies backup objects of a backup schedule in Database Backup (DBS).
//
// @param request - ModifyBackupObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupObjectsResponse
func (client *Client) ModifyBackupObjectsWithOptions(request *ModifyBackupObjectsRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupObjects) {
		query["BackupObjects"] = request.BackupObjects
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupObjects"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies backup objects of a backup schedule in Database Backup (DBS).
//
// @param request - ModifyBackupObjectsRequest
//
// @return ModifyBackupObjectsResponse
func (client *Client) ModifyBackupObjects(request *ModifyBackupObjectsRequest) (_result *ModifyBackupObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupObjectsResponse{}
	_body, _err := client.ModifyBackupObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a backup schedule.
//
// @param request - ModifyBackupPlanNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPlanNameResponse
func (client *Client) ModifyBackupPlanNameWithOptions(request *ModifyBackupPlanNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPlanNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupPlanName) {
		query["BackupPlanName"] = request.BackupPlanName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPlanName"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupPlanNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a backup schedule.
//
// @param request - ModifyBackupPlanNameRequest
//
// @return ModifyBackupPlanNameResponse
func (client *Client) ModifyBackupPlanName(request *ModifyBackupPlanNameRequest) (_result *ModifyBackupPlanNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupPlanNameResponse{}
	_body, _err := client.ModifyBackupPlanNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables, configures, or disables the automatic download feature.
//
// @param request - ModifyBackupSetDownloadRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupSetDownloadRulesResponse
func (client *Client) ModifyBackupSetDownloadRulesWithOptions(request *ModifyBackupSetDownloadRulesRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupSetDownloadRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupSetDownloadDir) {
		query["BackupSetDownloadDir"] = request.BackupSetDownloadDir
	}

	if !dara.IsNil(request.BackupSetDownloadTargetType) {
		query["BackupSetDownloadTargetType"] = request.BackupSetDownloadTargetType
	}

	if !dara.IsNil(request.BackupSetDownloadTargetTypeLocation) {
		query["BackupSetDownloadTargetTypeLocation"] = request.BackupSetDownloadTargetTypeLocation
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FullDataFormat) {
		query["FullDataFormat"] = request.FullDataFormat
	}

	if !dara.IsNil(request.IncrementDataFormat) {
		query["IncrementDataFormat"] = request.IncrementDataFormat
	}

	if !dara.IsNil(request.OpenAutoDownload) {
		query["OpenAutoDownload"] = request.OpenAutoDownload
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupSetDownloadRules"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupSetDownloadRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables, configures, or disables the automatic download feature.
//
// @param request - ModifyBackupSetDownloadRulesRequest
//
// @return ModifyBackupSetDownloadRulesResponse
func (client *Client) ModifyBackupSetDownloadRules(request *ModifyBackupSetDownloadRulesRequest) (_result *ModifyBackupSetDownloadRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupSetDownloadRulesResponse{}
	_body, _err := client.ModifyBackupSetDownloadRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the data source of a backup schedule.
//
// @param request - ModifyBackupSourceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupSourceEndpointResponse
func (client *Client) ModifyBackupSourceEndpointWithOptions(request *ModifyBackupSourceEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupSourceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupObjects) {
		query["BackupObjects"] = request.BackupObjects
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossAliyunId) {
		query["CrossAliyunId"] = request.CrossAliyunId
	}

	if !dara.IsNil(request.CrossRoleName) {
		query["CrossRoleName"] = request.CrossRoleName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupSourceEndpoint"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupSourceEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the data source of a backup schedule.
//
// @param request - ModifyBackupSourceEndpointRequest
//
// @return ModifyBackupSourceEndpointResponse
func (client *Client) ModifyBackupSourceEndpoint(request *ModifyBackupSourceEndpointRequest) (_result *ModifyBackupSourceEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupSourceEndpointResponse{}
	_body, _err := client.ModifyBackupSourceEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the backup time of a backup schedule.
//
// @param request - ModifyBackupStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupStrategyResponse
func (client *Client) ModifyBackupStrategyWithOptions(request *ModifyBackupStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupLogIntervalSeconds) {
		query["BackupLogIntervalSeconds"] = request.BackupLogIntervalSeconds
	}

	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupStartTime) {
		query["BackupStartTime"] = request.BackupStartTime
	}

	if !dara.IsNil(request.BackupStrategyType) {
		query["BackupStrategyType"] = request.BackupStrategyType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupStrategy"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the backup time of a backup schedule.
//
// @param request - ModifyBackupStrategyRequest
//
// @return ModifyBackupStrategyResponse
func (client *Client) ModifyBackupStrategy(request *ModifyBackupStrategyRequest) (_result *ModifyBackupStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupStrategyResponse{}
	_body, _err := client.ModifyBackupStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the lifecycle of data that is backed up based on a backup schedule.
//
// @param request - ModifyStorageStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyStorageStrategyResponse
func (client *Client) ModifyStorageStrategyWithOptions(request *ModifyStorageStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyStorageStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DuplicationArchivePeriod) {
		query["DuplicationArchivePeriod"] = request.DuplicationArchivePeriod
	}

	if !dara.IsNil(request.DuplicationInfrequentAccessPeriod) {
		query["DuplicationInfrequentAccessPeriod"] = request.DuplicationInfrequentAccessPeriod
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyStorageStrategy"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyStorageStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the lifecycle of data that is backed up based on a backup schedule.
//
// @param request - ModifyStorageStrategyRequest
//
// @return ModifyStorageStrategyResponse
func (client *Client) ModifyStorageStrategy(request *ModifyStorageStrategyRequest) (_result *ModifyStorageStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyStorageStrategyResponse{}
	_body, _err := client.ModifyStorageStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go backup schedule.
//
// Description:
//
// ## Impacts
//
// After a pay-as-you-go backup schedule is released, it stops providing services. Database Backup (DBS) no longer charges you fees for this backup schedule.
//
// @param request - ReleaseBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseBackupPlanResponse
func (client *Client) ReleaseBackupPlanWithOptions(request *ReleaseBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *ReleaseBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go backup schedule.
//
// Description:
//
// ## Impacts
//
// After a pay-as-you-go backup schedule is released, it stops providing services. Database Backup (DBS) no longer charges you fees for this backup schedule.
//
// @param request - ReleaseBackupPlanRequest
//
// @return ReleaseBackupPlanResponse
func (client *Client) ReleaseBackupPlan(request *ReleaseBackupPlanRequest) (_result *ReleaseBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseBackupPlanResponse{}
	_body, _err := client.ReleaseBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a backup schedule.
//
// @param request - RenewBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewBackupPlanResponse
func (client *Client) RenewBackupPlanWithOptions(request *RenewBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *RenewBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a backup schedule.
//
// @param request - RenewBackupPlanRequest
//
// @return RenewBackupPlanResponse
func (client *Client) RenewBackupPlan(request *RenewBackupPlanRequest) (_result *RenewBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewBackupPlanResponse{}
	_body, _err := client.RenewBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a backup schedule.
//
// @param request - StartBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBackupPlanResponse
func (client *Client) StartBackupPlanWithOptions(request *StartBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *StartBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a backup schedule.
//
// @param request - StartBackupPlanRequest
//
// @return StartBackupPlanResponse
func (client *Client) StartBackupPlan(request *StartBackupPlanRequest) (_result *StartBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartBackupPlanResponse{}
	_body, _err := client.StartBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a restore task.
//
// @param request - StartRestoreTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRestoreTaskResponse
func (client *Client) StartRestoreTaskWithOptions(request *StartRestoreTaskRequest, runtime *dara.RuntimeOptions) (_result *StartRestoreTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RestoreTaskId) {
		query["RestoreTaskId"] = request.RestoreTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRestoreTask"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRestoreTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a restore task.
//
// @param request - StartRestoreTaskRequest
//
// @return StartRestoreTaskResponse
func (client *Client) StartRestoreTask(request *StartRestoreTaskRequest) (_result *StartRestoreTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartRestoreTaskResponse{}
	_body, _err := client.StartRestoreTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a backup schedule.
//
// @param request - StopBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopBackupPlanResponse
func (client *Client) StopBackupPlanWithOptions(request *StopBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *StopBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StopMethod) {
		query["StopMethod"] = request.StopMethod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a backup schedule.
//
// @param request - StopBackupPlanRequest
//
// @return StopBackupPlanResponse
func (client *Client) StopBackupPlan(request *StopBackupPlanRequest) (_result *StopBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopBackupPlanResponse{}
	_body, _err := client.StopBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades a backup schedule.
//
// @param request - UpgradeBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeBackupPlanResponse
func (client *Client) UpgradeBackupPlanWithOptions(request *UpgradeBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *UpgradeBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a backup schedule.
//
// @param request - UpgradeBackupPlanRequest
//
// @return UpgradeBackupPlanResponse
func (client *Client) UpgradeBackupPlan(request *UpgradeBackupPlanRequest) (_result *UpgradeBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeBackupPlanResponse{}
	_body, _err := client.UpgradeBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
