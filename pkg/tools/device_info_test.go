package tools

import (
	"fmt"
	"testing"
)

func TestSpeedFun(t *testing.T) {
	startTestSpeed()
	fmt.Printf("%f", DeviceInfo.DownloadSpeed)
	fmt.Printf("%f", DeviceInfo.UploadSpeed)
}
