package template

var (
	MakefileSRV = `
.PHONY: {{dehyphen .Alias}}
demo:
	protoc \
     -I . \
     --micro_out=. \
     --go_out=. \
     ./{{dehyphen .Alias}}/{{dehyphen .Alias}}.proto
`
)



