如果长连接中,一端依旧工作,另一端程序崩溃,会出现什么问题
    
    2018/11/24 16:41:44 listen at :8080
    read: 4 element err %!s(<nil>) data ping      
    read: 0 element err read tcp 127.0.0.1:8080->127.0.0.1:58515: read: connection reset by peer data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello
    325 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58515: write: broken pipe data hello


    2018/11/24 16:42:23 listen at :8080
    read: 4 element err %!s(<nil>) data ping      
    read: 0 element err read tcp 127.0.0.1:8080->127.0.0.1:58521: read: connection reset by peer data           
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: protocol wrong type for socket data hello
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    read: 0 element err EOF data           
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: broken pipe data hello
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: broken pipe data hello
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: broken pipe data hello
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: broken pipe data hello
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: broken pipe data hello
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: broken pipe data hello
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: broken pipe data hello
    340 write: 0 element err write tcp 127.0.0.1:8080->127.0.0.1:58521: write: broken pipe data hello
    
解析下(随意点击运行几次):

    读取:
        首先读取的4个字节
        接下来,报错说 connection reset  
        然后报错说 EOF   
    写入:
        写入340个字节,然后报错说 broken pipe
        或者写入325个字节,然后报错说 broken pipe
        
分析:
    写入,是写入系统的缓冲栈中, 然后估计是当客户端断开连接后,发送broken pipe 
    (写入多少,跟对端什么时候发送broken pipe 有关,但是不管写入多少,对方肯定是都没有收到的)
    
    读取, 先告知 connection reset, 然后告知 EOF        
        