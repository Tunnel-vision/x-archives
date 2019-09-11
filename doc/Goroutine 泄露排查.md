### Goroutine 泄露排查

我们在发布一个 go 应用时，默认都会启用两个 http handler: 一个是 pprof，方便线上动态追踪问题；另外一个是 prometheus 的 metrics，这样就可以通过 grafana 准实时的监控当前 runtime 信息，及时预警。就像下面这样：
***
```
    package router

    import (
    	"net/http"
    	_ "net/http/pprof"

    	"github.com/prometheus/client_golang/prometheus/promhttp"
    )

    func InitAdmin() {
    	adminRouter := http.DefaultServeMux
    	adminRouter.Handle("/metrics", promhttp.Handler())

        adminServer := &http.Server{
    		Addr:           ":8081",
    		Handler:        adminRouter,
    	}

    	go func() {
    		if err := adminServer.ListenAndServe(); err != nil {
    			println("ListenAndServe admin: ", err.Error())
    		}
    	}()
    }
```


最近我在优化一个 push 服务的时候，便观察到了一个 goroutine 泄露问题：
![](img/img1.png)

测试的客户端仅仅 30 个左右，基本都不会很活跃，但是却看到 goroutine 在持续上涨。怎么查出那些异常的 goroutine 呢？我所知道的有下面几个方法：

1.给程序发送给 **SIGQUIT** 信号，也就是 **kill -3 pid**
2.程序中监听信号，之后通过 runtime.Stack() 获取所有 goroutine 信息
3.通过pprof 获取， /debug/pprof/goroutine

第一种 方法会丢失第一现场，所以一般都在测试阶段使用。
第二种 方法要在程序中提前埋好点，很多开源项目这样做。
第三种 我用的第三种方法，pprof 大家一般都是生成CPU，MEM火焰图来分析性能的问题。殊不知其HTTP
的接口是分析 goroutine 泄露的绝佳工具

***
通过访问 127.0.0.1/debug/pprof/goroutine?debug=1

