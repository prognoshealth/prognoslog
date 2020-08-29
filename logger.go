package prognoslog

import (
	"encoding/json"
	"fmt"
	"io"
)

// Log provides a way to write a semi structured logline based upon the desired
// type. Currently supports JSON, KVP and TXT.
//
// TODO: right now we are panicing on errors as this is primarily used for
// writing to stdout and one of the main purposes of this library is to reduce
// boilerplate. Returning errors limits that goal.
//
// One potential options is to push errors into the Log struct when they occur.
type Logger struct {
	out       io.Writer
	isVerbose bool
}

// enforce panics on error.
func enforce(n int, err error) {
	if err != nil {
		panic(err)
	}
}

// SetVerbosity sets a Log's verbosity (to true).
func (log *Logger) SetVerbosity(v ...bool) {
	if len(v) == 0 {
		log.isVerbose = true
	} else {
		log.isVerbose = v[0]
	}
}

// JSON logs to stdout an object marshaled as a JSON string referenced with the
// provided name.
func (log *Logger) JSON(name string, obj interface{}) {
	b, err := json.Marshal(obj)

	enforce(0, err)
	enforce(fmt.Fprintf(log.out, "JSON %v=%v\n", name, string(b)))
}

// JSONIfVerbose calls JSON, iff log.isVerbose.
func (log *Logger) JSONIfVerbose(name string, obj interface{}) {
	if log.isVerbose {
		log.JSON(name, obj)
	}
}

// JSONString logs to stdout a string that contains a JSON formatted string
func (log *Logger) JSONString(name string, s string) {
	log.JSON(name, json.RawMessage(s))
}

// JSONStringIfVerbose calls JSONString, iff log.isVerbose.
func (log *Logger) JSONStringIfVerbose(name string, s string) {
	if log.isVerbose {
		log.JSONString(name, s)
	}
}

// KVP logs the name,obj key value pair to stdout.
func (log *Logger) KVP(name string, obj interface{}) {
	enforce(fmt.Fprintf(log.out, "KVP %v=%#v\n", name, obj))
}

// KVPIfVerbose calls KVP, iff log.isVerbose.
func (log *Logger) KVPIfVerbose(name string, obj interface{}) {
	if log.isVerbose {
		log.KVP(name, obj)
	}
}

// Txt logs the provided string and args to stdout.
func (log *Logger) Txt(format string, a ...interface{}) {

	if len(a) == 0 {
		enforce(fmt.Fprintln(log.out, "TXT "+format))
	} else {
		enforce(fmt.Fprintf(log.out, "TXT "+format+"\n", a...))
	}
}

// TxtIfVerbose calls Txt, iff log.isVerbose.
func (log *Logger) TxtIfVerbose(format string, a ...interface{}) {
	if log.isVerbose {
		if len(a) != 0 {
			log.Txt(format, a...)
		} else {
			log.Txt(format)
		}
	}
}
