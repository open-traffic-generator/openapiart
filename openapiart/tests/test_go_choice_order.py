import os

from openapiart.openapiart import OpenApiArt as openapiart_class


# Test that the enum getters are outputted in the same order
# as the enum definition.  This was previously using a set,
# which altered the order on each run so created unnecessary
# diffs.  The real part of the schema that matters here is the
# ChoiceHolder choice enum, everything else is plumbing.
def test_go_choice_getters_preserve_enum_order(tmp_path):
    api_file = tmp_path / "api.yaml"
    api_file.write_text(
        """
paths:
  /choice:
    get:
      operationId: get_choice
      description: Gets the choice holder.
      responses:
        "200":
          description: Choice holder response.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ChoiceHolder"
          x-field-uid: 1
        default:
          description: Unexpected error.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Error"
          x-field-uid: 2
components:
  schemas:
    ChoiceHolder:
      description: Object with choice enum values that have no backing properties.
      type: object
      properties:
        choice:
          description: Selects one of the no-property choices.
          type: string
          default: alpha
          x-field-uid: 1
          x-enum:
            alpha:
              x-field-uid: 1
            beta:
              x-field-uid: 2
            gamma:
              x-field-uid: 3
            delta:
              x-field-uid: 4
    Error:
      description: Error response generated while serving API request.
      type: object
      required:
        - code
        - errors
      properties:
        code:
          description: Numeric status code.
          type: integer
          format: int32
          x-field-uid: 1
        errors:
          description: Error messages.
          type: array
          items:
            type: string
          x-field-uid: 2
""",
        encoding="utf-8",
    )
    artifact_dir = tmp_path / "artifacts"

    openapiart_class(
        api_files=[
            os.path.join(os.path.dirname(__file__), "api", "info.yaml"),
            str(api_file),
        ],
        artifact_dir=str(artifact_dir),
        extension_prefix="orderpkg",
    ).GenerateGoSdk(
        package_dir="github.com/open-traffic-generator/openapiart/orderpkg",
        package_name="orderpkg",
    )

    generated = (tmp_path / "orderpkg" / "choice_holder.go").read_text(
        encoding="utf-8"
    )
    # Get the offsets of Alpha, Beta, Gamma, Delta in the generated file.
    getter_order = [
        generated.index("func (obj *choiceHolder) Alpha()"),
        generated.index("func (obj *choiceHolder) Beta()"),
        generated.index("func (obj *choiceHolder) Gamma()"),
        generated.index("func (obj *choiceHolder) Delta()"),
    ]

    # This assertion is a little obscure; it asserts that the ordering
    # is preserved - Alpha is before Beta is before Gamma is before Delta.
    # It will show just a list of integers being different, but the
    # integers are offsets in the file, and they should be ordered.
    assert getter_order == sorted(getter_order)
