package idgen

import (
    "errors"
    "sync"
    "sync/atomic"
    "time"
)

type IGenerator interface {
    GenId() (uint64, error)
}

type generator struct {
    mu        sync.Mutex
    LastStamp int64
    AppId     int64
    MachineId int64
    Sequence  int64
}


func NewGenerator(appId, machineId int64) IGenerator {
    g := &generator{
        mu: sync.Mutex{},
        LastStamp: 0,
        AppId: appId,
        MachineId: machineId,
        Sequence: 0,
    }
    return g
}


func (g *generator)getMilliSecond() int64 {
    return time.Now().UnixNano() / 1e6
}


func (g *generator)GenId()(uint64, error) {
    g.mu.Lock()
    defer g.mu.Unlock()
    return g.genId()
}


func (g *generator)genId() (uint64, error) {
    ts := g.getMilliSecond()
    if ts < g.LastStamp {
        return 0, errors.New("时间戳异常")
    }

    if g.LastStamp == ts {
        atomic.AddInt64(&g.Sequence,1)
        g.Sequence = (g.Sequence + 1) & maxSequence

        if g.Sequence == 0 {
            for ts <= g.LastStamp {
                g.LastStamp = g.getMilliSecond()
            }
        }
    } else {
        g.Sequence = 0
    }

    g.LastStamp = ts
    id := ((ts - startTime) << timeLeft) | (g.AppId << appLeft) | (g.MachineId << machineLeft) | g.Sequence

    return uint64(id), nil
}