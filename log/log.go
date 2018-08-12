/*
   Copyright 2018 Assetnote

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package log

import (
	"github.com/sirupsen/logrus"
)

var (
	Error   = logrus.Error
	Errorf  = logrus.Errorf
	Errorln = logrus.Errorln
	Fatal   = logrus.Fatal
	Fatalf  = logrus.Fatalf
	Fatalln = logrus.Fatalln
	Panic   = logrus.Panic
	Panicf  = logrus.Panicf
	Panicln = logrus.Panicln
	Print   = logrus.Print
	Printf  = logrus.Printf
	Println = logrus.Println
	Debug   = logrus.Debug
	Debugf  = logrus.Debugf
	Debugln = logrus.Debugln
	Warn    = logrus.Warn
	Warnf   = logrus.Warnf
	Warnln  = logrus.Warnln
	Info    = logrus.Info
	Infof   = logrus.Infof
	Infoln  = logrus.Infoln
)

type Fields logrus.Fields

func WithFields(f Fields) *logrus.Entry {
	return logrus.WithFields(logrus.Fields(f))
}
