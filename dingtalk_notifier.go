package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DingTalkNotifier is a notifier that sends notifications to DingTalk.
type DingTalkNotifier struct {
	AccessToken string `yaml:"access_token"`
}

// Notify sends a notification to DingTalk.
func (n *DingTalkNotifier) Notify(event string, severity string, msg string) error {
	dingTalkURL := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", n.AccessToken)
	message := fmt.Sprintf("[Nuclei] %s %s\n%s", event, severity, msg)
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": message,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.Post(dingTalkURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("DingTalk notification failed with status code %d", resp.StatusCode)
	}
	return nil
}
