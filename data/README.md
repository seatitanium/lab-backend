# data

此页面用来存放后端可能会调用的数据，类似于本地数据库。这些数据会用 Python 进行处理。运行这些 Python 之前，请先初始化 virtual environment 后进入

```shell
python -m venv .
chmod +x ./bin/*
source ./bin/activate
```

安装依赖

```shell
pip install -r ./requirements.txt
```

后运行脚本。脚本内部分内容需要修改。对于包含阿里云 api 调用的操作，需要创建 `secret.py` 并填入 AccessKeyID 和 AccessKeySecret

```python
akid='' #...
aksecret='' #...
```