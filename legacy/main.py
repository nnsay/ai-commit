import dashscope
import subprocess
import sys
import json
from dashscope import Generation
from http import HTTPStatus

diffResult = subprocess.run(['git', 'diff', '--staged' ], stdout=subprocess.PIPE).stdout.decode('utf-8')
if (diffResult == ''):
  print('Please run git add something into stage firstly.')
  sys.exit(0)

history=[]
maxRetry=5
maxMessageWords=80
#data = '根据 git diff 的结果, 生成提交信息, 提交信息要满足以下要求: 1. 符合约定式提交的格式, 如: <type>[optional scope]: <description>; 2. 英文; 3. 一行. git diff 结果如下: \n' + diffResult
data = '根据git diff结果生成一句英文的提交信息, git diff结果如下: \n' + diffResult


print(data)

response = Generation.call(
    model='qwen-v1',
    prompt=data
)
history.append({"user": data, "bot": response.output.text})
# The response status_code is HTTPStatus.OK indicate success,
# otherwise indicate request is failed, you can get error code
# and message from code and message.
aiCommitMessage = ''
if response.status_code == HTTPStatus.OK:
    # print('>', json.dumps(response.output, indent=4, ensure_ascii=False))
    aiCommitMessage = response.output.text
    while len(response.output.text) > maxMessageWords and maxRetry > 0:
      data = '请再精简一些, 结果最好不超过{num}个单词的英文描述'.format(num=maxMessageWords)
      response = Generation.call(
        model='qwen-v1',
        prompt=data,
        history=history
      )
      history.append({"user": data, "bot": response.output.text})
      # print('>', json.dumps(response.output, indent=4, ensure_ascii=False))
      maxRetry -= 1

else:
    print('Code: %d, status: %s, message: %s' % (response.status_code, response.code, response.message))
    sys.exit(1)

print(aiCommitMessage)
