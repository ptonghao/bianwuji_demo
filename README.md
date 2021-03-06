# bianwuji_demo

一、需求分析:
    快速查询设备数据

二、解决方案:
    1、业务数据冷热分离,业务热数据放内存,定时更新
    2、兜底方案:未命中内存,查db
    3、业务热数据量较大的放分布式redis,设置过期时间并定期更新
    4、缓存失效时服务降级

三、方案缺点:
    1、业务热数据量大时,在更新重刷内存或者缓存时,所有请求会打到db
    2、内存里的数据必须都是有序的

四、模块逻辑说明:
    1、当前数据放在变量里
    2、前一天的数据放到内存(map)里,每天0点的时候定时更新.内存结构为slice+map,slice存采集数据的时间点,map存slice的下标和采集的数据, 
    3、上个月的数据打算放分布式redis里,但是不好测试所以放入go-cache里.每个月第一天0点定时更新
    4、代码里所有 // todo ... 的地方都是mock的数据
    5、使用时间戳查询,对slice进行二分查找搜索范围,找的最接近目标值的下标,然后使用下标到map里查询设备数据
    
五、调用方法:
    1、restful api 方式:
        接口: ip:8888/query?product=t&dim=day&begin_time=XXX&end_time=XXX
        调用方式: POST
        参数说明:
            product:    产品类型,t: 温度计温度, p: 机械臂位置, pr: 酶标仪
            dim:        查询维度, current: 当前, day: 昨天, month: 上个月
            begin_time: 查询开始时间戳
            end_time:   查询结束时间戳

六、测试方案:
    1、可以使用restful api 调用接口进行测试
    2、模块单测
        单测月级别数据: models/page/cache/month/month_test.go
        单测天级别数据: models/page/cache/yesterday/yesterday_test.go

七、目录结构说明:
    1、models/page 包里放业务逻辑代码
        1、cache 包:  缓存
        2、consts 包: 常量定义
        3、service包: 产品工厂(温度计温度、机械臂位置、酶标仪)
    2、library 包 辅助工具
    3、bootstrap 包 gin简单路由
    4、bootstrap/scheduler 定时任务
    5、conf 包 环境等配置相关

八、详细设计:
    1、中文文档: doc/边无际详设.docx
    2、英文文档: doc/Edgenesis detail design.docx
    


