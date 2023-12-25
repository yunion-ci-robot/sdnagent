// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	"yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"

	"yunion.io/x/onecloud/pkg/apis"
	"yunion.io/x/onecloud/pkg/util/choices"
)

const (
	CLOUD_PROVIDER_INIT          = "init"
	CLOUD_PROVIDER_CONNECTED     = "connected"
	CLOUD_PROVIDER_DISCONNECTED  = "disconnected"
	CLOUD_PROVIDER_START_DELETE  = "start_delete"
	CLOUD_PROVIDER_DELETING      = "deleting"
	CLOUD_PROVIDER_DELETED       = "deleted"
	CLOUD_PROVIDER_DELETE_FAILED = "delete_failed"
	CLOUD_PROVIDER_SYNC_NETWORK  = "sync_network"

	CLOUD_PROVIDER_SYNC_NETWORK_FAILED = "sync_net_failed"

	CLOUD_PROVIDER_SYNC_STATUS_QUEUING = "queuing"
	CLOUD_PROVIDER_SYNC_STATUS_QUEUED  = "queued"
	CLOUD_PROVIDER_SYNC_STATUS_SYNCING = "syncing"
	CLOUD_PROVIDER_SYNC_STATUS_IDLE    = "idle"
	CLOUD_PROVIDER_SYNC_STATUS_ERROR   = "error"

	CLOUD_PROVIDER_ONECLOUD       = compute.CLOUD_PROVIDER_ONECLOUD
	CLOUD_PROVIDER_VMWARE         = compute.CLOUD_PROVIDER_VMWARE
	CLOUD_PROVIDER_NUTANIX        = compute.CLOUD_PROVIDER_NUTANIX
	CLOUD_PROVIDER_ALIYUN         = compute.CLOUD_PROVIDER_ALIYUN
	CLOUD_PROVIDER_APSARA         = compute.CLOUD_PROVIDER_APSARA
	CLOUD_PROVIDER_QCLOUD         = compute.CLOUD_PROVIDER_QCLOUD
	CLOUD_PROVIDER_AZURE          = compute.CLOUD_PROVIDER_AZURE
	CLOUD_PROVIDER_AWS            = compute.CLOUD_PROVIDER_AWS
	CLOUD_PROVIDER_HUAWEI         = compute.CLOUD_PROVIDER_HUAWEI
	CLOUD_PROVIDER_HCSO           = compute.CLOUD_PROVIDER_HCSO
	CLOUD_PROVIDER_HCSOP          = compute.CLOUD_PROVIDER_HCSOP
	CLOUD_PROVIDER_HCS            = compute.CLOUD_PROVIDER_HCS
	CLOUD_PROVIDER_OPENSTACK      = compute.CLOUD_PROVIDER_OPENSTACK
	CLOUD_PROVIDER_UCLOUD         = compute.CLOUD_PROVIDER_UCLOUD
	CLOUD_PROVIDER_VOLCENGINE     = compute.CLOUD_PROVIDER_VOLCENGINE
	CLOUD_PROVIDER_ZSTACK         = compute.CLOUD_PROVIDER_ZSTACK
	CLOUD_PROVIDER_GOOGLE         = compute.CLOUD_PROVIDER_GOOGLE
	CLOUD_PROVIDER_CTYUN          = compute.CLOUD_PROVIDER_CTYUN
	CLOUD_PROVIDER_ECLOUD         = compute.CLOUD_PROVIDER_ECLOUD
	CLOUD_PROVIDER_JDCLOUD        = compute.CLOUD_PROVIDER_JDCLOUD
	CLOUD_PROVIDER_CLOUDPODS      = compute.CLOUD_PROVIDER_CLOUDPODS
	CLOUD_PROVIDER_BINGO_CLOUD    = compute.CLOUD_PROVIDER_BINGO_CLOUD
	CLOUD_PROVIDER_INCLOUD_SPHERE = compute.CLOUD_PROVIDER_INCLOUD_SPHERE
	CLOUD_PROVIDER_PROXMOX        = compute.CLOUD_PROVIDER_PROXMOX
	CLOUD_PROVIDER_REMOTEFILE     = compute.CLOUD_PROVIDER_REMOTEFILE
	CLOUD_PROVIDER_H3C            = compute.CLOUD_PROVIDER_H3C
	CLOUD_PROVIDER_KSYUN          = compute.CLOUD_PROVIDER_KSYUN
	CLOUD_PROVIDER_BAIDU          = compute.CLOUD_PROVIDER_BAIDU
	CLOUD_PROVIDER_CUCLOUD        = compute.CLOUD_PROVIDER_CUCLOUD
	CLOUD_PROVIDER_QINGCLOUD      = compute.CLOUD_PROVIDER_QINGCLOUD
	CLOUD_PROVIDER_ORACLE         = compute.CLOUD_PROVIDER_ORACLE

	CLOUD_PROVIDER_GENERICS3 = compute.CLOUD_PROVIDER_GENERICS3
	CLOUD_PROVIDER_CEPH      = compute.CLOUD_PROVIDER_CEPH
	CLOUD_PROVIDER_XSKY      = compute.CLOUD_PROVIDER_XSKY

	CLOUD_PROVIDER_HEALTH_NORMAL        = compute.CLOUD_PROVIDER_HEALTH_NORMAL        // 远端处于健康状态
	CLOUD_PROVIDER_HEALTH_INSUFFICIENT  = compute.CLOUD_PROVIDER_HEALTH_INSUFFICIENT  // 不足按需资源余额
	CLOUD_PROVIDER_HEALTH_SUSPENDED     = compute.CLOUD_PROVIDER_HEALTH_SUSPENDED     // 远端处于冻结状态
	CLOUD_PROVIDER_HEALTH_ARREARS       = compute.CLOUD_PROVIDER_HEALTH_ARREARS       // 远端处于欠费状态
	CLOUD_PROVIDER_HEALTH_UNKNOWN       = compute.CLOUD_PROVIDER_HEALTH_UNKNOWN       // 未知状态，查询失败
	CLOUD_PROVIDER_HEALTH_NO_PERMISSION = compute.CLOUD_PROVIDER_HEALTH_NO_PERMISSION // 没有权限获取账单信息

	ZSTACK_BRAND_DSTACK     = compute.ZSTACK_BRAND_DSTACK
	ONECLOUD_BRAND_ONECLOUD = "OneCloud"

	CLOUD_ACCOUNT_WIRE_LEVEL_VCENTER    = "vcenter"
	CLOUD_ACCOUNT_WIRE_LEVEL_DATACENTER = "datacenter"
	CLOUD_ACCOUNT_WIRE_LEVEL_CLUSTER    = "cluster"
)

var CLOUD_ACCOUNT_WIRE_LEVELS = choices.NewChoices(
	CLOUD_ACCOUNT_WIRE_LEVEL_VCENTER,
	CLOUD_ACCOUNT_WIRE_LEVEL_DATACENTER,
	CLOUD_ACCOUNT_WIRE_LEVEL_CLUSTER,
)

const (
	CLOUD_ACCESS_ENV_AWS_GLOBAL          = compute.CLOUD_ACCESS_ENV_AWS_GLOBAL
	CLOUD_ACCESS_ENV_AWS_CHINA           = compute.CLOUD_ACCESS_ENV_AWS_CHINA
	CLOUD_ACCESS_ENV_AZURE_GLOBAL        = compute.CLOUD_ACCESS_ENV_AZURE_GLOBAL
	CLOUD_ACCESS_ENV_AZURE_GERMAN        = compute.CLOUD_ACCESS_ENV_AZURE_GERMAN
	CLOUD_ACCESS_ENV_AZURE_US_GOVERNMENT = compute.CLOUD_ACCESS_ENV_AZURE_US_GOVERNMENT
	CLOUD_ACCESS_ENV_AZURE_CHINA         = compute.CLOUD_ACCESS_ENV_AZURE_CHINA
	CLOUD_ACCESS_ENV_HUAWEI_GLOBAL       = compute.CLOUD_ACCESS_ENV_HUAWEI_GLOBAL
	CLOUD_ACCESS_ENV_HUAWEI_CHINA        = compute.CLOUD_ACCESS_ENV_HUAWEI_CHINA
	CLOUD_ACCESS_ENV_ALIYUN_GLOBAL       = compute.CLOUD_ACCESS_ENV_ALIYUN_GLOBAL
	CLOUD_ACCESS_ENV_ALIYUN_FINANCE      = compute.CLOUD_ACCESS_ENV_ALIYUN_FINANCE
	CLOUD_ACCESS_ENV_CTYUN_CHINA         = compute.CLOUD_ACCESS_ENV_CTYUN_CHINA
	CLOUD_ACCESS_ENV_ECLOUD_CHINA        = compute.CLOUD_ACCESS_ENV_ECLOUD_CHINA
	CLOUD_ACCESS_ENV_JDCLOUD_CHINA       = compute.CLOUD_ACCESS_ENV_JDCLOUD_CHINA
)

var (
	CLOUD_PROVIDER_VALID_STATUS        = []string{CLOUD_PROVIDER_CONNECTED}
	CLOUD_PROVIDER_VALID_HEALTH_STATUS = []string{CLOUD_PROVIDER_HEALTH_NORMAL, CLOUD_PROVIDER_HEALTH_NO_PERMISSION}
	PRIVATE_CLOUD_PROVIDERS            = []string{CLOUD_PROVIDER_ZSTACK, CLOUD_PROVIDER_OPENSTACK, CLOUD_PROVIDER_APSARA,
		CLOUD_PROVIDER_HCSO, CLOUD_PROVIDER_HCS, CLOUD_PROVIDER_HCSOP,
		CLOUD_PROVIDER_INCLOUD_SPHERE, CLOUD_PROVIDER_PROXMOX, CLOUD_PROVIDER_REMOTEFILE,
		CLOUD_PROVIDER_H3C,
	}

	CLOUD_PROVIDERS = []string{
		CLOUD_PROVIDER_ONECLOUD,
		CLOUD_PROVIDER_VMWARE,
		CLOUD_PROVIDER_ALIYUN,
		CLOUD_PROVIDER_APSARA,
		CLOUD_PROVIDER_QCLOUD,
		CLOUD_PROVIDER_AZURE,
		CLOUD_PROVIDER_AWS,
		CLOUD_PROVIDER_HUAWEI,
		CLOUD_PROVIDER_HCSO,
		CLOUD_PROVIDER_HCSOP,
		CLOUD_PROVIDER_HCS,
		CLOUD_PROVIDER_OPENSTACK,
		CLOUD_PROVIDER_UCLOUD,
		CLOUD_PROVIDER_VOLCENGINE,
		CLOUD_PROVIDER_ZSTACK,
		CLOUD_PROVIDER_GOOGLE,
		CLOUD_PROVIDER_CTYUN,
		CLOUD_PROVIDER_ECLOUD,
		CLOUD_PROVIDER_JDCLOUD,
		CLOUD_PROVIDER_CLOUDPODS,
		CLOUD_PROVIDER_NUTANIX,
		CLOUD_PROVIDER_BINGO_CLOUD,
		CLOUD_PROVIDER_INCLOUD_SPHERE,
		CLOUD_PROVIDER_PROXMOX,
		CLOUD_PROVIDER_REMOTEFILE,
		CLOUD_PROVIDER_H3C,
		CLOUD_PROVIDER_KSYUN,
		CLOUD_PROVIDER_BAIDU,
		CLOUD_PROVIDER_CUCLOUD,
		CLOUD_PROVIDER_QINGCLOUD,
		CLOUD_PROVIDER_ORACLE,
	}

	CLOUD_PROVIDER_HOST_TYPE_MAP = map[string][]string{
		CLOUD_PROVIDER_ONECLOUD: {
			HOST_TYPE_KVM,
			HOST_TYPE_BAREMETAL,
			HOST_TYPE_HYPERVISOR,
		},
		CLOUD_PROVIDER_VMWARE: {
			HOST_TYPE_ESXI,
		},
		CLOUD_PROVIDER_ALIYUN: {
			HOST_TYPE_ALIYUN,
		},
		CLOUD_PROVIDER_APSARA: {
			HOST_TYPE_APSARA,
		},
		CLOUD_PROVIDER_QCLOUD: {
			HOST_TYPE_QCLOUD,
		},
		CLOUD_PROVIDER_AZURE: {
			HOST_TYPE_AZURE,
		},
		CLOUD_PROVIDER_AWS: {
			HOST_TYPE_AWS,
		},
		CLOUD_PROVIDER_HUAWEI: {
			HOST_TYPE_HUAWEI,
		},
		CLOUD_PROVIDER_HCSO: {
			HOST_TYPE_HCSO,
		},
		CLOUD_PROVIDER_HCSOP: {
			HOST_TYPE_HCSOP,
		},
		CLOUD_PROVIDER_HCS: {
			HOST_TYPE_HCS,
		},
		CLOUD_PROVIDER_OPENSTACK: {
			HOST_TYPE_OPENSTACK,
		},
		CLOUD_PROVIDER_UCLOUD: {
			HOST_TYPE_UCLOUD,
		},
		CLOUD_PROVIDER_VOLCENGINE: {
			HOST_TYPE_VOLCENGINE,
		},
		CLOUD_PROVIDER_ZSTACK: {
			HOST_TYPE_ZSTACK,
		},
		CLOUD_PROVIDER_GOOGLE: {
			HOST_TYPE_GOOGLE,
		},
		CLOUD_PROVIDER_CTYUN: {
			HOST_TYPE_CTYUN,
		},
		CLOUD_PROVIDER_ECLOUD: {
			HOST_TYPE_ECLOUD,
		},
		CLOUD_PROVIDER_JDCLOUD: {
			HOST_TYPE_JDCLOUD,
		},
		CLOUD_PROVIDER_CLOUDPODS: {
			HOST_TYPE_CLOUDPODS,
		},
		CLOUD_PROVIDER_NUTANIX: {
			HOST_TYPE_NUTANIX,
		},
		CLOUD_PROVIDER_BINGO_CLOUD: {
			HOST_TYPE_BINGO_CLOUD,
		},
		CLOUD_PROVIDER_INCLOUD_SPHERE: {
			HOST_TYPE_INCLOUD_SPHERE,
		},
		CLOUD_PROVIDER_PROXMOX: {
			HOST_TYPE_PROXMOX,
		},
		CLOUD_PROVIDER_REMOTEFILE: {
			HOST_TYPE_REMOTEFILE,
		},
		CLOUD_PROVIDER_H3C: {
			HOST_TYPE_H3C,
		},
		CLOUD_PROVIDER_KSYUN: {
			HOST_TYPE_KSYUN,
		},
		CLOUD_PROVIDER_BAIDU: {
			HOST_TYPE_BAIDU,
		},
		CLOUD_PROVIDER_CUCLOUD: {
			HOST_TYPE_CUCLOUD,
		},
		CLOUD_PROVIDER_QINGCLOUD: {
			HOST_TYPE_QINGCLOUD,
		},
		CLOUD_PROVIDER_ORACLE: {
			HOST_TYPE_ORACLE,
		},
	}
)

const (
	CLOUD_ENV_PUBLIC_CLOUD  = cloudprovider.CLOUD_ENV_PUBLIC_CLOUD
	CLOUD_ENV_PRIVATE_CLOUD = cloudprovider.CLOUD_ENV_PRIVATE_CLOUD
	CLOUD_ENV_ON_PREMISE    = cloudprovider.CLOUD_ENV_ON_PREMISE

	CLOUD_ENV_PRIVATE_ON_PREMISE = cloudprovider.CLOUD_ENV_PRIVATE_ON_PREMISE
)

const (
	CLOUD_ACCOUNT_SHARE_MODE_ACCOUNT_DOMAIN  = apis.CLOUD_ACCOUNT_SHARE_MODE_ACCOUNT_DOMAIN
	CLOUD_ACCOUNT_SHARE_MODE_SYSTEM          = apis.CLOUD_ACCOUNT_SHARE_MODE_SYSTEM
	CLOUD_ACCOUNT_SHARE_MODE_PROVIDER_DOMAIN = apis.CLOUD_ACCOUNT_SHARE_MODE_PROVIDER_DOMAIN
)

var (
	CLOUD_ACCOUNT_SHARE_MODES = []string{
		CLOUD_ACCOUNT_SHARE_MODE_ACCOUNT_DOMAIN,
		CLOUD_ACCOUNT_SHARE_MODE_SYSTEM,
		CLOUD_ACCOUNT_SHARE_MODE_PROVIDER_DOMAIN,
	}

	CLOUD_ENV_MAP = map[string]map[string]string{
		CLOUD_PROVIDER_AZURE: {
			"AzureGermanCloud":       CLOUD_ACCESS_ENV_AZURE_GERMAN,
			"AzureChinaCloud":        CLOUD_ACCESS_ENV_AZURE_CHINA,
			"AzureUSGovernmentCloud": CLOUD_ACCESS_ENV_AZURE_US_GOVERNMENT,
			"AzurePublicCloud":       CLOUD_ACCESS_ENV_AZURE_GLOBAL,
		},
		CLOUD_PROVIDER_AWS: {
			"InternationalCloud": CLOUD_ACCESS_ENV_AWS_GLOBAL,
			"ChinaCloud":         CLOUD_ACCESS_ENV_AWS_CHINA,
		},
		CLOUD_PROVIDER_HUAWEI: {
			"InternationalCloud": CLOUD_ACCESS_ENV_HUAWEI_GLOBAL,
			"ChinaCloud":         CLOUD_ACCESS_ENV_HUAWEI_CHINA,
		},
		CLOUD_PROVIDER_ALIYUN: {
			"InternationalCloud": CLOUD_PROVIDER_ALIYUN,
			"FinanceCloud":       CLOUD_ACCESS_ENV_ALIYUN_FINANCE,
		},
	}
)

func GetCloudEnv(provider, accessUrl string) string {
	envMap, ok := CLOUD_ENV_MAP[provider]
	if !ok {
		return provider
	}
	env, ok := envMap[accessUrl]
	if !ok {
		return provider
	}
	return env
}
