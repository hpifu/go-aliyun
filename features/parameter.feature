Feature: credential 测试

    Scenario: credential
        When http 请求 POST /parameter/imm/default
            """
            {
                "json": {
                    "filename": "test1",
                    "params": {
                        "Action": "ListProject",
                        "Usage": "ByCU"
                    }
                }
            }
            """
        Then http 检查 200
        When http 请求 GET /parameter/imm/default/test1
        Then http 检查 200
            """
            {
                "json": {
                    "Action": "ListProject",
                    "Usage": "ByCU"
                }
            }
            """
        When http 请求 GET /parameter/imm/default
        Then http 检查 200
            """
            {
                "json": [
                    "test1"
                ]
            }
            """
        When http 请求 DELETE /parameter/imm/default/test1
        Then http 检查 202