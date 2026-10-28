import inspect
from types import SimpleNamespace
from unittest.mock import MagicMock

import pytest


@pytest.mark.parametrize("api_class", ["HttpApi", "GrpcApi"])
@pytest.mark.parametrize(
    "method_name, parameters",
    [("set_config", ("self", "payload")), ("get_config", ("self",))],
)
def test_generated_method_metadata(api_class, method_name, parameters):
    cls = getattr(pytest.module.sanity, api_class)
    method = getattr(cls, method_name)
    original = inspect.unwrap(method)

    assert original is not method
    assert method.__wrapped__ is original
    assert method.__name__ == method_name
    assert method.__qualname__ == "{}.{}".format(api_class, method_name)
    assert method.__doc__ == original.__doc__
    assert tuple(inspect.signature(method).parameters) == parameters
    assert inspect.signature(method) == inspect.signature(original)

    instance = cls()
    bound_method = getattr(instance, method_name)
    assert bound_method.__qualname__ == method.__qualname__
    assert inspect.unwrap(bound_method) is original
    assert tuple(inspect.signature(bound_method).parameters) == parameters[1:]


def test_decorator_preserves_metadata():
    def operation(self, payload: str, *, timeout: int = 10) -> str:
        """Apply the requested operation."""
        return payload

    decorated = pytest.module.sanity.Telemetry.create_child_span(operation)

    assert decorated is not operation
    for attribute in (
        "__name__",
        "__qualname__",
        "__module__",
        "__doc__",
        "__annotations__",
    ):
        assert getattr(decorated, attribute) == getattr(operation, attribute)
    assert inspect.unwrap(decorated) is operation
    assert inspect.signature(decorated) == inspect.signature(operation)


@pytest.mark.parametrize("enabled", [False, True])
@pytest.mark.parametrize("raises", [False, True])
def test_decorator_execution(enabled, raises):
    tracer = MagicMock()
    client_kind = object()
    instance = SimpleNamespace(
        _telemetry=SimpleNamespace(
            is_telemetry_enabled=enabled,
            _spankind=SimpleNamespace(CLIENT=client_kind),
        ),
        tracer=MagicMock(return_value=tracer),
    )
    calls = []
    payload = object()
    result = object()
    error = ValueError("operation failed")

    def operation(self, *args, **kwargs):
        calls.append((self, args, kwargs))
        if enabled:
            tracer.start_as_current_span.return_value.__enter__.assert_called_once_with()
            tracer.start_as_current_span.return_value.__exit__.assert_not_called()
        if raises:
            raise error
        return result

    decorated = pytest.module.sanity.Telemetry.create_child_span(operation)
    if raises:
        with pytest.raises(ValueError) as exc_info:
            decorated(instance, payload, timeout=10)
        assert exc_info.value is error
    else:
        assert decorated(instance, payload, timeout=10) is result

    assert calls == [(instance, (payload,), {"timeout": 10})]
    if enabled:
        instance.tracer.assert_called_once_with()
        tracer.start_as_current_span.assert_called_once_with(
            "operation", kind=client_kind
        )
        span = tracer.start_as_current_span.return_value
        if raises:
            span.__exit__.assert_called_once()
            exc_type, exc_value, traceback = span.__exit__.call_args[0]
            assert exc_type is ValueError
            assert exc_value is error
            assert traceback is not None
        else:
            span.__exit__.assert_called_once_with(None, None, None)
    else:
        instance.tracer.assert_not_called()
        tracer.start_as_current_span.assert_not_called()
