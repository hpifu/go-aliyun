Feature: credential 测试

    Scenario: credential
        When http 请求 POST /endpoint/imm
            """
            {
                "json": {
                    "endpoint": "https://imm.cn-hangzhou.aliyuncs.com"
                }
            }
            """
        Then http 检查 200
        When http 请求 POST /endpoint/imm
            """
            {
                "json": {
                    "endpoint": "https://imm.cn-beijing.aliyuncs.com"
                }
            }
            """
        Then http 检查 200
        When http 请求 GET /endpoint/imm
        Then http 检查 200
            """
            {
                "json": [
                    "https://imm.cn-beijing.aliyuncs.com",
                    "https://imm.cn-hangzhou.aliyuncs.com"
                ]
            }
            """
        When http 请求 DELETE /endpoint/imm
            """
            {
                "json": {
                    "endpoint": "https://imm.cn-beijing.aliyuncs.com"
                }
            }
            """
        Then http 检查 202
        When http 请求 DELETE /endpoint/imm
            """
            {
                "json": {
                    "endpoint": "https://imm.cn-hangzhou.aliyuncs.com"
                }
            }
            """
        Then http 检查 202