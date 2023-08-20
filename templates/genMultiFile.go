package templates

import "github.com/dave/jennifer/jen"

func genDeclAt122x(v1, v2, v3, v4, v5 string) jen.Code {
	return jen.Null().Var().Id("dir").Op("=").Lit(v1).Line().Var().Id("paths").Op("=").Index().Id("string").Values(jen.Op(v2)).Line().Var().Id("local_paths").Op("=").Index().Id("string").Values(jen.Op(v3)).Line().Var().Id("content_types").Op("=").Index().Id("string").Values(jen.Op(v4)).Line().Var().Id("cache_controls").Op("=").Index().Id("string").Values(jen.Op(v5))
}
func genFuncinitx() jen.Code {
	return jen.Func().Id("init").Params().Block(jen.Qual("github.com/GoogleCloudPlatform/functions-framework-go/functions", "HTTP").Call(jen.Id("dir"), jen.Id("call")))
}
func genFunccallx() jen.Code {
	return jen.Func().Id("call").Params(jen.Id("w").Qual("net/http", "ResponseWriter"), jen.Id("r").Op("*").Qual("net/http", "Request")).Block(jen.Id("path").Op(":=").Id("r").Dot("URL").Dot("Path").Index(jen.Lit(1), jen.Empty()), jen.For(jen.List(jen.Id("i")).Op(":=").Range().Id("paths")).Block(jen.Id("p").Op(":=").Id("paths").Index(jen.Id("i")), jen.If(jen.Id("path").Op("==").Id("p")).Block(jen.List(jen.Id("file"), jen.Id("_")).Op(":=").Qual("os", "Open").Call(jen.Lit("./serverless_function_source_code/").Op("+").Id("local_paths").Index(jen.Id("i"))), jen.Defer().Id("file").Dot("Close").Call(), jen.Id("w").Dot("Header").Call().Dot("Set").Call(jen.Lit("Content-Type"), jen.Id("content_types").Index(jen.Id("i"))), jen.Id("w").Dot("Header").Call().Dot("Set").Call(jen.Lit("Cache-Control"), jen.Id("cache_controls").Index(jen.Id("i"))), jen.List(jen.Id("_"), jen.Id("_")).Op("=").Qual("io", "Copy").Call(jen.Id("w"), jen.Id("file")).Line().Return())), jen.Qual("net/http", "Error").Call(jen.Id("w"), jen.Lit("Not Found"), jen.Qual("net/http", "StatusNotFound")))
}
func GenMultiFile(dir, paths, localpaths, contenttypes, ma string) *jen.File {
	ret := jen.NewFile("autogcf_multi_file")
	ret.ImportName("github.com/GoogleCloudPlatform/functions-framework-go/functions", "functions")
	ret.Add(genDeclAt122x(dir, paths, localpaths, contenttypes, ma))
	ret.Add(genFuncinitx())
	ret.Add(genFunccallx())
	return ret
}
