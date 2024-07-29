# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
import os
import sys

from typing import List

from alibabacloud_tea_openapi.client import Client as OpenApiClient
from alibabacloud_tea_openapi import models as open_api_models
from alibabacloud_tea_util import models as util_models
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from datetime import datetime, timedelta
from dateutil.relativedelta import *
import json
from secret import akid, aksecret

product_code="yundisk" # yundisk, oss, ecs
subscription_type="PayAsYouGo" # PayAsYouGo, Subscription

class Sample:
    def __init__(self):
        pass

    @staticmethod
    def create_client() -> OpenApiClient:
        """
        使用AK&SK初始化账号Client
        @return: Client
        @throws Exception
        """
        # 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
        # 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378659.html。
        config = open_api_models.Config(
            # 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。,
            access_key_id=akid,
            # 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。,
            access_key_secret=aksecret
        )
        # Endpoint 请参考 https://api.aliyun.com/product/BssOpenApi
        config.endpoint = f'business.aliyuncs.com'
        return OpenApiClient(config)

    @staticmethod
    def create_api_info() -> open_api_models.Params:
        """
        API 相关
        @param path: params
        @return: OpenApi.Params
        """
        params = open_api_models.Params(
            # 接口名称,
            action='QueryBill',
            # 接口版本,
            version='2017-12-14',
            # 接口协议,
            protocol='HTTPS',
            # 接口 HTTP 方法,
            method='POST',
            auth_type='AK',
            style='RPC',
            # 接口 PATH,
            pathname=f'/',
            # 接口请求体内容格式,
            req_body_type='json',
            # 接口响应体内容格式,
            body_type='json'
        )
        return params

    @staticmethod
    def main(
        args: List[str],
    ) -> None:
        startCycle = datetime.strptime("2021-01", "%Y-%m")
        endCycle = datetime.strptime("2024-07", "%Y-%m")

        result = []

        while endCycle.timestamp() > startCycle.timestamp():
            startCycle += relativedelta(months=+1)
            cycle = startCycle.strftime("%Y-%m")
            for i in range(10):
                client = Sample.create_client()
                params = Sample.create_api_info()
                # query params
                queries = {}
                queries['BillingCycle'] = cycle
                queries['ProductCode'] = product_code
                queries['ProductType'] = None
                queries['SubscriptionType'] = subscription_type
                queries['IsHideZeroCharge'] = True
                queries['PageNum'] = i+1
                queries['PageSize'] = 300
                # runtime options
                runtime = util_models.RuntimeOptions()
                request = open_api_models.OpenApiRequest(
                    query=OpenApiUtilClient.query(queries)
                )
                # 复制代码运行请自行打印 API 的返回值
                # 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
                resp = client.call_api(params, request, runtime)
                print(f"Retrieved result cycle={cycle},message={resp['body']['Message']}")
                result.extend(resp["body"]["Data"]["Items"]["Item"])
                print(f"result length={len(result)}")

        with open(product_code + "-bills.json", "w") as f:
            json.dump(result, f)

    @staticmethod
    async def main_async(
        args: List[str],
    ) -> None:
        client = Sample.create_client()
        params = Sample.create_api_info()
        # query params
        queries = {}
        queries['BillingCycle'] = '2021-02'
        queries['ProductCode'] = 'ecs'
        queries['ProductType'] = None
        queries['SubscriptionType'] = 'PayAsYouGo'
        queries['IsHideZeroCharge'] = True
        queries['PageNum'] = 1
        queries['PageSize'] = 300
        # runtime options
        runtime = util_models.RuntimeOptions()
        request = open_api_models.OpenApiRequest(
            query=OpenApiUtilClient.query(queries)
        )
        # 复制代码运行请自行打印 API 的返回值
        # 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
        await client.call_api_async(params, request, runtime)


if __name__ == '__main__':
    Sample.main(sys.argv[1:])
