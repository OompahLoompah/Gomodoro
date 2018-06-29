// Copyright © 2018 Sean Heuer <seanmheuer@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0xAX/notificator"
	"github.com/spf13/cobra"
	t "github.com/OompahLoompah/Gomodoro/pkg/timer"
)

var seconds int
var breakSeconds int
var notify *notificator.Notificator
var pomo bool
var tag string

var rootCmd = &cobra.Command{
	Use:   "gomodoro",
	Short: "A pomodoro-inspired timer application",
	Long: `A pomodoro-inspired timer application that supports recording sessions
	to metrics servers via JSON.`,
        Run: func(cmd *cobra.Command, args []string) {
		if pomo {
			for ;; {
				err := t.Timer(1500, notifier, true, tag)
				if err != nil {
					log.Fatal(err)
				}
				err = t.Timer(300, nil, false, tag)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			if seconds > 0 {
				err := t.Timer(seconds, notifier, true, tag)
				if err != nil {
					log.Fatal(err)
				}
			}
			if breakSeconds > 0 {
				err := t.Timer(breakSeconds, nil, false, tag)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&seconds, "time", "T", 0, "time to count down from")
	rootCmd.Flags().IntVarP(&breakSeconds, "break", "b", 0, "Break time to count down from")
	rootCmd.Flags().BoolVarP(&pomo, "pomodoro", "p", false, "Run continuous pomodoro")
	rootCmd.Flags().StringVarP(&tag, "comment", "c", "", "Session tag")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pomodoro.yaml)")
}

func notifier() {
	notify = notificator.New(notificator.Options{
		AppName:          "Pomodoro Timer",
	})
	notify.Push("", "Time's up!", "", notificator.UR_CRITICAL)
}
