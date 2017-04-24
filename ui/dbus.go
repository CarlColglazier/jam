// Copyright (c) 2016, 2017 Evgeny Badin

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ui

import (
	"fmt"
	"github.com/godbus/dbus"
	"os"
)

func (app App) Next() *dbus.Error {
	app.Status.State <- next
	return nil
}

func (app App) Previous() *dbus.Error {
	app.Status.State <- prev
	return nil
}

func (app App) Pause() *dbus.Error {
	app.Status.State <- pause
	return nil
}

func (app App) PlayPause() *dbus.Error {
	app.Status.State <- pause
	return nil
}

func (app App) Stop() *dbus.Error {
	app.Status.State <- stop
	return nil
}

func (app App) Play() *dbus.Error {
	app.Status.State <- play
	return nil
}

func conn() *dbus.Conn {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	reply, err := conn.RequestName("org.mpris.MediaPlayer2.jam", dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Fprintln(os.Stderr, "name already taken")
		os.Exit(1)
	}
	return conn
}

func (app *App) dbusSetup() {
	conn := conn()
	conn.Export(*app, "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2.Player")
}
