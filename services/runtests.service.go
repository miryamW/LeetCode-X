package runtests

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func runTests(testID string) error {
	// יצירת לקוח Docker
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("Error creating Docker client: %v", err)
	}

	// הקשר הקונטיינר
	ctx := context.Background()

	// יצירת קונטיינר חדש
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "my-python-test-image", // התמונה שברצונך להריץ
		Env:   []string{fmt.Sprintf("TEST_ID=%s", testID)}, // העברת משתנים לקונטיינר
	}, nil, nil, nil, "")
	if err != nil {
		return fmt.Errorf("Error creating Docker container: %v", err)
	}

	// הפעלת הקונטיינר
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("Error starting Docker container: %v", err)
	}

	// המתנה לסיום הקונטיינר
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNextExit)
	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("Error waiting for container: %v", err)
		}
	case <-statusCh:
	}

	// קריאת הלוגים מהקונטיינר
	out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return fmt.Errorf("Error getting container logs: %v", err)
	}

	// הצגת הלוגים
	defer out.Close()
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return nil
}