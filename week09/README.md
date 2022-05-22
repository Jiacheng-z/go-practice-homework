# Socket粘包解法

1. length field based frame decoder

## 1. fix length

每次发送定长的数据，接收方，每次接收定长数据。

## 2. delimiter based

发送方在数据结尾添加指定分割符，例如HTTP的/r/n来区分包结束。 接收方接到特定字符后，则知道包结束了。

## 3. length field based frame decoder

在数据包头信息中加入body长度信息。例: Netty, http