# 编译调试
{
    alias gob='CGO_ENABLED=0 go build -v -gcflags "all=-N -l" -o volc-sdk-golang ./cmd'
    alias dlv='gob && dlv exec ./volc-sdk-golang --init=.dbg/volc-sdk-golang.dlv'
}

# 单元测试
{
    alias dlv='CGO_ENABLED=0 dlv test github.com/volcengine/volc-sdk-golang/example/rtc/v20230801 --init .dbg/volc-sdk-golang.dlv -- -test.run Test_GetRoomOnlineUsers -test.v'
}

# 上传云服务
{
    scp /home/qstesiro/github.com/qstesiro/volc-sdk-golang/volc-sdk-golang root@101.126.87.161:/root/
    scp /home/qstesiro/.volc/config root@101.126.87.161:/root/.volc/
}
