# ucgo
go框架

go的web框架，集成了热重启，gin路由，rsa+des+sign标准的接口加密方式，各种中间件可自行扩展，有mongo、redis的链接示例、网上go使用以太坊的教程也比较少都是web3的，集成了TokenERC20的示例

##需求

####软件

* dep

####dep
* name = "github.com/sirupsen/logrus"
  version = "1.0.6"
* name = "gopkg.in/mgo.v2"
  branch = "v2"
* name = "github.com/ethereum/go-ethereum"
  version = "1.8.13"
* name = "github.com/go-redis/redis"
  version = "6.13.2"
* name = "github.com/gorilla/websocket"
  version = "1.2.0"
* name = "github.com/gin-gonic/gin"
  version = "1.3.0"

####安装

dep ensure

###目录结构
<pre>

+ app                    
  |+ controller
    |+ demo
      |- controller.go
    .......
  |+ data  
    |+ http
      |+ httpbase
        |- base.go
    |+ mongo
      |+ mongos
      |+ replset
    |+ redis
      |+ cluster
  |+ model
    |+ demo
      |- code.go
      |- config.yaml
      |- getconf.go
      |- model.go
      |- service.go
      |- util.go
    .......
  |+ route
    |- inner.go
    |- outer.go
    |- script.go
  |+ script
    |+ demo
      |- script.go
    .......
+ library
  |+ icode
     |- code.go
  |+ ilogger
     |- logrus.go
  |+ imongo
     |- mongos.go
     |- replset.go
  |+ iredis
     |- cluster.go
  |+ ires
     |- response.go
  |+ isafe
    |+ keystore
      |+ v1
        |- private_android.key
        |- private_ios.key
        |- public_android.key
        |- public_ios.key
      |+ v2
     |- login.go
     |- network.go
     |- post.go
     |- sign.go
   |+ iutil
     |- des.go
     |- device.go
     |- md5.go
     |- params.go
     |- rsa.go
     |- time.go
     |- version.go
- .gitignore
- build.sh                    //编译脚本
- Gopkg.lock                         
- Gopkg.toml                
- README.md               
- run.sh                      //启动脚本包含编译跟服务启动
- service.sh                  //服务启动脚本
- ucgo.go                     //框架入口

</pre>
