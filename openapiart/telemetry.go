import (
	"fmt"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type telemetry struct {
	transport     string
	endpoint      string
	rootCtx       context.Context
	traceProvider *sdktrace.TracerProvider
}

type Telemetry interface {
	isOTLPEnabled() bool
	getRootContext() context.Context
	WithExporterTrasnport(transport string) Telemetry
	WithExporterEndPoint(endpoint string) Telemetry
	WithRootContext(ctx context.Context) Telemetry
	StartTracing() (Telemetry, error)
	StopTracing()
	NewSpan(ctx context.Context, name string) (context.Context, trace.Span)
	SetSpanStatus(span trace.Span, code codes.Code, description string)
	SetSpanAttributes(span trace.Span, attrs []attribute.KeyValue)
	CloseSpan(span trace.Span)
	View()
}

var tracer = otel.Tracer("gosnappi-tracer")

// just a debug function to view the telemetry struct
// TODO: Remove this.
func (t *telemetry) View() {
	fmt.Println("fetching the tracer")
	fmt.Printf("tel is %v\n", t)
	fmt.Println(t.transport)
}

// Internal fucntion to check wheather telemetry is enabled or not.
// Used by rest of the functions to become no-ops.
func (t *telemetry) isOTLPEnabled() bool {
	return t.endpoint != ""
}

// Internal function to fetch the root context if provided by user.
// Default context is background
func (t *telemetry) getRootContext() context.Context {
	if t.rootCtx != nil {
		return t.rootCtx
	}

	return context.Background()
}

// Sets the transport which is to be used to communicate with the OTLP collector.
// Available options are HTTP/GRPC
// Default value is HTTP
func (t *telemetry) WithExporterTrasnport(transport string) Telemetry {
	t.transport = transport
	return t
}

// Sets The endpoint of the OTLP collector which will collect the data and visualize it.
// For HTTP the endpoint shoule be like "IP:PORT" e.g. "127.0.0.1:4138"
func (t *telemetry) WithExporterEndPoint(endpoint string) Telemetry {
	t.endpoint = endpoint
	return t
}

// Sets The Root Context.
// If the user wants all the spans to be the child of a single span, this method should be used.
func (t *telemetry) WithRootContext(ctx context.Context) Telemetry {
	t.rootCtx = ctx
	return t
}

// Initiates the trace provider with proper resources, exporter information
// and span processors
func (t *telemetry) StartTracing() (Telemetry, error) {

	if t.isOTLPEnabled() {

		// creating exporter for now only concentrating on http
		exporter, err := otlptrace.New(
			context.Background(),
			otlptracehttp.NewClient(
				otlptracehttp.WithInsecure(),
				otlptracehttp.WithEndpoint(t.endpoint),
			),
		)

		// raising error if exporter creation had some issues
		if err != nil {
			return nil, fmt.Errorf("Error creating OTLP trace exporter: %v", err)
		}

		// defining the service name
		resources, err := resource.New(
			context.Background(),
			resource.WithAttributes(
				attribute.String("service.name", "go-snappi"),
				attribute.String("application", "go-snappi"),
			),
		)

		if err != nil {
			return nil, fmt.Errorf("Error setting resources for OTLP trace: %v", err)
		}

		// Selecting batch span processor as of now
		spanProcessor := sdktrace.NewBatchSpanProcessor(exporter)

		// Creating the traceProvider
		traceProvider := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithSpanProcessor(spanProcessor),
			sdktrace.WithResource(resources),
		)

		// a TextMapPropagator is a component that injects values ​​into and extracts values ​​
		// from transport as string key/value pairs
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{}),
		)

		otel.SetTracerProvider(traceProvider)
		t.traceProvider = traceProvider

		return t, nil
	}

	return nil, nil
}

// Gracefully shuts down the trace provider and flushes the collector streams.
func (t *telemetry) StopTracing() {
	if t.isOTLPEnabled() {
		if err := t.traceProvider.Shutdown(context.Background()); err != nil {
			fmt.Println("Failed shutting down trace provider")
		}
		fmt.Println("shut down successful !!!!")
	}
}

// Creates a new span , if a parent context is passed then it will be child span.
// By default the span will be root span.
func (t *telemetry) NewSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	if t.isOTLPEnabled() {
		return tracer.Start(ctx, name)
	}

	return nil, nil
}

// Closes a Span.
func (t *telemetry) CloseSpan(span trace.Span) {
	if t.isOTLPEnabled() {
		span.End()
	}
}

// This part is generally used to set errors and stuff with proper code and desc.
func (t *telemetry) SetSpanStatus(span trace.Span, code codes.Code, description string) {
	if t.isOTLPEnabled() {
		span.SetStatus(code, description)
	}
}

// This method is used to set attributes to a particular span.
func (t *telemetry) SetSpanAttributes(span trace.Span, attrs []attribute.KeyValue) {
	if t.isOTLPEnabled() {
		span.SetAttributes(attrs...)
	}
}
