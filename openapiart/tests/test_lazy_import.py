"""Tests to verify that grpc and protobuf imports are lazily loaded.

The generated code should NOT import grpc, protobuf stubs (*_pb2, *_pb2_grpc)
at module level. These should only be imported inside GrpcApi.__init__ via
importlib.import_module, so that users who only need HTTP transport are not
required to have grpc/protobuf installed.
"""

import ast
import os
import sys
import pytest


def _get_generated_source():
    """Return the source code of the generated sanity.py file."""
    src_path = os.path.join(
        pytest.artifacts_path, pytest.module_name, pytest.module_name + ".py"
    )
    with open(src_path, "r") as f:
        return f.read()


def _get_top_level_imports(source):
    """Parse the module AST and return all top-level import names."""
    tree = ast.parse(source)
    imports = set()
    for node in ast.iter_child_nodes(tree):
        if isinstance(node, ast.Import):
            for alias in node.names:
                imports.add(alias.name)
        elif isinstance(node, ast.ImportFrom):
            if node.module is not None:
                imports.add(node.module)
    return imports


def test_grpc_not_in_top_level_imports():
    """grpc should not be imported at the module level."""
    source = _get_generated_source()
    top_imports = _get_top_level_imports(source)
    assert (
        "grpc" not in top_imports
    ), "grpc should not be a top-level import in the generated module"


def test_protobuf_stubs_not_in_top_level_imports():
    """protobuf stubs (sanity_pb2, sanity_pb2_grpc) should not be imported
    at the module level."""
    source = _get_generated_source()
    top_imports = _get_top_level_imports(source)
    pb2_imports = [
        name
        for name in top_imports
        if name.endswith("_pb2") or name.endswith("_pb2_grpc")
    ]
    assert len(pb2_imports) == 0, (
        "protobuf stubs should not be top-level imports, found: %s"
        % pb2_imports
    )


def test_grpc_lazy_loaded_in_grpc_api_init():
    """The GrpcApi.__init__ should call importlib.import_module for grpc,
    protobuf json_format, and the pb2 stubs."""
    source = _get_generated_source()
    tree = ast.parse(source)

    # Find the GrpcApi class
    grpc_api_class = None
    for node in ast.walk(tree):
        if isinstance(node, ast.ClassDef) and node.name == "GrpcApi":
            grpc_api_class = node
            break
    assert (
        grpc_api_class is not None
    ), "GrpcApi class not found in generated code"

    # Find __init__ method
    init_method = None
    for item in grpc_api_class.body:
        if isinstance(item, ast.FunctionDef) and item.name == "__init__":
            init_method = item
            break
    assert init_method is not None, "GrpcApi.__init__ not found"

    # Collect all importlib.import_module call string arguments in __init__
    lazy_imports = set()
    for node in ast.walk(init_method):
        if isinstance(node, ast.Call):
            func = node.func
            # Match importlib.import_module(...) or self._xxx = importlib.import_module(...)
            is_importlib_call = False
            if (
                isinstance(func, ast.Attribute)
                and func.attr == "import_module"
            ):
                if isinstance(func.value, ast.Attribute):
                    is_importlib_call = True
                elif (
                    isinstance(func.value, ast.Name)
                    and func.value.id == "importlib"
                ):
                    is_importlib_call = True
            if is_importlib_call and node.args:
                arg = node.args[0]
                if isinstance(arg, ast.Constant) and isinstance(
                    arg.value, str
                ):
                    lazy_imports.add(arg.value)

    assert (
        "grpc" in lazy_imports
    ), "grpc should be lazily imported in GrpcApi.__init__"
    assert (
        "google.protobuf.json_format" in lazy_imports
    ), "google.protobuf.json_format should be lazily imported in GrpcApi.__init__"
    pb2_lazy = [m for m in lazy_imports if "pb2" in m]
    assert len(pb2_lazy) >= 2, (
        "Both pb2 and pb2_grpc stubs should be lazily imported in "
        "GrpcApi.__init__, found: %s" % pb2_lazy
    )


def test_http_api_does_not_load_grpc(monkeypatch):
    """Creating an HttpApi should not trigger any grpc or pb2 imports.

    We verify this by temporarily making grpc un-importable and confirming
    that HttpApi instantiation still works.
    """
    import importlib

    original_import = importlib.import_module

    grpc_modules_loaded = []

    def tracking_import(name, *args, **kwargs):
        if "grpc" in name or "_pb2" in name:
            grpc_modules_loaded.append(name)
        return original_import(name, *args, **kwargs)

    monkeypatch.setattr(importlib, "import_module", tracking_import)

    # Clear any previously cached state
    grpc_modules_loaded.clear()

    # Creating an HTTP API should not trigger grpc/pb2 imports
    module = pytest.module
    http_api = module.api(
        location="http://127.0.0.1:12345",
        transport=module.Transport.HTTP,
        verify=False,
    )
    assert http_api is not None

    grpc_related = [
        m for m in grpc_modules_loaded if "grpc" in m or "_pb2" in m
    ]
    assert len(grpc_related) == 0, (
        "Creating HttpApi should not trigger grpc/pb2 imports, but loaded: %s"
        % grpc_related
    )


def test_grpc_api_triggers_lazy_imports(monkeypatch):
    """Creating a GrpcApi should trigger lazy imports of grpc and pb2 modules."""
    import importlib

    original_import = importlib.import_module

    grpc_modules_loaded = []

    def tracking_import(name, *args, **kwargs):
        if "grpc" in name or "_pb2" in name:
            grpc_modules_loaded.append(name)
        return original_import(name, *args, **kwargs)

    monkeypatch.setattr(importlib, "import_module", tracking_import)

    grpc_modules_loaded.clear()

    module = pytest.module
    grpc_api = module.api(
        location="localhost:50051",
        transport=module.Transport.GRPC,
    )
    assert grpc_api is not None

    # Verify grpc and pb2 modules were loaded
    assert any(
        "grpc" == m or m.endswith("_grpc") for m in grpc_modules_loaded
    ), ("GrpcApi should lazily import grpc, loaded: %s" % grpc_modules_loaded)
    assert any("_pb2" in m for m in grpc_modules_loaded), (
        "GrpcApi should lazily import pb2 stubs, loaded: %s"
        % grpc_modules_loaded
    )
    grpc_api.close()
