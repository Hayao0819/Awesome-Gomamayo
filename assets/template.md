# Awesome Gomamayo

偉業なゴママヨ実装を纏めています。PR大歓迎です。

## ゴママヨとは

複合語の中に連続する音や文字が含まれている状態、またはその語のことをゴママヨと呼びます。

ここではゴママヨの語を検出するためのライブラリ/ツールをまとめています。

ゴママヨには[様々な種類](https://3qua9la-notebook.hatenablog.com/entry/2021/04/10/220317)が存在しています。

## 実装

現在 {{ len .Codes }} 個の言語による {{ .RepoCount }} 個の実装が掲載されています。

{{- range $lang, $codes := .Codes }}
### {{ $lang }}
{{- range $codes }}
- [{{ .Label }}]({{ .Url }})  
  {{ .Desc }}

{{- end }}
{{- end }}

## 参考文献

{{- range .Websites }}
- [{{ .Title }}]({{ .Url }})
{{- end }}
