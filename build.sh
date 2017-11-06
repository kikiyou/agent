# export $PATH=$PATH:/home/yh/go/bin
~/go/bin/go-bindata -pkg templates -o templates/bindata.go templates/*.tmpl

~/go/bin/go-bindata-assetfs static/...

~/go/bin/go-bindata -pkg shell -o shell/bindata.go shell/linux_json_api.sh
