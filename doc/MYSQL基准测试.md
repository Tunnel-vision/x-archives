### MYSQL 基准测试
***

基准测试（benchmark) 事针对系统设计的压力测试
sysbench mysql基准测试工具
***
2.1 为什么需要基准测试
***
2.2 基准测试的策略  
    * 针对整个系统的测试 集成式测试  
    * 单独测试，单组件式测试
    2.2.1 确定测试何种指标  
    * 吞吐量  
    * 响应时间延迟  
    * 并发性  
    * 可扩展性  
***
2.3 基准测试的方法
    2.3.1 设计和规划基准测试
    2.3.2 基准测试应该运行多长时间  
    2.3.3 获取系统性能和状态  
*** 
2.4基准测试工具
   2.4.1 集成测试工具  
   ab   
   http_load  
   JMeter  
   2.4.2 单组件测试工具  
     mysqlslap  
     Mysql Benchmark Suite（sql-bench）  
     super Smack  
     Database Test Suite
     sysbench
     
     