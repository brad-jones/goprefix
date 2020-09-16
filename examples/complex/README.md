# Complex

This example demonstrates how the `prefixer` & `colorchooser` can be used together.
This is more or less what `goexec` is doing in this example
<https://github.com/brad-jones/goexec/tree/v2/examples/prefixed>

## Expected Output

_Depending on OS & network conditions of course_

```
ping 127.0.0.1 |
ping 127.0.0.1 | Pinging 127.0.0.1 with 32 bytes of data:
ping 127.0.0.1 | Reply from 127.0.0.1: bytes=32 time<1ms TTL=128
ping 127.0.0.2 |
ping 127.0.0.2 | Pinging 127.0.0.2 with 32 bytes of data:
ping 127.0.0.2 | Reply from 127.0.0.2: bytes=32 time<1ms TTL=128
ping 127.0.0.2 | Reply from 127.0.0.2: bytes=32 time<1ms TTL=128
ping 127.0.0.1 | Reply from 127.0.0.1: bytes=32 time<1ms TTL=128
ping 127.0.0.2 | Reply from 127.0.0.2: bytes=32 time<1ms TTL=128
ping 127.0.0.1 | Reply from 127.0.0.1: bytes=32 time<1ms TTL=128
ping 127.0.0.2 | Reply from 127.0.0.2: bytes=32 time<1ms TTL=128
ping 127.0.0.1 | Reply from 127.0.0.1: bytes=32 time<1ms TTL=128
ping 127.0.0.1 |
ping 127.0.0.1 | Ping statistics for 127.0.0.1:
ping 127.0.0.1 | Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
ping 127.0.0.1 | Approximate round trip times in milli-seconds:
ping 127.0.0.1 | Minimum = 0ms, Maximum = 0ms, Average = 0ms
ping 127.0.0.2 |
ping 127.0.0.2 | Ping statistics for 127.0.0.2:
ping 127.0.0.2 | Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
ping 127.0.0.2 | Approximate round trip times in milli-seconds:
ping 127.0.0.2 | Minimum = 0ms, Maximum = 0ms, Average = 0ms
```
