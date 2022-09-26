package types

import (
	"fmt"

	pbacme "github.com/starknet-graph/firehose-starknet/types/pb/zklend/starknet/type/v1"
	"github.com/streamingfast/bstream"
	pbbstream "github.com/streamingfast/pbgo/sf/bstream/v1"
	"google.golang.org/protobuf/proto"
)

func BlockFromProto(b *pbacme.Block) (*bstream.Block, error) {
	content, err := proto.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal to binary form: %s", err)
	}

	block := &bstream.Block{
		Id:             b.ID(),
		Number:         b.Number(),
		PreviousId:     b.PreviousID(),
		Timestamp:      b.Time(),
		LibNum:         b.Number() - 1,
		PayloadKind:    pbbstream.Protocol_UNKNOWN,
		PayloadVersion: 1,
	}

	return bstream.GetBlockPayloadSetter(block, content)
}
