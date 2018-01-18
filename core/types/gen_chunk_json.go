// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/common/hexutil"
)

var _ = (*chunkMarshalling)(nil)

func (c Chunk) MarshalJSON() ([]byte, error) {
	type Chunk struct {
		Index hexutil.Uint64 `json:"index"  gencodec:"required"`
		Data  hexutil.Bytes  `json:"bytes"  gencodec:"required"`
		Proof common.Hash    `json:"proof"  gencodec:"required"`
	}
	var enc Chunk
	enc.Index = hexutil.Uint64(c.Index)
	enc.Data = c.Data
	enc.Proof = c.Proof
	return json.Marshal(&enc)
}

func (c *Chunk) UnmarshalJSON(input []byte) error {
	type Chunk struct {
		Index *hexutil.Uint64 `json:"index"  gencodec:"required"`
		Data  hexutil.Bytes   `json:"bytes"  gencodec:"required"`
		Proof *common.Hash    `json:"proof"  gencodec:"required"`
	}
	var dec Chunk
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Index == nil {
		return errors.New("missing required field 'index' for Chunk")
	}
	c.Index = uint64(*dec.Index)
	if dec.Data == nil {
		return errors.New("missing required field 'bytes' for Chunk")
	}
	c.Data = dec.Data
	if dec.Proof == nil {
		return errors.New("missing required field 'proof' for Chunk")
	}
	c.Proof = *dec.Proof
	return nil
}
