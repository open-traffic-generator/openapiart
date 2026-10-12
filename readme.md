# openapiart - OpenAPI Artifact Generator

[![CICD](https://github.com/open-traffic-generator/openapiart/workflows/CICD/badge.svg)](https://github.com/open-traffic-generator/openapiart/actions)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![python](https://img.shields.io/pypi/pyversions/openapiart.svg)](https://pypi.python.org/pypi/openapiart)
[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://en.wikipedia.org/wiki/MIT_License)

The `OpenAPIArt` (OpenAPI Artifact Generator) python package does the following:
- pre-processes OpenAPI yaml files according to the [MODELGUIDE](../main/MODELGUIDE.md)
- using the path keyword bundles all dependency OpenAPI yaml files into a single openapi.yaml file
- post-processes any [MODELGUIDE](../main/MODELGUIDE.md) extensions
- validates the bundled openapi.yaml file
- generates a `.proto` file from the openapi file
- optionally generates a static redocly documentation file 
- optionally generates a `python ux sdk` from the openapi file
- optionally generates a `go ux sdk` from the openapi file

## Getting started
Install the package
```
pip install openapiart
```

Generate artifacts from OpenAPI files
```python
""" 
The following command will produce these artifacts:
    - ./artifacts/openapi.yaml
    - ./artifacts/openapi.json
    - ./artifacts/openapi.html
    - ./artifacts/sample.proto
    - ./artifacts/sample/__init__.py
    - ./artifacts/sample/sample.py
    - ./artifacts/sample/sample_pb2.py
    - ./artifacts/sample/sample_pb2_grpc.py
    - ./pkg/openapiart.go
    - ./pkg/go.mod
    - ./pkg/go.sum
    - ./pkg/sanity/sanity_grpc.pb.go
    - ./pkg/sanity/sanity.pb.go
"""
import openapiart

# bundle api files
# validate the bundled file
# generate the documentation file
art = openapiart.OpenApiArt(
    api_files=[
        "./openapiart/tests/api/info.yaml",
        "./openapiart/tests/common/common.yaml",
        "./openapiart/tests/api/api.yaml",
    ],
    artifact_dir="./artifacts",
    protobuf_name="sanity",
    extension_prefix="sanity",
)

# optionally generate a python ux sdk and python protobuf/grpc stubs
art.GeneratePythonSdk(
    package_name="sanity"
)

# optionally generate a go ux sdk and go protobuf/grpc stubs
art.GenerateGoSdk(
    package_dir="github.com/open-traffic-generator/openapiart/pkg", 
    package_name="openapiart"
)
```

## Generated SDK RPC logging

Generated Python and Go SDKs can log unary HTTP and gRPC request/response
exchanges by setting a package-specific environment variable named
`<PACKAGE>_RPC_LOG`, where `<PACKAGE>` is the generated SDK package name
uppercased with non-alphanumeric characters converted to `_`. For example, a
generated `snappi` SDK uses `SNAPPI_RPC_LOG`. `OPENAPIART_RPC_LOG` is also
honored as a generic fallback.

- `stdout` or `-`: write RPC logs to stdout
- `stderr`: write RPC logs to stderr
- `<path>`: append RPC logs to the file at `<path>`

Each exchange is written as a YAML document separated by `---`.
Failures are represented under `response.error`; HTTP entries retain the
response status and any body bytes read before a body-read failure.

## Specifications
> This repository is based on the [OpenAPI specification](
https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md) 
which is a standard, language-agnostic interface to RESTful APIs. 

> [Modeling guide specific to this package](../main/MODELGUIDE.md)
