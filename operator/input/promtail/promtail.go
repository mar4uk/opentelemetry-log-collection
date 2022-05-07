package promtail

import (
	// "github.com/grafana/loki/clients/pkg/promtail/positions"
	// "github.com/grafana/loki/clients/pkg/promtail/scrapeconfig"
	// "github.com/grafana/loki/clients/pkg/promtail/targets/file"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-log-collection/operator"
	"github.com/open-telemetry/opentelemetry-log-collection/operator/helper"
)

func init() {
	operator.Register("promtail_input", func() operator.Builder { return NewPromtailInputConfig("") })
}

func NewPromtailInputConfig(operatorID string) *PromtailInputConfig {
	return &PromtailInputConfig{
		InputConfig: helper.NewInputConfig(operatorID, "promtail_input"),
	}
}

// PromtailInputConfig is the configuration of a journald input operator
type PromtailInputConfig struct {
	helper.InputConfig `mapstructure:",squash" yaml:",inline"`

	PromtailConfig          PromtailConfig `mapstructure:"config"`
}

type PromtailConfig struct {
	// PositionsConfig positions.Config      `yaml:"positions,omitempty"`
	// ScrapeConfig    []scrapeconfig.Config `yaml:"scrape_configs,omitempty"`
	// TargetConfig    file.Config           `yaml:"target_config,omitempty"`
}

// Build will build a promtail input operator from the supplied configuration
func (c PromtailInputConfig) Build(logger *zap.SugaredLogger) (operator.Operator, error) {
	inputOperator, err := c.InputConfig.Build(logger)
	if err != nil {
		return nil, err
	}

	// TODO: initialize Promtail Managers here

	// return &PromtailInput{
	// 	InputOperator: inputOperator,
	// }, nil

	return &PromtailInput{
		InputOperator: inputOperator,
	}, nil
}

type PromtailInput struct {
	helper.InputOperator
}

func (operator *PromtailInput) Start(_ operator.Persister) error {
	panic("not implemented")
	// TODO: start Promtail Managers here
	// NB: Persister is a cursor in input where we stopped last time
}

func (operator *PromtailInput) Stop() error {
	panic("not implemented")
	// TODO: stop Promtail Managers here
}