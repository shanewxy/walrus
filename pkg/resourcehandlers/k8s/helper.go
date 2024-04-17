package k8s

import (
	"context"
	"fmt"

	"github.com/seal-io/utils/json"
	"k8s.io/client-go/rest"
)

func IsConnected(ctx context.Context, r rest.Interface) error {
	body, err := r.Get().
		AbsPath("/version").
		Do(ctx).
		Raw()
	if err != nil {
		return err
	}

	var info struct {
		Major    string `json:"major"`
		Minor    string `json:"minor"`
		Compiler string `json:"compiler"`
		Platform string `json:"platform"`
	}

	err = json.Unmarshal(body, &info)
	if err != nil {
		return fmt.Errorf("unable to parse the server version: %w", err)
	}

	return nil
}
