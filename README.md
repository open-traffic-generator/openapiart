# OpenAPIArt 

[![CICD](https://github.com/open-traffic-generator/openapiart/workflows/CICD/badge.svg)](https://github.com/open-traffic-generator/openapiart/actions)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![pypi](https://img.shields.io/pypi/v/openapiart.svg)](https://pypi.org/project/openapiart)
[![python](https://img.shields.io/pypi/pyversions/openapiart.svg)](https://pypi.python.org/pypi/openapiart)
[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/open-traffic-generator/openapiart.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/openapiart/alerts/)
[![Language grade: Python](https://img.shields.io/lgtm/grade/python/g/open-traffic-generator/openapiart.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/openapiart/context:python)

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

## Specifications
> This repository is based on the [OpenAPI specification](
https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md) 
which is a standard, language-agnostic interface to RESTful APIs. 

> [Modeling guide specific to this package](../main/MODELGUIDE.md)


