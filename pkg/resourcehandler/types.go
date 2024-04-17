package resourcehandler

import (
	"context"
	"io"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

// Type indicates the type of ResourceHandler,
// e.g. Kubernetes, AWS, etc.
type Type = string

// ResourceHandler holds the actions that an resourcehandlers must satisfy.
type ResourceHandler interface {
	// Type returns Type.
	Type() Type

	// IsConnected validates whether is connected.
	IsConnected(context.Context) error

	// Burst returns the maximum number of operations that can be called at once.
	Burst() int

	// ID returns an operation identifier of the resourcehandlers.
	//
	// The result is not a unique notation of in the traditional sense,
	// which means we can't use it to identify the only resourcehandlers,
	// but rather a generalization of the characteristics of the resourcehandlers.
	//
	// If two operators have the same ID,
	// we can group them together for some operations.
	//
	// ID returns a blank string if no that kind of identifier.
	ID() string

	// GetKeys returns keys from the given resource.
	//
	// The given walrus.ResourceComponent item must specify the following fields:
	// ID, DeployerType, Type and Name.
	GetKeys(context.Context, *walrus.ResourceComponents) (*ResourceComponentOperationKeys, error)

	// GetStatus gets status of the given resource.
	//
	// The given walrus.ResourceComponent item must specify the following fields:
	// ID, DeployerType, Type and Name.
	GetStatus(context.Context, *walrus.ResourceComponents) ([]meta.Condition, error)

	// GetComponents gets components of the given resource,
	// returns list must not be `nil` unless unexpected input or raising error,
	// it can be used to clean stale items safety if got an empty list.
	GetComponents(context.Context, *walrus.ResourceComponents) ([]*walrus.ResourceComponents, error)

	// Log gets logs from the given key.
	Log(context.Context, string, LogOptions) error

	// Exec executes commands to the given key.
	Exec(context.Context, string, ExecOptions) error
}

// LogOptions holds the options of ResourceHandler's Log action.
type LogOptions struct {
	// Out receives the output.
	Out io.Writer
	// WithoutFollow returns logs without following.
	WithoutFollow bool
	// Previous indicates to get the previous log of the accessing target.
	Previous bool
	// SinceSeconds returns logs newer than a relative duration.
	SinceSeconds *int64
	// Timestamps returns logs with RFC3339 or RFC3339Nano timestamp.
	Timestamps bool
	// TailLines indicates to get the lines from end of the logs.
	TailLines *int64
}

// ExecOptions holds the options of ResourceHandler's Exec action.
type ExecOptions struct {
	// Out receives the output.
	Out io.Writer
	// In passes the input.
	In io.Reader
	// Shell indicates to launch what kind of shell.
	Shell string
	// Resizer indicates to resize the size(width, height) of the terminal.
	Resizer TerminalResizer
}

// TerminalResizer holds the options to resize the terminal.
type TerminalResizer interface {
	// Next returns the new terminal size after the terminal has been resized.
	// It returns false when monitoring has been stopped.
	Next() (width, height uint16, ok bool)
}

type CreateOptions struct {
	Connector walruscore.Connector
}

type Creator func(context.Context, CreateOptions) (ResourceHandler, error)

type ResourceComponentOperationKeys struct {
	// TODO: Need to consider how the resource component is used
}
