# OpenAPIArt 
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
openapiart.OpenApiArt(api_files=['./tests/api/api.yaml'], python_package_name='test')
```

## Specifications
> This repository is based on the [OpenAPI specification](
https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md) which is a standard, language-agnostic interface to RESTful APIs. 

> Modeling guide specific to this package


