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

Using the validated openapi.yaml file it then:
- generates a static redocly documentation file 
- generates a `protobuf` file
- generates protobuf based python files
- generates an enhanced ux python module

## Getting started
Install the package
```
pip install openapiart
```

Generate artifacts from OpenAPI files
```python
import openapiart

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
"""
openapiart.OpenApiArt(
    api_files=[
        './tests/api/api.yaml'
        './tests/api/info.yaml'
        './tests/common/common.yaml'
        ], 
    python_module_name='sample', 
    protobuf_file_name='sample',
    protobuf_package_name='sample',
    output_dir='./artifacts',
    extension_prefix='sample'
)
```

## Specifications
> This repository is based on the [OpenAPI specification](
https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md) 
which is a standard, language-agnostic interface to RESTful APIs. 

> [Modeling guide specific to this package](../main/MODELGUIDE.md)


