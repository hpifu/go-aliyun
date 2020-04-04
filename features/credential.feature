Feature: echo 测试

    Scenario: echo
        When http 请求 POST /credential
            """
            {
                "json": {
                    "filename": "test1",
                    "accessKeyID": "9ede1245e77c06df5b",
                    "accessKeySecret": "36ae8f74164a5af0f66b9c25a6146583"
                }
            }
            """
        Then http 检查 200
        When http 请求 GET /credential/test1
        Then http 检查 200
            """
            {
                "accessKeyID": "9ede1245e77c06df5b",
                "accessKeySecret": "36ae8f74164a5af0f66b9c25a6146583"
            }
            """