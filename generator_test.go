package idgen

import (
    "github.com/stretchr/testify/assert"
    "sync"
    "testing"
)

func TestGenerator_GenId(t *testing.T) {
    g := NewGenerator(1,1)
    count := 8000
    ch := make(chan uint64, count)
    wg := sync.WaitGroup{}
    wg.Add(count)
    defer close(ch)

    for i :=0; i < count; i++ {
        go func() {
            defer wg.Done()
            id, err := g.GenId()
            assert.Nil(t, err)
            ch <- id
        }()
    }

    wg.Wait()

    m := map[uint64]int{}
    for i:=0; i<count;i++ {
        id := <-ch
        _, ok := m[id]
        if ok {
            t.Logf("repeat id %d", id)
            t.Fail()
        }

        m[id]=1
    }
    t.Log(len(m))
    assert.Equal(t, len(m), count)

}
