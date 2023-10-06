package main

import (
	"context"
	"fmt"

	"github.com/kr/pretty"
	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	Key    string
	Api    string
	Server string
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
	c := Config{
		Key:    action.GetInput("key"),
		Api:    action.GetInput("api"),
		Server: action.GetInput("server"),
	}
	return &c, nil
}

func run(ctx context.Context, cfg *Config) error {
	if _, err := pretty.Println(cfg); err != nil {
		return err
	}
	return nil
}
