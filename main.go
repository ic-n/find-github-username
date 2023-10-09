package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var (
	fOpts = &cli.StringSliceFlag{
		Name:        "opts",
		Required:    true,
		Usage:       "Comma-separated list of possible parts in username",
		DefaultText: "nice,cool,frog",
	}
	fIterations = &cli.IntFlag{
		Name:        "iterations",
		Value:       10,
		Aliases:     []string{"i"},
		Usage:       "Number of generation to test",
		DefaultText: "10",
	}
	fConcats = &cli.IntFlag{
		Name:        "concats",
		Value:       3,
		Aliases:     []string{"c"},
		Usage:       "Number of concat operations to username (how many times pick a chunk)",
		DefaultText: "3",
	}
	fHypenChance = &cli.IntFlag{
		Name:        "hypen",
		Value:       100,
		Usage:       "% chance of hypen inserstion",
		DefaultText: "100",
	}
)

func main() {
	app := &cli.App{
		Name:   "find",
		Usage:  "find vacant github login",
		Action: find,
		Flags:  []cli.Flag{fOpts, fIterations, fConcats, fHypenChance},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func find(c *cli.Context) error {
	opts := fOpts.Get(c)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var (
		v string
		b strings.Builder
	)

	for i := 0; i < fIterations.Get(c); i++ {
		for i := 0; i < fConcats.Get(c); i++ {
			b.WriteString(opts[r.Intn(len(opts))])
			if r.Intn(100) < fHypenChance.Get(c) {
				b.WriteRune('-')
			}
		}
		v = b.String()
		v = strings.Trim(v, "-")
		b.Reset()

		if len(v) < 3 {
			continue
		}

		resp, err := http.Get(fmt.Sprintf("https://github.com/%s", v))
		if err != nil {
			return errors.Wrap(err, "failed to request github")
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusOK:
			fmt.Printf("(-) %s is busy\n", v)
		case http.StatusNotFound:
			fmt.Printf("(V) %s is vacant\n", v)
		default:
			return fmt.Errorf("unexpected status code %d %s", resp.StatusCode, resp.Status)
		}
	}

	return nil
}
