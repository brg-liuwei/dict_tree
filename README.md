#dict_tree

### 字典树

一个简单的字典树

### 功能

用于对一系列关键字建立字典树，逐行处理文本，对每一条记录打标签

### 依赖

需要安装golang

### 示例

    git clone https://github.com/brg-liuwei/dict_tree
    cd dict_tree
    make
    bin/dict

### 数据集

`data/cities`: 该文件是标签文件，其中每一行的第一列是城市名，第二列是标签
`data/records`: 该文件的每一行是待处理的文本

### 测试示例

    make test

### 性能测试

    make bench

