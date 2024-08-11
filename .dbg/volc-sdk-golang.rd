# 调试
{
    alias dlv='CGO_ENABLED=0 dlv test github.com/volcengine/volc-sdk-golang/example/rtc/v20230801 --init .dbg/volc-sdk-golang.dlv -- -test.run Test_GetRoomOnlineUsers -test.v'
}
