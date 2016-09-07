package loggly

import (
	"encoding/json"
	"fmt"
    "net/http"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/op"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
    
)

func init() {
	outputs.RegisterOutputPlugin("loggly", New)
}

type loggly struct {
	config config
}

func New(beatName string, config *common.Config, _ int) (outputs.Outputer, error) {
	//check for url in config here
    c := &loggly{config: defaultConfig}
    return c, nil
}



// Implement Outputer
func (c *loggly) Close() error {
	return nil
}

func (c *loggly) PublishEvent(
	s op.Signaler,
	opts outputs.Options,
	data outputs.Data,
) error {
	var jsonEvent []byte
	var err error

    jsonEvent, err = json.MarshalIndent(data.Event, "", "  ")
	
    if err != nil {
		logp.Err("Fail to convert the event to JSON (%v): %#v", err, data.Event)
		op.SigCompleted(s)
		return err
	}

	if err = c.writeBuffer(jsonEvent); err != nil {
		goto fail
	}
	op.SigCompleted(s)
	return nil
fail:
	if opts.Guaranteed {
		logp.Critical("Unable to publish events to console: %v", err)
	}
	op.SigFailed(s, err)
	return err
}

func (c *loggly) writeBuffer(buf []byte) error {
	resp, err := http.Post(c.Url, "application/json", &buf)
    if err != nil {
        // handle error
        return err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
	return nil
}
