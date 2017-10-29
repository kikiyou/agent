package g

import (
	"crypto/md5"
	"encoding/hex"
)

// func GetCurrPluginVersion() string {
// 	if !Config().Plugin.Enabled {
// 		return "plugin not enabled"
// 	}

// 	pluginDir := Config().Plugin.Dir
// 	if !file.IsExist(pluginDir) {
// 		return "plugin dir not existent"
// 	}

// 	cmd := exec.Command("git", "rev-parse", "HEAD")
// 	cmd.Dir = pluginDir

// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		return fmt.Sprintf("Error:%s", err.Error())
// 	}

// 	return strings.TrimSpace(out.String())
// }

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
