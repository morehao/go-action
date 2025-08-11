package main

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/loader/file"
	"github.com/cloudwego/eino-ext/components/document/parser/xlsx"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	callbacksHelper "github.com/cloudwego/eino/utils/callbacks"
	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"
)

func XlsxHandler(ctx *gin.Context) {
	glog.Info(ctx, ("===== call XLSX Parser directly ====="))
	// Initialize the parser
	parser, err := xlsx.NewXlsxParser(ctx, nil)
	if err != nil {
		glog.Errorf(ctx, "xlsx.NewXLSXParser failed, err=%v", err)
		gincontext.Fail(ctx, err)
	}

	// Initialize the loader
	loader, err := file.NewFileLoader(ctx, &file.FileLoaderConfig{
		Parser: parser,
	})
	if err != nil {
		glog.Errorf(ctx, "file.NewFileLoader failed, err=%v", err)
		gincontext.Fail(ctx, err)
	}

	// Load the document
	filePath := "./testdata/test.xlsx"
	docs, err := loader.Load(ctx, document.Source{
		URI: filePath,
	})
	if err != nil {
		glog.Errorf(ctx, "loader.Load failed, err=%v", err)
		gincontext.Fail(ctx, err)
	}

	glog.Info(ctx, "===== Documents Content =====")
	for _, doc := range docs {
		glog.Infof(ctx, "Id %v content: %v metadata: %v", doc.ID, doc.Content, doc.MetaData)
	}

	glog.Info(ctx, "===== call XLSX Parser in Chain =====")
	// Create callback handler
	handlerHelper := &callbacksHelper.LoaderCallbackHandler{
		OnStart: func(ctx context.Context, info *callbacks.RunInfo, input *document.LoaderCallbackInput) context.Context {
			glog.Infof(ctx, "start loading docs...: %s\n", input.Source.URI)
			return ctx
		},
		OnEnd: func(ctx context.Context, info *callbacks.RunInfo, output *document.LoaderCallbackOutput) context.Context {
			glog.Infof(ctx, "complete loading docs，total loaded docs: %d\n", len(output.Docs))
			return ctx
		},
		// OnError
	}

	// Use callback handler
	handler := callbacksHelper.NewHandlerHelper().
		Loader(handlerHelper).
		Handler()

	chain := compose.NewChain[document.Source, []*schema.Document]()
	chain.AppendLoader(loader)
	// Use at runtime
	run, err := chain.Compile(ctx)
	if err != nil {
		glog.Errorf(ctx, "chain.Compile failed, err=%v", err)
	}

	outDocs, err := run.Invoke(ctx, document.Source{
		URI: filePath,
	}, compose.WithCallbacks(handler))
	if err != nil {
		glog.Errorf(ctx, "run.Invoke failed, err=%v", err)
		gincontext.Fail(ctx, err)
	}

	glog.Info(ctx, "===== Documents Content =====")
	for _, doc := range outDocs {
		glog.Infof(ctx, "Id %v content: %v metadata: %v", doc.ID, doc.Content, doc.MetaData)
	}
}
