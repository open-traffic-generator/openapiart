# OpenAPIArt 

[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![Build](https://github.com/open-traffic-generator/openapiart/workflows/build/badge.svg)](https://github.com/open-traffic-generator/openapiart/actions)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/open-traffic-generator/openapiart.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/openapiart/alerts/)
[![Language grade: Python](https://img.shields.io/lgtm/grade/python/g/open-traffic-generator/openapiart.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/openapiart/context:python)
[![pypi](https://img.shields.io/pypi/v/openapiart.svg)](https://pypi.org/project/openapiart)
[![python](https://img.shields.io/pypi/pyversions/snappi.svg)](https://pypi.python.org/pypi/snappi)


The `OpenAPIArt` (OpenAPI Artifact Generator) python package does the following:
- bundles individual yaml files into a single file
- post process x- extensions
- validates the bundled openapi.yaml file
- creates an enhanced ux python file containing all the classes generated from 
  the bundled openapi.yaml file

## Getting started
Install the package
```
pip install openapiart
```

Generate artifacts from OpenAPI models
```python
import openapiart

# the following command produces these artifacts
# openapi.yaml
# openapi.json
# openapi.html
# test.py
openapiart.OpenApiArt(api_files=['./tests/api/api.yaml'], python_module_name='sample')
```

## Specifications
> This repository is based on the [OpenAPI specification](
https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md) which is a standard, language-agnostic interface to RESTful APIs. 

> Modeling guide specific to this package


