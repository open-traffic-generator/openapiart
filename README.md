# OpenAPIArt 

[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![CICD](https://github.com/open-traffic-generator/openapiart/workflows/CICD/badge.svg)](https://github.com/open-traffic-generator/openapiart/actions)
[![pypi](https://img.shields.io/pypi/v/openapiart.svg)](https://pypi.org/project/openapiart)
[![python](https://img.shields.io/pypi/pyversions/snappi.svg)](https://pypi.python.org/pypi/snappi)


The `OpenAPIArt` (OpenAPI Artifact Generator) python package does the following:
- bundles individual yaml files into a single file
- post process x- extensions
- validates the bundled openapi.yaml file
- generates enhanced ux python classes from the bundled openapi.yaml file
- creates a single combined python file for all generated and common classes
> This python package DOES NOT create a python package for the generated artifacts.


## Getting started
Install the package
```
pip install openapiart
```

Generate artifacts from OpenAPI files
```python
import openapiart

""" 
the following command produces these artifacts
    - openapi.yaml
    - openapi.json
    - openapi.html
    - sample.py
and writes them to the output_dir
"""
openapiart.OpenApiArt(
    api_files=['./tests/api/api.yaml'], 
    python_module_name='sample', 
    output_dir='../../artifacts'
)
```

## Specifications
> This repository is based on the [OpenAPI specification](
https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md) which is a standard, language-agnostic interface to RESTful APIs. 

> Modeling guide specific to this package


