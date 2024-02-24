package snowball

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	// Length of the timestamp section.
	TimestampLen uint8 = 42
	// Length of the Server ID section.
	ServerIdLen uint8 = 11
	// Length of the sequence section.
	SequenceLen uint8 = 11
	// Maximum value for the sequence section.
	MaxSequence uint16 = 1<<SequenceLen - 1
	// Maximum value for the Server ID section.
	MaxServerId uint16 = 1<<ServerIdLen - 1
	// Maximum value for the timestamp section.
	MaxTimestamp uint64 = 1<<TimestampLen - 1

	serverIdShift  = SequenceLen
	timestampShift = ServerIdLen + SequenceLen
)

var (
	ServerId  uint64    = GetServerId()
	Epoch     uint64    = GetEpoch()
	EpochTime time.Time = time.Unix(int64(Epoch)/1000, (int64(Epoch)%1000)*1000000)
)

// Contains basic information used to generate Snowball IDs
type SnowballNode struct {
	mutex sync.Mutex
	epoch time.Time

	serverId uint64
	currTime uint64
	currSeq  uint64
}

// Creates and returns a new node object for generating Snowball IDs
func InitNode() (*SnowballNode, error) {
	if SequenceLen+ServerIdLen > 22 {
		return nil, errors.New("initialization failed: sequence and server ID length is invalid")
	}

	result := SnowballNode{}
	result.serverId = ServerId
	if result.serverId > uint64(MaxServerId) {
		return nil, errors.New(
			"initialization failed: server ID must be between 0 and " + strconv.FormatInt(int64(MaxServerId), 10),
		)
	}

	var now = time.Now()
	result.epoch = now.Add(EpochTime.Sub(now))
	return &result, nil
}

// Creates and returns a new, unique Snowball ID.
func (node *SnowballNode) GenerateID() uint64 {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	now := time.Since(node.epoch).Milliseconds()
	if now == int64(node.currTime) {
		node.currSeq = (node.currSeq + 1) & uint64(MaxSequence)
		if node.currSeq == 0 {
			for now <= int64(node.currTime) {
				now = int64(time.Since(node.epoch).Milliseconds())
			}
		}
	} else {
		node.currSeq = 0
	}

	node.currTime = uint64(now)
	result := uint64(
		now<<int64(timestampShift) | (int64(node.serverId) << int64(serverIdShift)) | int64(node.currSeq),
	)
	return result
}