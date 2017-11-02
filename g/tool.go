package g

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	// "html/template"
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

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)
	// fmt.Println("uuuuuuuuuuuuuuu")
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)

	}
}

// template is nil, it is created from the first file.
// func parseFiles(t *Template, filenames ...string) (*Template, error) {
// 	if err := t.checkCanParse(); err != nil {
// 		return nil, err
// 	}

// 	if len(filenames) == 0 {
// 		// Not really a problem, but be consistent.
// 		return nil, fmt.Errorf("html/template: no files named in call to ParseFiles")
// 	}
// 	for _, filename := range filenames {
// 		b, err := ioutil.ReadFile(filename)
// 		if err != nil {
// 			return nil, err
// 		}
// 		s := string(b)
// 		name := filepath.Base(filename)
// 		// First template becomes return value if not already defined,
// 		// and we use that one for subsequent New calls to associate
// 		// all the templates together. Also, if this file has the same name
// 		// as t, this file becomes the contents of t, so
// 		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
// 		// works. Otherwise we create a new template associated with t.
// 		var tmpl *Template
// 		if t == nil {
// 			t = New(name)
// 		}
// 		if name == t.Name() {
// 			tmpl = t
// 		} else {
// 			tmpl = t.New(name)
// 		}
// 		_, err = tmpl.Parse(s)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return t, nil
// }
