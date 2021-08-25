log封装的

写入日志底层都是封装 log.Logger 

------------------writer 文件---------------------------------

主体接口LevelWriter 包括三种接口，FormatWriter格式化接口，io.Writer，以及实现WriteLevel

defaultFormatTimestamp、defaultFormatLevel、defaultFormatCaller、defaultFormatMessage 是默认的格式化函数

levelWriterAdapter levelWriter适配器，非主体接口LevelWriter的io.Writer可以通过该结构体封装

multiLevelWriter LevelWriter数组，可以将一个消息循环写入到不同的LevelWriter数组中

------------------access_writer 文件-----------------------

AccessWriter 是一种LevelWriter， 其中补充了日志清理，循环删除等操作，更加丰富。
所以不管是file 还是stdout直接组合AccessWriter 并重写其中的自定义函数就可以


------------------global 文件------------------------------

用来保存全局信息

------------------util 文件--------------------------------

公用功能函数