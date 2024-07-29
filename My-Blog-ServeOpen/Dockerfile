# 使用官方的 Python 运行时作为父镜像
FROM python:3.10-slim
RUN python -m pip install --upgrade pip
# 设置工作目录
WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends default-libmysqlclient-dev gcc && apt-get clean

# 设置环境变量

RUN pip config --user set global.index https://mirrors.huaweicloud.com/repository/pypi
RUN pip config --user set global.index-url https://mirrors.huaweicloud.com/repository/pypi/simple
RUN pip config --user set global.trusted-host mirrors.huaweicloud.com
# 复制 requirements.txt 到容器中
COPY requirements.txt .

# 安装依赖
RUN pip install --no-cache-dir -r requirements.txt

# 复制项目文件到容器中
COPY . /app/


# 暴露端口
EXPOSE 8000

# 运行命令
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]