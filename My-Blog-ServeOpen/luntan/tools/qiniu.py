import uuid
import io

import qiniu
from PIL import Image
from qiniu import BucketManager

from luntan.settings import QINIU_ACCESS_KEY, QINIU_SECRET_KEY, QINIU_BUCKET_NAME, QINIU_DOMAIN

access_key = QINIU_ACCESS_KEY
secret_key = QINIU_SECRET_KEY
bucket_name = QINIU_BUCKET_NAME
domain = QINIU_DOMAIN
# 初始化 Auth 对象
q = qiniu.Auth(access_key=access_key, secret_key=secret_key)


def upload_to_qiniu(image_file, key=None):
    """上传图片到七牛云并返回图片 URL"""
    try:
        key = 'images/' + str(uuid.uuid1()).replace('-', '')
        _img = image_file.read()
        image = Image.open(io.BytesIO(_img))
        # 获取image的后缀名
        name = 'upfile.{0}'.format(image.format)
        image.save('./' + name)
        path = './' + name

        # 生成上传 Token，可以指定过期时间等
        token = q.upload_token(bucket=bucket_name, key=key, expires=3600)

        # 调用 put_file 方法上传图片，返回图片的哈希值和信息
        ret, info = qiniu.put_file(token, key, path)

        # 判断图片是否上传成功
        if info.status_code == 200:
            # 返回图片的访问 URL，拼接方式：http://domain/key
            return f'{domain}{ret["key"]}'
        else:
            print(info.exception())
    except Exception as e:
        print(e)


def delete_image(image_url):
    try:
        q = qiniu.Auth(access_key=access_key, secret_key=secret_key)
        bucket = qiniu.BucketManager(q)
        image_key = image_url.replace(f'{domain}', '')
        ret, info = bucket.delete(bucket_name, image_key)
        if info.status_code == 200:
            return True
        else:
            return False
    except Exception as error:
        print(error)