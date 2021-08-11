package idgen

const (
    timeBits      = uint64(41) //毫秒级时间戳，占41位，(2^41-1)/(1000 * 60 * 60 * 24 * 365) = 69年
    machineIdBits = uint64(5)  //机器节点占5位，2^5-1=31
    appIdBits     = uint64(5)  //appid占5位，2^5-1=31
    sequenceBits  = uint64(12) //每个机器、每个appid、每毫秒产生的序列号，占12位，2^12-1=4095

    maxAppId     = int64(-1) ^ (int64(-1) << appIdBits)     //appid最大值，防止溢出
    maxMachineId = int64(-1) ^ (int64(-1) << machineIdBits) //machine最大值，防止溢出
    maxSequence  = int64(-1) ^ (int64(-1) << sequenceBits)  //sequence最大值，防止溢出

    machineLeft = uint8(12) //机器节点向左偏移量
    appLeft     = uint8(17) //appid向左偏移量
    timeLeft    = uint8(22) //时间戳向左偏移量

    startTime = int64(1628611200000) //初始时间戳（毫秒）2021-08-11 00:00:00
)
