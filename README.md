# C-OCR
An online optical character recognition (OCR) website based on the gin framework, using the rpc architecture to call remote methods to achieve OCR.  

一个基于gin框架的在线图片文字识别（OCR）网站，用rpc架构调用远程方法实现OCR。

## 部署C-OCR：

#### 一、部署并启动paddleOCR-Server
见 [heejinzzz/paddleOCR-Server](https://github.com/heejinzzz/paddleOCR-Server)

#### 二、下载代码
    git clone https://github.com/heejinzzz/C-OCR
    cd C-OCR

#### 三、使用consul
安装并启动consul，将 paddleOCR-Server 提供的 paddleOCR 服务注册进consul中。

#### 四、使用redis
安装并启动redis。

#### 五、使用nginx
安装nginx，修改 nginx.conf 文件后，用其替换nginx安装位置下的 conf/nginx.conf 文件，启动nginx。

#### 六、修改源码中的参数：
1. 修改 paddleOCRclient/paddleOCRclient.go 文件中的 consulAddress
2. 修改 redisDB/redisDB.go 文件中的 redisServerAddress
3. 修改 main.go 文件中的 C_OCRAddress

#### 七、启动C-OCR
执行：

    go run main.go
此时在浏览器输入在nginx配置的地址，即可访问C-OCR。
