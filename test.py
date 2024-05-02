import requests
import json

def upload_image(file_path):
    url = 'http://localhost:5000/upload'
    files = {'file': open(file_path, 'rb')}
    response = requests.post(url, files=files)
    text=response.text
    print(text)
    print(json.loads(text))
# 提交图片示例
image_path = 'a.png'
upload_image(image_path)