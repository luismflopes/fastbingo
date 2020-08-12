package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func RequestResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		log.Println(c.Request.URL.Path, " Request: ", readBody(rdr1)) // Print request body to output log file
		c.Request.Body = rdr2

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		log.Println(c.Request.URL.Path, " Response: ", blw.body.String()) // Print Response body to output log file

	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()

	replacer := strings.NewReplacer("\n", "", "\t", "", " ", "")
	s = replacer.Replace(s)

	re := regexp.MustCompile("(\"password\"|\"new_password\"|\"old_password\"|\"access_token\"|\"jsc\"):\"[^\"]*\"")

	return re.ReplaceAllString(s, "$1:\"XXXXXXX\"")
}
