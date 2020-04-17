package main

/*
在streamserver中Streaming和Upload files都需要保持长连接，
和之前的api短连接是不一样的。所以在多路长连接同时保持的时候就会出现一个问题，
如果再不断的发起连接，打开视频，总有一个时刻会把我们的server crash掉，因此我们需要
一个流控部分来控制connection
*/
