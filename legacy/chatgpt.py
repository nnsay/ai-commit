import openai
import os

openai.organization = "org-icOCl0ze4WZzF0OoJp1kLc7j"
openai.api_key = os.getenv("OPENAI_API_KEY")

# List llm
# models = openai.Model.list()
# print(models)

# Call the openai ChatCompletion endpoint
response = openai.ChatCompletion.create(
    model="gpt-3.5-turbo",
    messages=[{"role": "user", "content": "Hello World!"}],
)
# Extract the response
print(response["choices"][0]["message"]["content"])
