package templates

import "github.com/dave/jennifer/jen"

/*
Mostly generated by github.com/aloder/tojen, slightly modified by hand.
All updates to template code must be done like this.
*/

func genDeclAt122(v1, v2, v3 string) jen.Code {
	return jen.Null().Const().Id("path").Op("=").Lit(v1).Line().Const().Id("content_type").Op("=").Lit(v2).Line().Const().Id("cache_control").Op("=").Lit(v3)
}
func genFuncinit() jen.Code {
	return jen.Func().Id("init").Params().Block(jen.Qual("github.com/GoogleCloudPlatform/functions-framework-go/functions", "HTTP").Call(jen.Id("path"), jen.Id("call")))
}
func genFunccall() jen.Code {
	return jen.Func().Id("call").Params(jen.Id("w").Qual("net/http", "ResponseWriter"), jen.Id("r").Op("*").Qual("net/http", "Request")).Block(jen.List(jen.Id("file"), jen.Id("_")).Op(":=").Qual("os", "Open").Call(jen.Lit("./serverless_function_source_code/").Op("+").Id("path")), jen.Defer().Id("file").Dot("Close").Call(), jen.Id("w").Dot("Header").Call().Dot("Set").Call(jen.Lit("Content-Type"), jen.Id("content_type")), jen.Id("w").Dot("Header").Call().Dot("Set").Call(jen.Lit("Cache-Control"), jen.Id("cache_control")), jen.List(jen.Id("_"), jen.Id("_")).Op("=").Qual("io", "Copy").Call(jen.Id("w"), jen.Id("file")))
}
func GenSingleFile(path string, contentType string, ma string) *jen.File {
	ret := jen.NewFile("autogcf_single_file")
	ret.ImportName("github.com/GoogleCloudPlatform/functions-framework-go/functions", "functions")
	ret.Add(genDeclAt122(path, contentType, ma))
	ret.Add(genFuncinit())
	ret.Add(genFunccall())
	return ret
}