import dashscope
import subprocess
from dashscope import Generation
from http import HTTPStatus

diffResult = subprocess.run(['git', 'diff', '--staged' ], stdout=subprocess.PIPE).stdout.decode('utf-8')
data = '根据git diff 的结果, 生成提交信息, 提交信息要满足一下要求: 1. 符合约定式提交的格式, 如: <type>[optional scope]: <description>; 2. 英文; 3. 一行. \n git diff 结果如下:\n' + diffResult

print(data)
response = Generation.call(
    model='qwen-v1',
    prompt=data
)
# The response status_code is HTTPStatus.OK indicate success,
# otherwise indicate request is failed, you can get error code
# and message from code and message.
if response.status_code == HTTPStatus.OK:
    print(response.output)  # The output text
    print(response.usage)  # The usage information
else:
    print(response.code)  # The error code.
    print(response.message)  # The error message.
