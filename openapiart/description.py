def render_description(openapi_object, default=None):
    """Return description text with rendered OpenAPIArt metadata.

    The returned value is intended for generated comments/docstrings. The
    OpenAPI object is not modified.
    """
    description = openapi_object.get("description", default)
    parts = []

    status = openapi_object.get("x-status")
    if status is not None:
        status_type = status.get("status")
        if status_type is not None:
            status_type = status_type.replace("-", "_")
        information = status.get("information", "Information TBD")
        if status_type == "deprecated":
            parts.append("Deprecated: {}".format(information))
        elif status_type == "under_review":
            parts.append("Under Review: {}".format(information))

    if description is not None and len(description) > 0:
        parts.append(description)

    constraints = openapi_object.get("x-constraint")
    if constraints is not None and len(constraints) > 0:
        lines = ["x-constraint:"]
        lines.extend("- {}".format(constraint) for constraint in constraints)
        # Preserve the trailing blank comment line emitted by the former
        # bundler-based rendering path.
        parts.append("\n".join(lines) + "\n")

    return "\n\n".join(parts)
