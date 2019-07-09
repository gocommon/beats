package worktile

import (
	"bytes"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/outputs/codec"
	"github.com/elastic/beats/libbeat/publisher"
)

var _ outputs.Client = &Worktile{}

// Worktile Worktile
type Worktile struct {
	cfg      config
	beat     beat.Info
	observer outputs.Observer
	codec    codec.Codec
	buff     bytes.Buffer
	ch       chan []byte
}

func (p *Worktile) proc() {
	go func() {
		for{
			select{
			case <- p.ch:
			default:
				if p.buff.Len() == 0 {
					select {
					case <- p.ch:
					}
				}

				// do
				p.buff.Reset()
			}
		}
	}()
}

// Close Close
func (p *Worktile) Close() error {
	panic("not implemented")
}

// Publish Publish
func (p *Worktile) Publish(publisher.Batch) error {
	panic("not implemented")
}

// String String
func (p *Worktile) String() string {
	panic("not implemented")
}

// NewFactory NewFactory
func NewFactory(im outputs.IndexManager, beat beat.Info, stats outputs.Observer, cfg *common.Config) (outputs.Group, error) {
	config := defaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}

	codec, err := codec.CreateEncoder(beat, config.Codec)
	if err != nil {
		return outputs.Fail(err)
	}

	c := &Worktile{
		beat:     beat,
		observer: stats,
		cfg:      config,
		codec:    codec,
		ch : make(chan []byte,100),
	}

	return outputs.Success(-1, 0, c)
}

func init() {
	outputs.RegisterType("worktile", NewFactory)
}
