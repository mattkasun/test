package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kr/pretty"
	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	Role     string
	Duration int
}

func main() {
	ctx := context.Background()
	action := githubactions.New()
	cfg, err := NewFromInput(action)
	if err != nil {
		action.Fatalf(err.Error())
	}
	if err := run(ctx, cfg); err != nil {
		action.Fatalf(err.Error())
	}
}

func NewFromInput(action *githubactions.Action) (*Config, error) {
	context, err := action.Context()
	if err != nil {
		return nil, err
	}
	if !context.Actions {
		return nil, fmt.Errorf("not running in GitHub Actions")
	}
	pretty.Println(context)
	length := action.GetInput("duration")
	d, err := strconv.Atoi(length)
	if err != nil {
		return nil, err
	}
	c := Config{
		Role:     action.GetInput("role"),
		Duration: d,
	}
	return &c, nil
}

func run(ctx context.Context, cfg *Config) error {
	fmt.Printf("role: %s\n", cfg.Role)
	fmt.Printf("duration: %d\n", cfg.Duration)
	return nil
}
