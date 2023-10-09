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
		Name: "opts",
	}
	fIterations = &cli.IntFlag{
		Name:    "iterations",
		Value:   10,
		Aliases: []string{"i"},
	}
	fConcats = &cli.IntFlag{
		Name:    "concats",
		Value:   3,
		Aliases: []string{"c"},
	}
	fHypenChance = &cli.IntFlag{
		Name:  "hypen",
		Value: 100,
		Usage: "% chance of hypen creation",
	}
)

func main() {
	app := &cli.App{
		Name:   "boom",
		Usage:  "make an explosive entrance",
		Action: boom,
		Flags: []cli.Flag{
			fOpts,
			fIterations, fConcats, fHypenChance,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func boom(c *cli.Context) error {
	values := fOpts.Get(c)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rs := func() string { return values[r.Intn(len(values))] }

	var (
		v string
		b strings.Builder
	)

	for i := 0; i < fIterations.Get(c); i++ {
		for i := 0; i < fConcats.Get(c); i++ {
			b.WriteString(rs())
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
