package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

var options *idgen.IdGeneratorOptions

func init() {
	// 自定义参数
	workerId := uint16(global.GVA_CONFIG.System.WorkerID)
	options = idgen.NewIdGeneratorOptions(workerId) 	// 设置WorkerId = 1
	options.WorkerIdBitLength = 15                      // workerId长度，最多 2^15 = 32768
	options.SeqBitLength = 6                            // 同一毫秒最多59个顺序号，等于3时最多3个，4时最多11个，5时最多27，6时最多59
	options.BaseTime = time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC).UnixNano() / 1e6
	idgen.SetIdGenerator(options)
}

// GetWorkerId 机器ID
func GetWorkerId(sid uint64) (workerId uint16) {
	maxWorkerIdNumber := uint16(1<<options.WorkerIdBitLength) - 1
	workerId = uint16(sid>>options.SeqBitLength) & maxWorkerIdNumber
	return
}

// GetGenTimestamp 获取创建ID时的时间戳
func GetGenTimestamp(sid uint64) (timestamp uint64) {
	timestamp = GetTimestamp(sid) + uint64(options.BaseTime)
	return
}

// GetGenTime 获取创建ID时的时间字符串(精度：秒)
func GetGenTime(sid uint64) (t string) {
	// 需将GetGenTimestamp获取的时间戳/1000转换成秒
	t = time.Unix(int64(GetGenTimestamp(sid))/1000, 0).Format("2006-01-02 15:04:05")
	return
}

// GetTimestamp 获取时间戳
func GetTimestamp(sid uint64) (timestamp uint64) {
	timestampShift := options.WorkerIdBitLength + options.SeqBitLength
	timestampMax := uint64(1<<timestampShift) - 1
	timestamp = (sid >> timestampShift) & timestampMax
	return
}

//export SetOptions
func SetOptions(workerId uint16) {
	// 创建 IdGeneratorOptions 对象，请在构造函数中输入 WorkerId：
	options = idgen.NewIdGeneratorOptions(workerId)
	// 保存参数（必须的操作，否则以上设置都不能生效）
	idgen.SetIdGenerator(options)
	// 以上初始化过程只需全局一次，且必须在第2步之前设置。//第2步，生成ID：
}

//export NextId
func NextId() uint64 {
	return idgen.NextId()
}
