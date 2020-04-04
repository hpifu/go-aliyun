# Feature: credential 测试

#     Scenario: credential
#         When http 请求 POST /credential
#             """
#             {
#                 "json": {
#                     "filename": "hatlonely",
#                     "accessKeyID": "key",
#                     "accessKeySecret": "secret"
#                 }
#             }
#             """
#         Then http 检查 200
#         When http 请求 POST /imm
#             """
#             {
#                 "json": {
#                     "endpoint": "https://imm.cn-hangzhou.aliyuncs.com",
#                     "credential": "hatlonely",
#                     "params": {
#                         "Action": "ListProjects"
#                     }
#                 }
#             }
#             """
#         Then http 检查 200
