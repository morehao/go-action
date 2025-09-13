package infra

import (
	"fmt"
	"net/http"
)

func WriteSSE(w http.ResponseWriter, id, event string, data []byte) error {
	// 确保支持 flush
	flusher, ok := w.(http.Flusher)
	if !ok {
		return fmt.Errorf("ResponseWriter does not support flushing")
	}

	if len(id) > 0 {
		if _, err := fmt.Fprintf(w, "id: %s\n", id); err != nil {
			return err
		}
	}
	if len(event) > 0 {
		if _, err := fmt.Fprintf(w, "event: %s\n", event); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
		return err
	}

	flusher.Flush()
	return nil
}
